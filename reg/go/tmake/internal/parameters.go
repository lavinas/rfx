package internal

import (
	"fmt"
	"log"
	"os"
	"time"
)

// GetParameters gets start date, end date and quantity of transactions from command line arguments
func GetParameters() (time.Time, time.Time, int) {
	if len(os.Args) != 4 {
		log.Fatal("Usage: tmake <start_date> <end_date> <quantity>")
	}
	startDate, err := time.Parse("2006-01-02", os.Args[1])
	if err != nil {
		log.Fatal("Invalid start date format. Use YYYY-MM-DD.")
	}
	endDate, err := time.Parse("2006-01-02", os.Args[2])
	if err != nil {
		log.Fatal("Invalid end date format. Use YYYY-MM-DD.")
	}
	qty := 0
	_, err = fmt.Sscanf(os.Args[3], "%d", &qty)
	if err != nil || qty <= 0 {
		log.Fatal("Quantity must be a positive integer.")
	}
	return startDate, endDate, qty
}