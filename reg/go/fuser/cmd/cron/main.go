package main

import (
	"context"
	"os"
	"time"

	"fuser/internal/adapters/driven"
	"fuser/internal/adapters/driver"
	"fuser/internal/core/service"
)

// Main function to initialize and run the application
func main() {
	// Initialize the logger
	logger := driven.NewSimpleLogger()
	// Load configuration
	cfg, err := driven.NewConfig("fuser.json")
	if err != nil {
		logger.Println("Error loading configuration:", err)
		os.Exit(1)
	}
	// Set the local time zone based on the configuration
	loc, _ := time.LoadLocation(cfg.GetDBTimeZone())
	time.Local = loc
	// Initialize the repository with the database connection
	ctx := context.Background()
	repo, err := driven.NewPostgresRepository(cfg, &ctx)
	if err != nil {
		logger.Println("Error initializing repository:", err)
		os.Exit(1)
	}
	defer repo.Close()
	// initialize service with the repository and logger
	service := service.NewFuseService(repo, logger)
	// initialize driver with the service, logger, and configuration
	drv := driver.NewCronDriver(service, logger, cfg)
	// Run the driver
	if err := drv.Run(); err != nil {
		logger.Println(err)
		os.Exit(1)
	}
	logger.Println("Fully exiting !!")
}
