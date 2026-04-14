package service

import (
	"time"

	"consolidator/internal/core/ports"
	domain_target "consolidator/internal/core/domain/target"
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
func (s *ConsolidateService) Run(year int, quarter int, focus string, delete bool, start *time.Time, end *time.Time) error {
	// Log the start of the consolidation process
	s.Logger.Printf("Starting consolidation process for year: %d, quarter: %d\n", year, quarter)

	var desconto domain_target.Desconto
	descontoMap := make(map[string]*domain_target.Desconto)

	// Loop through the dates in the specified quarter and process transactions
	start_date, end_date := s.getQuarterDates(year, quarter)
	if start != nil {
		start_date = *start
	}
	if end != nil {
		end_date = *end
	}
	for date := start_date; !date.After(end_date); date = date.AddDate(0, 0, 1) {
		s.Logger.Printf("	Processing transactions for date: %s\n", date.Format("2006-01-02"))
		transactions, err := s.Repository.GetTransactionsByDate(date)
		if err != nil {
			s.Logger.Printf("	Error fetching transactions for date %s: %v\n", date.Format("2006-01-02"), err)
			continue
		}
		s.Logger.Printf("	Got %d transactions for date: %s\n", len(transactions), date.Format("2006-01-02"))
		// Process transactions and consolidate data
		desconto.AddTransactions(transactions, descontoMap)
		s.Logger.Printf("	Consolidated Desconto for date: %s\n", date.Format("2006-01-02"))

	}

	// Transform the descontoMap to a slice for saving to the database
	descontoSlice := make([]*domain_target.Desconto, 0, len(descontoMap))
	for _, d := range descontoMap {
		descontoSlice = append(descontoSlice, d)
	}

	// Save the consolidated Desconto data to the database
	s.Logger.Printf("Saving %d consolidated Desconto data to the database\n", len(descontoMap))
	if err := s.Repository.SaveDesconto(descontoSlice); err != nil {
		s.Logger.Printf("	Error saving consolidated Desconto: %v\n", err)
		return err
	}	
	s.Logger.Printf("Saved %d consolidated Desconto data to the database successfully\n", len(descontoMap))

	// Consolidate data based on the focus (e.g., "transactions", "accounts", etc.)
	s.Logger.Printf("Consolidating data with focus: %s\n", focus)


	// Log the completion of the consolidation process
	s.Logger.Printf("Consolidation process completed successfully for year: %d, quarter: %d\n", year, quarter)
	return nil
}


// getQuarterDates calculates the start and end dates for a given year and quarter.
func (s *ConsolidateService) getQuarterDates(year int, quarter int) (time.Time, time.Time) {
	var start_date, end_date time.Time
	switch quarter {
	case 1:
		start_date = time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
		end_date = time.Date(year, time.March, 31, 23, 59, 59, 0, time.UTC)
	case 2:
		start_date = time.Date(year, time.April, 1, 0, 0, 0, 0, time.UTC)
		end_date = time.Date(year, time.June, 30, 23, 59, 59, 0, time.UTC)
	case 3:
		start_date = time.Date(year, time.July, 1, 0, 0, 0, 0, time.UTC)
		end_date = time.Date(year, time.September, 30, 23, 59, 59, 0, time.UTC)
	case 4:
		start_date = time.Date(year, time.October, 1, 0, 0, 0, 0, time.UTC)
		end_date = time.Date(year, time.December, 31, 23, 59, 59, 0, time.UTC)
	}
	return start_date, end_date
}
