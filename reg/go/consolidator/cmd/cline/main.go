package main

import (
	"context"
	"time"

	"consolidator/internal/adapters/driven"
	"consolidator/internal/adapters/driver"
	"consolidator/internal/core/service"
)

// main is the entry point of the consolidation application. It initializes the necessary components and starts the consolidation process.
func main() {
	// Load configuration
	cfg, err := driven.NewConfig("consolidator.json")
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
	// Initialize the repository (this is a placeholder, you would need to implement the actual repository)
	ctx := context.Background()
	repository, err := driven.NewPostgresRepository(cfg, &ctx)
	if err != nil {
		logger.IPrintf(0, "Failed to initialize repository: %v\n", err)
		return
	}
	defer repository.Close()
	// Initialize and run the consolidation service
	service, err := service.NewConsolidateService(repository, logger, cfg)
	if err != nil {
		logger.IPrintf(0, "Failed to initialize consolidation service: %v\n", err)
		return
	}
	// Initialize driver and run the service
	driver := driver.NewFlagDriver(service)
	if err := driver.Run(); err != nil {
		logger.IPrintf(0, "Error running the consolidation service: %v\n", err)
	}
	logger.IPrintf(0, "Fully exiting command line\n")
}
