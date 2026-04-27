package service

import (
	"time"

	domain_source "consolidator/internal/core/domain/source"
	domain_target "consolidator/internal/core/domain/target"
	"consolidator/internal/core/ports"
)

// ConsolidateService is responsible for consolidating data from various sources and storing it in the target database.
type ConsolidateService struct {
	Repository     ports.Repository
	Logger         ports.Logger
	Config         ports.Config
	consolidations map[string]ports.Domain
}

func NewConsolidateService(repository ports.Repository, logger ports.Logger, config ports.Config) (*ConsolidateService, error) {
	bins, err := GetBins(repository, logger)
	if err != nil {
		return nil, err
	}
	consolidations := map[string]ports.Domain{
		"Desconto": domain_target.NewDesconto(),
		"Ranking":  domain_target.NewRanking(),
		"Intercam": domain_target.NewIntercam(bins),
		"ConcCred": domain_target.NewConccred(),
		"Segmento": domain_target.NewSegmento(),
		"infrterm": domain_target.NewInfrterm(),
		"infresta": domain_target.NewInfresta(),
	}
	return &ConsolidateService{
		Repository:     repository,
		Logger:         logger,
		Config:         config,
		consolidations: consolidations,
	}, nil
}

// GetBins retrieves BIN information from the repository
func GetBins(repository ports.Repository, logger ports.Logger) (map[int64]*domain_source.Bin, error) {
	logger.IPrintf(0, "Fetching BIN information from the repository\n")
	bins := make(map[int64]*domain_source.Bin)
	binList, err := repository.GetBins()
	if err != nil {
		return nil, err
	}
	for _, bin := range binList {
		bins[bin.Bin] = bin
	}
	logger.IPrintf(0, "Fetched %d BIN records from the repository\n", len(bins))
	return bins, nil
}

// Run executes the consolidation process for a specific date.
func (s *ConsolidateService) Run(year int, quarter int, days int) error {
	s.Logger.IPrintf(0, "Starting consolidation process for year: %d, quarter: %d\n", year, quarter)
	// delete existing consolidated data for transactions if the delete flag is set
	if err := s.delete(year, quarter); err != nil {
		return err
	}
	// consolidate transactions for the specified year and quarter
	if err := s.consolidateTransactions(year, quarter, days); err != nil {
		return err
	}
	// consolidate establishments for the specified year and quarter and get a map of establishment codes to federation units
	estabMap, err := s.consolidateEstablishments(year, quarter)
	if err != nil {
		return err
	}
	// consolidate terminals for the specified year and quarter using the map of establishment codes to federation units
	if err := s.consolidateTerminals(year, quarter, estabMap); err != nil {
		return err
	}
	// build the consolidated data
	if err := s.build(year, quarter); err != nil {
		return err
	}
	// save the consolidated data to the repository
	if err := s.save(year, quarter); err != nil {
		return err
	}
	s.Logger.IPrintf(0, "Consolidation process completed successfully for year: %d, quarter: %d\n", year, quarter)

	return nil
}

// consolidateTransactions processes transactions for a specific date and updates the consolidated data maps
func (s *ConsolidateService) consolidateTransactions(year int, quarter int, days int) error {
	s.Logger.IPrintf(1, "Consolidating transactions for year: %d, quarter: %d\n", year, quarter)
	// calulate start and end dates for the specified quarter
	startDate, endDate := s.getDates(year, quarter, days)
	// iterate over each date in the specified range and process transactions for that date
	for date := startDate; !date.After(endDate); date = date.AddDate(0, 0, 1) {
		s.Logger.IPrintf(2, "Processing transactions for date: %s\n", date.Format("2006-01-02"))
		// Fetch transactions for the date
		transactions, err := s.Repository.GetTransactionsByDate(date)
		if err != nil {
			return err
		}
		s.Logger.IPrintf(3, "Got %d transactions for date: %s\n", len(transactions), date.Format("2006-01-02"))
		// Update each consolidation with the transactions for the date
		for name, consolidation := range s.consolidations {
			consolidation.AddTransactions(transactions)
			s.Logger.IPrintf(3, "Consolidated transactions for %s for date: %s\n", name, date.Format("2006-01-02"))
		}
		s.Logger.IPrintf(2, "Processed transactions for date: %s\n", date.Format("2006-01-02"))
	}
	s.Logger.IPrintf(1, "Consolidated transactions for year: %d, quarter: %d\n", year, quarter)
	return nil
}

