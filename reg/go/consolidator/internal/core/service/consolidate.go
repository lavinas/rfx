package service

import (
	"consolidator/internal/core/ports"
)

// ConsolidateService is responsible for consolidating data from various sources and storing it in the target database.
type ConsolidateService struct {
	Repository ports.Repository
	Logger     ports.Logger
	Config     ports.Config
}

func NewConsolidateService(repository ports.Repository, logger ports.Logger, config ports.Config) *ConsolidateService {
	return &ConsolidateService{
		Repository: repository,
		Logger:     logger,
		Config:     config,
	}
}

// Run executes the consolidation process for a specific date.
func (s *ConsolidateService) Run(year int, quarter int, days int) error {
	// Log the start of the consolidation process
	s.Logger.IPrintf(0, "Starting consolidation process for year: %d, quarter: %d\n", year, quarter)
	
	// running the consolidation process for transactions dependencies
	var err error

	// running the consolidation process for transactions dependencies
	err = s.runTransaction(year, quarter, days)
	if err != nil {
		s.Logger.IPrintf(1, "Error running consolidation transaction: %v\n", err)
		return err
	}

	// running the consolidation process for other dependencies
	err = s.runOthers(year, quarter)
	if err != nil {
		s.Logger.IPrintf(1, "Error running consolidation others: %v\n", err)
		return err
	}

	// Log the completion of the consolidation process
	s.Logger.IPrintf(0, "Consolidation process completed successfully for year: %d, quarter: %d\n", year, quarter)

	// If all processes ran successfully, return nil
	return nil
}
