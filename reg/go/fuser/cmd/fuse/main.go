package main

import (
	"context"
	"os"

	"fuser/internal/adapters/driven"
	"fuser/internal/adapters/driver"
	"fuser/internal/core/service"

)

// Main function to initialize and run the application
func main() {
	// Initialize the logger
	logger := driven.NewSimpleLogger()
	// Load configuration
	cfg, err := driven.LoadJsonConfig("config.json")
	if err != nil {
		logger.Println("Error loading configuration:", err)
		os.Exit(1)
	}
	// Initialize the repository with the database connection
	ctx := context.Background()
	repo, err := driven.NewGormRepository(cfg.GetDNS(), &ctx)
	if err != nil {
		logger.Println("Error initializing repository:", err)
		os.Exit(1)
	}
	defer repo.Close()
	// Create the driver with the service
	drv := driver.NewFlagDriver(service.NewFuseService(repo, logger))
	// Run the driver
	if err := drv.Run(); err != nil {
		logger.Println(err)
		os.Exit(1)
	}
}
