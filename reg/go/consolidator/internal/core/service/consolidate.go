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

	// If all processes ran successfully, return nil
	return nil
}
