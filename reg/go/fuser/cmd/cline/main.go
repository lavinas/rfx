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
	// Create the driver with the service
	drv := driver.NewFlagDriver(service.NewFuseService(repo, logger))
	// Run the driver
	if err := drv.Run(); err != nil {
		logger.IPrintf(0, "Error running FlagDriver: %v\n", err)
		return
	}
	logger.IPrintf(0, "Fully exiting command line\n")
}
