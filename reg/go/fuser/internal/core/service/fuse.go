package service

import (
	"fmt"

	"fuser/internal/ports"
)

// FuseService is the service layer that interacts with the repository to perform business logic
type FuseService struct {
	Repository ports.Repository
}

// NewFuseService creates a new instance of FuseService with the provided repository
func NewFuseService(repository ports.Repository) *FuseService {
	return &FuseService{Repository: repository}
}

// Run executes the main logic of the FuseService (placeholder for actual implementation)
func (s *FuseService) Run() error {
	// Placeholder for the main logic of the service, such as fetching transactions and processing them
	fmt.Println("Running FuseService...")

	return nil
}
