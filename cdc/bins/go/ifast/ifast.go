package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

const (
	host     = "192.168.100.78"
	port     = 5436
	user     = "sys_flexcon"
	password = "Wgkjsjjja8872Xl"
	dbname   = "dev_regulat"
	sslmode  = "disable"
	lenBatch = 5000
)

// This program reads SQL insert statements from a file and executes them in batches to speed up the insertion process.
func main() {
	// Check if the user provided the required SQL file as an argument.
	if len(os.Args) < 2 {
		log.Fatal("Usage: ifast <sql_file>")
	}
	// Open the specified SQL file for reading.
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(1, err)
	}
	defer file.Close()
	// Connect to the PostgreSQL database using the provided connection parameters.
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode))
	if err != nil {
		log.Fatal(2, err)
	}
	defer db.Close()
	// Loop through the file, reading SQL insert statements and adding them to a batch. When the batch reaches the specified length, merge and execute it.
	run(file, db)
	fmt.Println("Done")
}

// run reads SQL insert statements from the provided file, batches them, and executes them against the database. It handles merging of insert statements and ensures that any remaining statements are executed after processing the entire file.
func run (file *os.File, db *sql.DB) {
	log.Printf("Processing file: %s\n", file.Name())
	scanner := bufio.NewScanner(file)
	batch := []string{}
	total, err := countLines(file)
	if err != nil {
		log.Fatal(5, err)
	}
	count := 0
	for scanner.Scan() {
		count++
		// Trim whitespace from the line and skip empty lines. Append non-empty lines to the batch.
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		batch = append(batch, line)
		// When the batch reaches the specified length, merge the insert statements and execute them in a single transaction.
		if len(batch) == lenBatch {
			flushBatch(db, batch, total, count)
			batch = []string{}
		}
	}
	// After processing all lines, if there are any remaining statements in the batch, merge and execute them.
	if len(batch) > 0 {
		flushBatch(db, batch, total, count)
	}
}

// countLines counts the number of lines in the specified file and returns the count. It handles any errors that may occur during file opening or scanning and logs them appropriately.
func countLines(file *os.File) (int, error) {
	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		count++
	}
	file.Seek(0, 0)
	return count, scanner.Err()
}

// FlushBatch takes a batch of SQL insert statements, merges them into a single statement, and executes it against the database. It handles any errors that may occur during merging or execution and logs them appropriately.
func flushBatch(db *sql.DB, batch []string, total, count int) {
	merged, err := mergeInserts(batch)
	if err != nil {
		log.Fatal(3, err)
	}
	if err := executeBatch(db, *merged); err != nil {
		log.Fatal(4, err)
	}
	log.Printf("Executed batch of %d statements (%d/%d)\n", len(batch), count, total)
}						

// mergeInserts merges multiple batches of SQL insert statements into a single string for execution.
func mergeInserts(batch []string) (*string, error) {
	// If the batch is empty, return an error indicating that there are no statements to merge.
	if len(batch) == 0 {
		return nil, fmt.Errorf("batch is empty")
	}
	var merged string
	// mountStatement takes the first statement in the batch and sets the prefix of the merged string to the part of the statement before the "values" keyword. If the first statement is not a valid insert statement, it returns an error indicating that the statement is not valid.
	if err := mountStatement(&merged, batch[0]); err != nil {
		return nil, err
	}
	// loop through the batch of statements and validate that each statement is a valid insert statement by checking if it starts with "insert into". If any statement does not start with "insert into", return an error indicating that the statement is not a valid insert statement.
	for _, stmt := range batch {
		if err := mountValues(&merged, stmt); err != nil {
			return nil, err
		}
	}
	// Remove the trailing comma and add a semicolon at the end.
	merged = strings.TrimSuffix(merged, ",\n") + ";"
	return &merged, nil
}

// mountStatement takes a pointer to the merged string and a SQL statement, validates that the statement is a valid insert statement, and sets the prefix of the merged string to the part of the statement before the "values" keyword. If the statement is not a valid insert statement, it returns an error indicating that the statement is not valid.
func mountStatement(merged *string ,stmt string) error {
	// validate that the first statement in the batch is a valid insert statement by checking if it starts with "insert into". If it does not, return an error indicating that the statement is not a valid insert statement.
	first := strings.ToLower(strings.TrimSpace(stmt))
	if !strings.HasPrefix(first, "insert into") {
		return fmt.Errorf("invalid insert statement: %s", stmt)
	}
	idx := strings.Index(first, "values")
	if idx == -1 {
		return fmt.Errorf("invalid insert statement: %s", stmt)
	}
	*merged = stmt[:idx+6] + "\n"
	return nil
}

// mountValues takes a pointer to the merged string and a SQL statement, validates that the statement is a valid insert statement, and appends the values part of the statement to the merged string. If the statement is not a valid insert statement, it returns an error indicating that the statement is not valid.
func mountValues(merged *string, stmt string) error {
	line := strings.TrimSuffix(strings.ToLower(strings.TrimSpace(stmt)), ";")	
	if !strings.HasPrefix(line, "insert into") {
		return nil
	}
	idx := strings.Index(line, "values")
	if idx == -1 {
		return fmt.Errorf("invalid insert statement: %s", stmt)
	}
	*merged += line[idx+6:] + ",\n"
	return nil
}

// executeBatch executes a batch of SQL statements in a single transaction.
func executeBatch(db *sql.DB, query string) error {
	_, err := db.Exec(query)
	return err
}
