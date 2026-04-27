package service

import (
	"time"

	domain_source "consolidator/internal/core/domain/source"
	domain_target "consolidator/internal/core/domain/target"
)

// runTransaction executes the consolidation process for a specific date.
func (s *ConsolidateService) runTransaction(year int, quarter int, days int) error {
	// Log the start of the consolidation process
	s.Logger.IPrintf(1, "Starting consolidation transaction process for year: %d, quarter: %d\n", year, quarter)

	// Calculate the start and end dates for the specified quarter
	start_date, end_date := s.getDates(year, quarter, days)

	// Get bins for mapping BIN numbers to product and card type information
	bins, err := s.GetBins()
	if err != nil {
		s.Logger.IPrintf(2, "Error fetching BIN information: %v\n", err)
		return err
	}

	// Process transactions for each date in the specified range and update the consolidation maps
	for date := start_date; !date.After(end_date); date = date.AddDate(0, 0, 1) {
		if err := s.processDate(date, bins); err != nil {
			s.Logger.IPrintf(2, "Error processing date %s: %v\n", date.Format("2006-01-02"), err)
			return err
		}
	}

	// filter ranking data if the filter_ranking flag is set
	s.Logger.IPrintf(1, "Filtering ranking data from %d records\n", len(s.ranking))
	s.rankingFiltered = s.FilterRanking(s.ranking)
	s.Logger.IPrintf(1, "Filtered ranking data to %d records\n", len(s.rankingFiltered))


	// Log the completion of the consolidation process
	s.Logger.IPrintf(1, "Consolidation transaction process completed successfully for year: %d, quarter: %d\n", year, quarter)
	return nil
}

// GetBins retrieves BIN information from the repository
func (s *ConsolidateService) GetBins() (map[int64]*domain_source.Bin, error) {
	s.Logger.IPrintf(2, "Fetching BIN information from the repository\n")
	bins := make(map[int64]*domain_source.Bin)
	binList, err := s.Repository.GetBins()
	if err != nil {
		return nil, err
	}
	for _, bin := range binList {
		bins[bin.Bin] = bin
	}
	s.Logger.IPrintf(2, "Fetched %d BIN records from the repository\n", len(bins))
	return bins, nil
}

// processDate processes transactions for a specific date and updates the consolidated data maps
func (s *ConsolidateService) processDate(date time.Time, bins map[int64]*domain_source.Bin) error {
	s.Logger.IPrintf(2, "Processing for date: %s\n", date.Format("2006-01-02"))

	// Fetch transactions for the date
	transactions, err := s.Repository.GetTransactionsByDate(date)
	if err != nil {
		s.Logger.IPrintf(2, "Error fetching transactions for date %s: %v\n", date.Format("2006-01-02"), err)
		return err
	}
	s.Logger.IPrintf(3, "Got %d transactions for date: %s\n", len(transactions), date.Format("2006-01-02"))

	// Add transactions to the respective consolidated data maps
	domain_target.NewDesconto().AddTransactions(transactions, s.desconto)
	s.Logger.IPrintf(3, "Consolidated Desconto for date: %s\n", date.Format("2006-01-02"))
	domain_target.NewRanking().AddTransactions(transactions, s.ranking)
	s.Logger.IPrintf(3, "Consolidated Ranking for date: %s\n", date.Format("2006-01-02"))
	domain_target.NewIntercam(bins).AddTransactions(transactions, s.intercam, bins)
	s.Logger.IPrintf(3, "Consolidated Intercam for date: %s\n", date.Format("2006-01-02"))
	domain_target.NewConccred().AddTransactions(transactions, s.concred)
	s.Logger.IPrintf(3, "Consolidated ConcCred for date: %s\n", date.Format("2006-01-02"))
	domain_target.NewSegmento().AddTransactions(transactions, s.segmento)
	s.Logger.IPrintf(3, "Consolidated Segmento for date: %s\n", date.Format("2006-01-02"))

	// Log the completion of processing for the date
	s.Logger.IPrintf(2, "Processed  for date: %s\n", date.Format("2006-01-02"))
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
