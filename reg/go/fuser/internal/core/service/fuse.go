package service

import (
	"time"

	"fuser/internal/ports"
)

// FuseService is the service layer that interacts with the repository to perform business logic
type FuseService struct {
	Repository ports.Repository
	Logger    ports.Logger
}

// NewFuseService creates a new instance of FuseService with the provided repository and logger
func NewFuseService(repository ports.Repository, logger ports.Logger) *FuseService {
	return &FuseService{Repository: repository, Logger: logger}
}

// Run executes the main logic of the FuseService (placeholder for actual implementation)
func (s *FuseService) Run(start_date time.Time, end_date time.Time) error {
	s.Logger.Println("Running FuseService...")
	// Placeholder for the main logic of the service, such as fetching transactions and processing them
	for date := start_date; !date.After(end_date); date = date.AddDate(0, 0, 1) {
		s.Logger.Printf("Processing transactions for date: %s\n", date.Format("2006-01-02"))
	}
	s.Logger.Println("Finished processing transactions.")
	return nil
}
