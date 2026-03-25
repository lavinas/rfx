package internal

import (
	"fmt"
	"log"
	"time"
)

// InsertTransactions generates and inserts transactions into the database
func InsertTransactions(transactionStartDate, transactionEndDate time.Time, transactionQtty int, db *DB) {
	transaction := NewTransaction()
	lastID := transaction.GetLastID(db)
	// Generate and insert transactions
    var count int64 = 1
	for date := transactionStartDate; date.Before(transactionEndDate) || date.Equal(transactionEndDate); date = date.AddDate(0, 0, 1) {
		for i := 1; i <= transactionQtty; i++ {
			transaction.GenerateData(count+lastID, date)
			transaction.Insert(db)
			if i%1000 == 0 {
				db.Commit()
				log.Printf("Inserted %d transactions for date %s\n", i, date.Format("2006-01-02"))
			}
			count++
		}
		db.Commit()
		log.Printf("Finished inserting %d transactions for date %s\n", transactionQtty, date.Format("2006-01-02"))
	}
	fmt.Println("Transactions inserted successfully")
}
