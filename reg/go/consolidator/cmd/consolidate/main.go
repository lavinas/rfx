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

	// Initialize the logger
	logger := driven.NewSimpleLogger()

	// Initialize the configuration
	config, err := driven.NewJsonConfig("consolidator.json")
	if err != nil {
		logger.Printf("Failed to initialize configuration: %v\n", err)
		return
	}

	// Set the local time zone based on the configuration
	loc, _ := time.LoadLocation(config.GetDBTimeZone())
	time.Local = loc

	// Initialize the repository (this is a placeholder, you would need to implement the actual repository)
	ctx := context.Background()
	repository, err := driven.NewGormRepository(config, &ctx)
	if err != nil {
		logger.Printf("Failed to initialize repository: %v\n", err)
		return
	}
	defer repository.Close()

	// Initialize and run the consolidation service
	service, err := service.NewConsolidateService(repository, logger, config)
	if err != nil {
		logger.Printf("Failed to initialize consolidation service: %v\n", err)
		return
	}

	// Initialize driver and run the service
	driver := driver.NewFlagDriver(service)
	if err := driver.Run(); err != nil {
		logger.Printf("Error running the consolidation service: %v\n", err)
	}

}
