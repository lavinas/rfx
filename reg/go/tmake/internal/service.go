package internal

import (
	"log"
	"time"
)

// InsertTransactions generates and inserts transactions into the database
func InsertTransactions(transactionStartDate, transactionEndDate time.Time, transactionQtty int, db *DB, commit_qtty int) {
	batch := NewBatch()
	lastID := batch.GetLastID(db)
	// Generate and insert transactions
	log.Println("Transactions insertion started")
	var count int64 = 1
	for date := transactionStartDate; date.Before(transactionEndDate) || date.Equal(transactionEndDate); date = date.AddDate(0, 0, 1) {
		for i := 1; i <= transactionQtty; i++ {
			transaction := NewTransaction()
			transaction.GenerateData(count+lastID, date)
			batch.AddTransaction(transaction)
			if i%commit_qtty == 0 {
				batch.Insert(db)
				db.Commit()
				log.Printf("Inserted %d transactions for date %s\n", i, date.Format("2006-01-02"))
			}
			count++

		}
		batch.Insert(db)
		db.Commit()
		log.Printf("Finished inserting %d transactions for date %s\n", transactionQtty, date.Format("2006-01-02"))		
	}
	log.Println("Transactions inserted successfully")
}
