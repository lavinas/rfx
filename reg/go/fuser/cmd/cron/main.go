package main

import (
	"context"
	"time"

	"fuser/internal/adapters/driven"
	"fuser/internal/adapters/driver"
	"fuser/internal/core/service"
)

// Main function to initialize and run the application
func main() {
	// Load configuration
	cfg, err := driven.NewConfig("fuser.json")
	if err != nil {
		panic(err)
	}
	// Initialize the logger
	logger, err := driven.NewSimpleLogger(cfg)
	if err != nil {
		panic(err)
	}
	defer logger.Close()
	// Set the local time zone based on the configuration
	loc, _ := time.LoadLocation(cfg.GetDBTimeZone())
	time.Local = loc
	// Initialize the repository with the database connection
	ctx := context.Background()
	repo, err := driven.NewPostgresRepository(cfg, &ctx)
	if err != nil {
		logger.IPrintf(0, "Error initializing repository: %v\n", err)
		return
	}
	defer repo.Close()
	// initialize service with the repository and logger
	service := service.NewFuseService(repo, logger)
	// initialize driver with the service, logger, and configuration
	drv := driver.NewCronDriver(service, logger, cfg)
	// Run the driver
	if err := drv.Run(); err != nil {
		logger.IPrintf(0, "Error running CronDriver: %v\n", err)
		return
	}
	logger.IPrintf(0, "Fully exiting CronDriver\n")
}
