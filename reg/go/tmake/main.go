package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var (
	transactionDate = time.Date(2024, 6, 1, 0, 0, 0, 0, time.UTC)
	transactionQtty = 100
	host            = "localhost"
	port            = 5434
	user            = "root"
	password        = "root"
	dbname          = "reg"
)

// main function to generate and insert transactions into the database
func main() {
	// Database connection
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	transaction := NewTransaction()

	// Generate and insert transactions
	for i := 0; i < 100; i++ {
		transaction.GenerateData(int16(i), time.Now()) // Assuming you have a method to generate random data for the transaction
		transaction.Insert(db)
	}
	fmt.Println("Transactions inserted successfully")
}