// consolidateEstablishments processes establishments for a specific date and updates the consolidated data maps
func (s *ConsolidateService) consolidateEstablishments(year int, quarter int) (map[int64]string, error) {
	s.Logger.IPrintf(1, "Consolidating establishments for year: %d, quarter: %d\n", year, quarter)
	// Fetch establishments for the specified year and quarter
	establishments, err := s.Repository.GetEstablishments()
	if err != nil {
		return nil, err
	}
	for name, consolidation := range s.consolidations {
		consolidation.AddEstablishments(year, quarter, establishments)
		s.Logger.IPrintf(2, "Consolidated establishments for %s for year: %d, quarter: %d\n", name, year, quarter)
	}
	estabMap := make(map[int64]string)
	for _, estab := range establishments {
		if !estab.IsAccredited(year, quarter) {
			continue
		}
		estabMap[estab.EstablishmentCode] = *estab.FederationUnit
	}
	s.Logger.IPrintf(1, "Consolidated establishments for year: %d, quarter: %d\n", year, quarter)
	return estabMap, nil
}

// consolidateTerminals processes terminals for a specific date and updates the consolidated data maps
func (s *ConsolidateService) consolidateTerminals(year int, quarter int, establishments map[int64]string) error {
	s.Logger.IPrintf(1, "Consolidating terminals for year: %d, quarter: %d\n", year, quarter)
	// Fetch terminals for the specified year and quarter
	terminals, err := s.Repository.GetTerminals()
	if err != nil {
		return err
	}
	for name, consolidation := range s.consolidations {
		consolidation.AddTerminals(year, quarter, terminals, establishments)
		s.Logger.IPrintf(2, "Consolidated terminals for %s for year: %d, quarter: %d\n", name, year, quarter)
	}
	s.Logger.IPrintf(1, "Consolidated terminals for year: %d, quarter: %d\n", year, quarter)
	return nil
}

// getDates calculates the start and end dates for a given year and quarter.
func (s *ConsolidateService) getDates(year int, quarter int, days int) (time.Time, time.Time) {
	// Calculate the default start and end dates for the specified quarter
	startMonth := (quarter-1)*3 + 1
	start_date := time.Date(year, time.Month(startMonth), 1, 0, 0, 0, 0, time.UTC)

	// end date is the last day of the quarter, which is 3 months later minus 1 day
	end_date := start_date.AddDate(0, 3, -1)
	if days > 0 {
		end_date = start_date.AddDate(0, 0, days-1) // add the specified number of days to the start date
	}
	// Return the calculated start and end dates
	return start_date, end_date
}

// delete deletes existing consolidated data from the database for the specified year and quarter
func (s *ConsolidateService) delete(year int, quarter int) error {
	// Log the start of deleting existing consolidated data for the specified year and quarter
	s.Logger.IPrintf(1, "Deleting existing consolidated data for year: %d, quarter: %d\n", year, quarter)

	// Iterate over each consolidation and delete the existing data for the specified year and quarter
	for name, consolidation := range s.consolidations {
		if err := consolidation.Delete(year, quarter, s.Repository); err != nil {
			return err
		}
		s.Logger.IPrintf(2, "Deleted existing consolidated data for %s for year: %d, quarter: %d\n", name, year, quarter)
	}

	// Log the completion of deleting existing consolidated data for the specified year and quarter
	s.Logger.IPrintf(1, "Deleted existing consolidated data for year: %d, quarter: %d\n", year, quarter)
	return nil
}

// build calls the Build method on each consolidation to perform any necessary processing or calculations after all transactions, establishments, and terminals have been added.
func (s *ConsolidateService) build(year int, quarter int) error {
	s.Logger.IPrintf(1, "Building consolidated data for year: %d, quarter: %d\n", year, quarter)
	for name, consolidation := range s.consolidations {
		consolidation.Build()
		s.Logger.IPrintf(2, "Built consolidated data for %s for year: %d, quarter: %d\n", name, year, quarter)
	}
	s.Logger.IPrintf(1, "Built consolidated data for year: %d, quarter: %d\n", year, quarter)
	return nil
}

// save saves the consolidated data to the repository for the specified year and quarter
func (s *ConsolidateService) save(year int, quarter int) error {
	// Log the start of saving consolidated data for the specified year and quarter
	s.Logger.IPrintf(1, "Saving consolidated data for year: %d, quarter: %d\n", year, quarter)

	// Iterate over each consolidation and save the consolidated data for the specified year and quarter
	for name, consolidation := range s.consolidations {
		if err := consolidation.Save(s.Repository); err != nil {
			return err
		}
		s.Logger.IPrintf(2, "Saved consolidated data for %s for year: %d, quarter: %d\n", name, year, quarter)
	}

	// Log the completion of saving consolidated data for the specified year and quarter
	s.Logger.IPrintf(1, "Saved consolidated data for year: %d, quarter: %d\n", year, quarter)
	return nil
}
