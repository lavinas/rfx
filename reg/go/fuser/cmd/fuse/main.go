package main


import (
	"context"

	"fuser/internal/adapters/config"
	"fuser/internal/adapters/repository"
	"fuser/internal/core/service"
)

// Main function to initialize and run the application
func main() {
	// Load configuration
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		panic(err)
	}
	// Initialize the repository with the database connection
	ctx := context.Background()
	repo, err := repository.NewGormRepository(cfg.GetDNS(), &ctx)
	if err != nil {
		panic(err)
	}
	defer repo.Close()
	// Create the service layer with the repository
	svc := service.NewFuseService(repo)
	// Run the service
	if err := svc.Run(); err != nil {
		panic(err)
	}

}
