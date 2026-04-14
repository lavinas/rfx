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
func (s *ConsolidateService) Run(year int, quarter int, delete bool, start *time.Time, end *time.Time) error {
	// Log the start of the consolidation process
	s.Logger.IPrintf(0, "Starting consolidation process for year: %d, quarter: %d\n", year, quarter)
	
	// Calculate the start and end dates for the specified quarter
	start_date, end_date := s.getQuarterDates(year, quarter, start, end)

	// Delete existing consolidated data if the delete flag is set
	if err := s.deleteAll(year, quarter, delete); err != nil {
		s.Logger.IPrintf(1, "Error deleting existing consolidated data: %v\n", err)
		return err
	}

	// consolidation maps to hold the consolidated data for each type
	descontoMap := make(map[string]*domain_target.Desconto)
	rankingMap := make(map[string]*domain_target.Ranking)
	intercamMap := make(map[string]*domain_target.Intercam)
	conccredMap := make(map[string]*domain_target.ConcCred)
	for date := start_date; !date.After(end_date); date = date.AddDate(0, 0, 1) {
		if err := s.processDate(date, descontoMap, rankingMap, intercamMap, conccredMap); err != nil {
			s.Logger.IPrintf(1, "Error processing date %s: %v\n", date.Format("2006-01-02"), err)
			return err
		}
	}

	// save consolidated data to the database
	if err := s.saveAll(descontoMap, rankingMap, intercamMap, conccredMap); err != nil {
		s.Logger.IPrintf(1, "Error saving consolidations: %v\n", err)
		return err
	}

	// Log the completion of the consolidation process
	s.Logger.IPrintf(0, "Consolidation process completed successfully for year: %d, quarter: %d\n", year, quarter)
	return nil
}


// deleteAll deletes existing consolidated data from the database for the specified year and quarter
func (s *ConsolidateService) deleteAll(year int, quarter int, delete bool) error {
	// if delete flag is false, skip deletion and log it
	if !delete {
		s.Logger.IPrintf(1, "Skipping deletion of existing consolidated data for year: %d, quarter: %d\n", year, quarter)
		return nil
	}

	// delete discounts
	s.Logger.IPrintf(1, "Deleting existing consolidated data for year: %d, quarter: %d\n", year, quarter)
	if err := s.Repository.DeleteDesconto(year, quarter); err != nil {
		s.Logger.IPrintf(2, "Error deleting Desconto data: %v\n", err)
		return err
	}

	// delete ranking
	if err := s.Repository.DeleteRanking(year, quarter); err != nil {
		s.Logger.IPrintf(2, "Error deleting Ranking data: %v\n", err)
		return err	
	}

	// delete intercam
	if err := s.Repository.DeleteIntercam(year, quarter); err != nil {
		s.Logger.IPrintf(2, "Error deleting Intercam data: %v\n", err)
		return err
	}

	// delete conccred
	if err := s.Repository.DeleteConcCred(year, quarter); err != nil {
		s.Logger.IPrintf(2, "Error deleting ConcCred data: %v\n", err)
		return err	
	}
	s.Logger.IPrintf(1, "Deleted existing consolidated data  for year: %d, quarter: %d\n", year, quarter)

	return nil
}

// processDate processes transactions for a specific date and updates the consolidated data maps
func (s *ConsolidateService) processDate(date time.Time, descontoMap map[string]*domain_target.Desconto, rankingMap map[string]*domain_target.Ranking, intercamMap map[string]*domain_target.Intercam, conccredMap map[string]*domain_target.ConcCred) error {
	s.Logger.IPrintf(1, "Processing for date: %s\n", date.Format("2006-01-02"))

	// Fetch transactions for the date
	transactions, err := s.Repository.GetTransactionsByDate(date)
	if err != nil {
		s.Logger.IPrintf(2, "Error fetching transactions for date %s: %v\n", date.Format("2006-01-02"), err)
		return err
	}
	s.Logger.IPrintf(2, "Got %d transactions for date: %s\n", len(transactions), date.Format("2006-01-02"))
	
	// Process transactions and consolidate data
	var desconto domain_target.Desconto
	var ranking domain_target.Ranking
	var intercam domain_target.Intercam
	var conccred domain_target.ConcCred
	
	// Add transactions to the respective consolidated data maps
	desconto.AddTransactions(transactions, descontoMap)
	s.Logger.IPrintf(2, "Consolidated Desconto for date: %s\n", date.Format("2006-01-02"))
	ranking.AddTransactions(transactions, rankingMap)
	s.Logger.IPrintf(2, "Consolidated Ranking for date: %s\n", date.Format("2006-01-02"))
	intercam.AddTransactions(transactions, intercamMap)
	s.Logger.IPrintf(2, "Consolidated Intercam for date: %s\n", date.Format("2006-01-02"))
	conccred.AddTransactions(transactions, conccredMap)
	s.Logger.IPrintf(2, "Consolidated ConcCred for date: %s\n", date.Format("2006-01-02"))
	s.Logger.IPrintf(1, "Processed  for date: %s\n", date.Format("2006-01-02"))
	return nil
}

