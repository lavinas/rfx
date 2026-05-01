package main

import (
	"fmt"
	"validator/internal/adapter"
	"validator/internal/service"
)

// main function to run the ValidatorService
func main() {
	// Load configuration and initialize repository
	config, err := adapter.NewConfig("./validator.json")
	if err != nil {
		panic(err)
	}
	//
	repo, err := adapter.NewPostgresGormAdapter(config)
	if err != nil {
		panic(err)
	}
	defer repo.Close()
	// Create service instance
	service := service.NewValidatorService(repo)
	// create and run the driver
	if err := adapter.NewFlagDriver(service).Run(); err != nil {
		fmt.Printf("Error running the driver: %v\n", err)
	}
	
}
