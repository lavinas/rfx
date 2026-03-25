package main

import (
	"tmake/internal"
)

// main function to generate and insert transactions into the database
func main() {
	// Establish database connection
	config, err := internal.LoadConfig("config.json")
	if err != nil {
		panic(err)
	}
	db := internal.NewBD(config.Database.Host, config.Database.Port, config.Database.User, config.Database.Password, config.Database.DBName, config.Database.SSLMode)
	defer db.CloseDBConnection()
	// Get parameters from command line arguments
	transactionStartDate, transactionEndDate, transactionQtty := internal.GetParameters()
	// Insert transactions into the database
	internal.InsertTransactions(transactionStartDate, transactionEndDate, transactionQtty, db)
}