// saveAll saves all the consolidated data to the database
func (s *ConsolidateService) saveAll(descontoMap map[string]*domain_target.Desconto, rankingMap map[string]*domain_target.Ranking, intercamMap map[string]*domain_target.Intercam, conccredMap map[string]*domain_target.ConcCred) error {
	s.Logger.IPrintf(1, "Saving consolidated data to the database\n")
	
	// save discounts
	if err := s.saveDesconto(descontoMap); err != nil {
		return err
	}

	// save ranking
	if err := s.saveRanking(rankingMap); err != nil {
		return err
	}

	// save intercam
	if err := s.saveIntercam(intercamMap); err != nil {
		return err
	}
	
	// save conccred
	if err := s.saveConcCred(conccredMap); err != nil {
		return err
	}
	s.Logger.IPrintf(1, "Saved consolidated data to the database\n")
	return nil
}

// saveDesconto saves the consolidated Desconto data to the database
func (s *ConsolidateService) saveDesconto(descontoMap map[string]*domain_target.Desconto) error {
	s.Logger.IPrintf(1, "Saving %d consolidated Desconto\n", len(descontoMap))
	descontoList := make([]*domain_target.Desconto, 0, 2000)
	count := 0
	for _, desconto := range descontoMap {
		descontoList = append(descontoList, desconto)
		count++
		if count%2000 == 0 {
			if err := s.Repository.SaveDesconto(descontoList); err != nil {
				return err
			}
			s.Logger.IPrintf(2, "Saved batch of 2000 consolidated Desconto\n")
			descontoList = make([]*domain_target.Desconto, 0, 2000)
		}
	}
	if err := s.Repository.SaveDesconto(descontoList); err != nil {
		return err
	}
	s.Logger.IPrintf(1, "Saved  %d consolidated Desconto\n", len(descontoMap))


	return s.Repository.SaveDesconto(descontoList)
}

// saveRanking saves the consolidated Ranking data to the database
func (s *ConsolidateService) saveRanking(rankingMap map[string]*domain_target.Ranking) error {
	s.Logger.IPrintf(1, "Saving %d consolidated Ranking\n", len(rankingMap))
	rankingList := make([]*domain_target.Ranking, 0, 2000)
	count := 0
	for _, ranking := range rankingMap {
		rankingList = append(rankingList, ranking)
		count++
		if count%2000 == 0 {
			if err := s.Repository.SaveRanking(rankingList); err != nil {
				return err
			}
			s.Logger.IPrintf(2, "Saved batch of 2000 consolidated Ranking\n")
			rankingList = make([]*domain_target.Ranking, 0, 2000)
		}
	}
	if err := s.Repository.SaveRanking(rankingList); err != nil {
		return err
	}
	s.Logger.IPrintf(1, "Saved  %d consolidated Ranking\n", len(rankingMap))
	return nil
}

// saveIntercam saves the consolidated Intercam data to the database
func (s *ConsolidateService) saveIntercam(intercamMap map[string]*domain_target.Intercam) error {
	s.Logger.IPrintf(1, "Saving %d consolidated Intercam\n", len(intercamMap))
	intercamList := make([]*domain_target.Intercam, 0, 2000)
	count := 0
	for _, intercam := range intercamMap {
		intercamList = append(intercamList, intercam)
		count++
		if count%2000 == 0 {
			if err := s.Repository.SaveIntercam(intercamList); err != nil {
				return err
			}
			s.Logger.IPrintf(2, "Saved batch of 2000 consolidated Intercam\n")
			intercamList = make([]*domain_target.Intercam, 0, 2000)
		}
	}
	if err := s.Repository.SaveIntercam(intercamList); err != nil {
		return err
	}
	s.Logger.IPrintf(1, "Saved  %d consolidated Intercam\n", len(intercamMap))
	return nil
}

// saveConcCred saves the consolidated ConcCred data to the database
func (s *ConsolidateService) saveConcCred(conccredMap map[string]*domain_target.ConcCred) error {
	s.Logger.IPrintf(1, "Saving %d consolidated ConcCred\n", len(conccredMap))
	conccredList := make([]*domain_target.ConcCred, 0, 2000)
	count := 0
	for _, conccred := range conccredMap {
		conccredList = append(conccredList, conccred)
		count++
		if count%2000 == 0 {
			if err := s.Repository.SaveConcCred(conccredList); err != nil {
				return err
			}
			s.Logger.IPrintf(2, "Saved batch of 2000 consolidated ConcCred\n")
			conccredList = make([]*domain_target.ConcCred, 0, 2000)
		}
	}
	if err := s.Repository.SaveConcCred(conccredList); err != nil {
		return err
	}
	s.Logger.IPrintf(1, "Saved  %d consolidated ConcCred\n", len(conccredMap))
	return nil
}

// getQuarterDates calculates the start and end dates for a given year and quarter.
func (s *ConsolidateService) getQuarterDates(year int, quarter int, start, end *time.Time) (time.Time, time.Time) {
	startMonth := (quarter-1)*3 + 1
	start_date := time.Date(year, time.Month(startMonth), 1, 0, 0, 0, 0, time.UTC)
	end_date := start_date.AddDate(0, 3, -1) // end of the quarter is 3 months later minus 1 day
	// override with provided start and end dates if they are not nil
	if start != nil {
		start_date = *start
	}
	if end != nil {
		end_date = *end
	}
	return start_date, end_date
}
