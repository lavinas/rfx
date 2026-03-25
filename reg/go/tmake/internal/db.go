package internal

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// DB struct to hold the database connection
type DB struct {
	*sql.DB
	tx *sql.Tx
	
}
// NewBD creates a new database connection and begins a transaction
func NewBD(host string, port int, user, password, dbname string, sslmode bool) *DB {
	txtSSLMode := "disable"
	if sslmode {
		txtSSLMode = "enable"
	}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, txtSSLMode)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging the database: ", err)
	}
	tx, err := db.Begin()
	if err != nil {
		log.Fatal("Error beginning transaction: ", err)
	}
	return &DB{db, tx}
}

// CloseDBConnection closes the database connection
func (db *DB) CloseDBConnection() {
	db.RollbackTransaction()
	err := db.Close()
	if err != nil {
		log.Fatal("Error closing the database connection: ", err)
	}
}

// BeginTransaction begins a new transaction
func (db *DB) BeginTransaction() {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal("Error beginning transaction: ", err)
	}
	db.tx = tx
}

// Commit commits the transaction
func (db *DB) Commit() {
	err := db.tx.Commit()
	if err != nil {
		log.Fatal("Error committing transaction: ", err)
	}
	db.BeginTransaction()
}

// RollbackTransaction rolls back the transaction in case of an error
func (db *DB) RollbackTransaction() {
	err := db.tx.Rollback()
	if err != nil {
		log.Fatal("Error rolling back transaction: ", err)
	}
	db.BeginTransaction()
}