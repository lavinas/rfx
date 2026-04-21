package service

import (
	"time"

	domain_source "consolidator/internal/core/domain/source"
	domain_target "consolidator/internal/core/domain/target"
)

// runTransaction executes the consolidation process for a specific date.
func (s *ConsolidateService) runTransaction(year int, quarter int, days int) error {
	// Log the start of the consolidation process
	s.Logger.IPrintf(0, "Starting consolidation transaction process for year: %d, quarter: %d\n", year, quarter)

	// Calculate the start and end dates for the specified quarter
	start_date, end_date := s.getDates(year, quarter, days)

	// Delete existing consolidated data if the delete flag is set
	if err := s.deleteTransactions(year, quarter); err != nil {
		s.Logger.IPrintf(1, "Error deleting existing consolidated data: %v\n", err)
		return err
	}

	// Get bins for mapping BIN numbers to product and card type information
	bins, err := s.GetBins()
	if err != nil {
		s.Logger.IPrintf(1, "Error fetching BIN information: %v\n", err)
		return err
	}

	// consolidation maps to hold the consolidated data for each type
	descontoMap := make(map[string]*domain_target.Desconto)
	rankingMap := make(map[string]*domain_target.Ranking)
	intercamMap := make(map[string]*domain_target.Intercam)
	conccredMap := make(map[string]*domain_target.ConcCred)
	segmentoMap := make(map[string]*domain_target.Segmento)
	luccredMap := make(map[string]*domain_target.Luccred)

	// Process transactions for each date in the specified range and update the consolidation maps
	for date := start_date; !date.After(end_date); date = date.AddDate(0, 0, 1) {
		if err := s.processDate(date, descontoMap, rankingMap, intercamMap, conccredMap, segmentoMap, luccredMap, bins); err != nil {
			s.Logger.IPrintf(1, "Error processing date %s: %v\n", date.Format("2006-01-02"), err)
			return err
		}
	}

	// filter ranking data if the filter_ranking flag is set
	s.Logger.IPrintf(1, "Filtering ranking data from %d records\n", len(rankingMap))
	rankingfilteredMap := domain_target.FilterRanking(rankingMap)
	s.Logger.IPrintf(1, "Filtered ranking data to %d records\n", len(rankingfilteredMap))

	// save consolidated data to the database
	if err := s.saveTransactions(descontoMap, rankingMap, rankingfilteredMap, intercamMap, conccredMap, segmentoMap, luccredMap); err != nil {
		s.Logger.IPrintf(1, "Error saving consolidations: %v\n", err)
		return err
	}
	// Log the completion of the consolidation process
	s.Logger.IPrintf(0, "Consolidation transaction process completed successfully for year: %d, quarter: %d\n", year, quarter)
	return nil
}

// GetBins retrieves BIN information from the repository
func (s *ConsolidateService) GetBins() (map[int64]*domain_source.Bin, error) {
	s.Logger.IPrintf(1, "Fetching BIN information from the repository\n")
	bins := make(map[int64]*domain_source.Bin)
	binList, err := s.Repository.GetBins()
	if err != nil {
		return nil, err
	}
	for _, bin := range binList {
		bins[bin.Bin] = bin
	}
	s.Logger.IPrintf(1, "Fetched %d BIN records from the repository\n", len(bins))
	return bins, nil
}

// deleteTransactions deletes existing consolidated data from the database for the specified year and quarter
func (s *ConsolidateService) deleteTransactions(year int, quarter int) error {
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

	// delete ranking filtered
	if err := s.Repository.DeleteRankingFiltered(year, quarter); err != nil {
		s.Logger.IPrintf(2, "Error deleting RankingFiltered data: %v\n", err)
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

	// delete segmento
	if err := s.Repository.DeleteSegmento(year, quarter); err != nil {
		s.Logger.IPrintf(2, "Error deleting Segmento data: %v\n", err)
		return err
	}

	// delete luccred
	if err := s.Repository.DeleteLuccred(year, quarter); err != nil {
		s.Logger.IPrintf(2, "Error deleting Luccred data: %v\n", err)
		return err
	}

	// Log the completion of deleting existing consolidated data for the specified year and quarter
	s.Logger.IPrintf(1, "Deleted existing consolidated data  for year: %d, quarter: %d\n", year, quarter)

	return nil
}

// processDate processes transactions for a specific date and updates the consolidated data maps
func (s *ConsolidateService) processDate(date time.Time, descontoMap map[string]*domain_target.Desconto, rankingMap map[string]*domain_target.Ranking,
	intercamMap map[string]*domain_target.Intercam, conccredMap map[string]*domain_target.ConcCred, segmentoMap map[string]*domain_target.Segmento,
	luccredMap map[string]*domain_target.Luccred, bins map[int64]*domain_source.Bin) error {
	s.Logger.IPrintf(1, "Processing for date: %s\n", date.Format("2006-01-02"))

	// Fetch transactions for the date
	transactions, err := s.Repository.GetTransactionsByDate(date)
	if err != nil {
		s.Logger.IPrintf(2, "Error fetching transactions for date %s: %v\n", date.Format("2006-01-02"), err)
		return err
	}
	s.Logger.IPrintf(2, "Got %d transactions for date: %s\n", len(transactions), date.Format("2006-01-02"))

	// Add transactions to the respective consolidated data maps
	domain_target.NewDesconto().AddTransactions(transactions, descontoMap)
	s.Logger.IPrintf(2, "Consolidated Desconto for date: %s\n", date.Format("2006-01-02"))
	domain_target.NewRanking().AddTransactions(transactions, rankingMap)
	s.Logger.IPrintf(2, "Consolidated Ranking for date: %s\n", date.Format("2006-01-02"))
	domain_target.NewIntercam(bins).AddTransactions(transactions, intercamMap, bins)
	s.Logger.IPrintf(2, "Consolidated Intercam for date: %s\n", date.Format("2006-01-02"))
	domain_target.NewConcCred().AddTransactions(transactions, conccredMap)
	s.Logger.IPrintf(2, "Consolidated ConcCred for date: %s\n", date.Format("2006-01-02"))
	domain_target.NewSegmento().AddTransactions(transactions, segmentoMap)
	s.Logger.IPrintf(2, "Consolidated Segmento for date: %s\n", date.Format("2006-01-02"))
	domain_target.NewLuccred().AddTransactions(transactions, luccredMap)
	s.Logger.IPrintf(2, "Consolidated Luccred for date: %s\n", date.Format("2006-01-02"))

	// Log the completion of processing for the date
	s.Logger.IPrintf(1, "Processed  for date: %s\n", date.Format("2006-01-02"))
	return nil
}

// saveAll saves all the consolidated data to the database
func (s *ConsolidateService) saveTransactions(descontoMap map[string]*domain_target.Desconto, rankingMap map[string]*domain_target.Ranking,
	rakkintFiltered map[string]*domain_target.RankingFiltered, intercamMap map[string]*domain_target.Intercam, conccredMap map[string]*domain_target.ConcCred,
	segmentoMap map[string]*domain_target.Segmento, luccredMap map[string]*domain_target.Luccred) error {
	s.Logger.IPrintf(1, "Saving consolidated data to the database\n")

	// save discounts
	if err := s.saveDesconto(descontoMap); err != nil {
		return err
	}

	// save ranking
	if err := s.saveRanking(rankingMap); err != nil {
		return err
	}

	// save ranking filtered
	if err := s.saveRankingFiltered(rakkintFiltered); err != nil {
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

	// save segmento
	if err := s.saveSegmento(segmentoMap); err != nil {
		return err
	}

	// save luccred
	if err := s.saveLuccred(luccredMap); err != nil {
		return err
	}

	// Log the completion of saving consolidated data to the database
	s.Logger.IPrintf(1, "Saved consolidated data to the database\n")
	return nil
}

// saveDesconto saves the consolidated Desconto data to the database
func (s *ConsolidateService) saveDesconto(descontoMap map[string]*domain_target.Desconto) error {
	// Log the number of consolidated Desconto records being saved
	s.Logger.IPrintf(2, "Saving %d consolidated Desconto\n", len(descontoMap))

	// Convert the map of Desconto to a slice for batch saving
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

	// Save any remaining Desconto records that were not saved in the batch loop
	if err := s.Repository.SaveDesconto(descontoList); err != nil {
		return err
	}

	// Log the completion of saving consolidated Desconto records to the database
	s.Logger.IPrintf(2, "Saved  %d consolidated Desconto\n", len(descontoMap))

	// Return nil to indicate successful completion of the save operation
	return nil
}

// saveRanking saves the consolidated Ranking data to the database
func (s *ConsolidateService) saveRanking(rankingMap map[string]*domain_target.Ranking) error {
	// Log the number of consolidated Ranking records being saved
	s.Logger.IPrintf(2, "Saving %d consolidated Ranking\n", len(rankingMap))

	// Convert the map of Ranking to a slice for batch saving
	rankingList := make([]*domain_target.Ranking, 0, 2000)
	count := 0

	// Iterate over the rankingMap and append each Ranking to the rankingList slice
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

	// Save any remaining Ranking records that were not saved in the batch loop
	if err := s.Repository.SaveRanking(rankingList); err != nil {
		return err
	}

	// Log the completion of saving consolidated Ranking records to the database
	s.Logger.IPrintf(2, "Saved  %d consolidated Ranking\n", len(rankingMap))
	return nil
}

// saveRankingFiltered saves the consolidated RankingFiltered data to the database
func (s *ConsolidateService) saveRankingFiltered(rankingFilteredMap map[string]*domain_target.RankingFiltered) error {
	// Log the number of consolidated RankingFiltered records being saved
	s.Logger.IPrintf(2, "Saving %d consolidated RankingFiltered\n", len(rankingFilteredMap))

	// Convert the map of RankingFiltered to a slice for batch saving
	rankingFilteredList := make([]*domain_target.RankingFiltered, 0, 2000)
	count := 0

	// Iterate over the rankingFilteredMap and append each RankingFiltered to the rankingFilteredList slice
	for _, rankingFiltered := range rankingFilteredMap {
		rankingFilteredList = append(rankingFilteredList, rankingFiltered)
		count++
		if count%2000 == 0 {
			if err := s.Repository.SaveRankingFiltered(rankingFilteredList); err != nil {
				return err
			}
			s.Logger.IPrintf(2, "Saved batch of 2000 consolidated RankingFiltered\n")
			rankingFilteredList = make([]*domain_target.RankingFiltered, 0, 2000)
		}
	}

	// Save any remaining RankingFiltered records that were not saved in the batch loop
	if err := s.Repository.SaveRankingFiltered(rankingFilteredList); err != nil {
		return err
	}

	// Log the completion of saving consolidated RankingFiltered records to the database
	s.Logger.IPrintf(2, "Saved  %d consolidated RankingFiltered\n", len(rankingFilteredMap))
	return nil
}

// saveIntercam saves the consolidated Intercam data to the database
func (s *ConsolidateService) saveIntercam(intercamMap map[string]*domain_target.Intercam) error {
	// Log the number of consolidated Intercam records being saved
	s.Logger.IPrintf(2, "Saving %d consolidated Intercam\n", len(intercamMap))

	// Convert the map of Intercam to a slice for batch saving
	intercamList := make([]*domain_target.Intercam, 0, 2000)
	count := 0

	// Iterate over the intercamMap and append each Intercam to the intercamList slice
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

	// Save any remaining Intercam records that were not saved in the batch loop
	if err := s.Repository.SaveIntercam(intercamList); err != nil {
		return err
	}

	// Log the completion of saving consolidated Intercam records to the database
	s.Logger.IPrintf(2, "Saved  %d consolidated Intercam\n", len(intercamMap))
	return nil
}

// saveConcCred saves the consolidated ConcCred data to the database
func (s *ConsolidateService) saveConcCred(conccredMap map[string]*domain_target.ConcCred) error {
	// Log the number of consolidated ConcCred records being saved
	s.Logger.IPrintf(2, "Saving %d consolidated ConcCred\n", len(conccredMap))

	// Convert the map of ConcCred to a slice for batch saving
	conccredList := make([]*domain_target.ConcCred, 0, 2000)
	count := 0

	// Iterate over the conccredMap and append each ConcCred to the conccredList slice
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

	// Save any remaining ConcCred records that were not saved in the batch loop
	if err := s.Repository.SaveConcCred(conccredList); err != nil {
		return err
	}

	// Log the completion of saving consolidated ConcCred records to the database
	s.Logger.IPrintf(2, "Saved  %d consolidated ConcCred\n", len(conccredMap))
	return nil
}

// saveSegmento saves the consolidated Segmento data to the database
func (s *ConsolidateService) saveSegmento(segmentoMap map[string]*domain_target.Segmento) error {
	// Log the number of consolidated Segmento records being saved
	s.Logger.IPrintf(2, "Saving %d consolidated Segmento\n", len(segmentoMap))

	// Convert the map of Segmento to a slice for batch saving
	segmentoList := make([]*domain_target.Segmento, 0, 2000)
	count := 0

	// Iterate over the segmentoMap and append each Segmento to the segmentoList slice
	for _, segmento := range segmentoMap {
		segmentoList = append(segmentoList, segmento)
		count++
		if count%2000 == 0 {
			if err := s.Repository.SaveSegmento(segmentoList); err != nil {
				return err
			}
			s.Logger.IPrintf(2, "Saved batch of 2000 consolidated Segmento\n")
			segmentoList = make([]*domain_target.Segmento, 0, 2000)
		}
	}

	// Save any remaining Segmento records that were not saved in the batch loop
	if err := s.Repository.SaveSegmento(segmentoList); err != nil {
		return err
	}

	// Log the completion of saving consolidated Segmento records to the database
	s.Logger.IPrintf(2, "Saved  %d consolidated Segmento\n", len(segmentoMap))
	return nil
}

// saveLuccred saves the consolidated Luccred data to the database
func (s *ConsolidateService) saveLuccred(luccredMap map[string]*domain_target.Luccred) error {
	// Log the number of consolidated Luccred records being saved
	s.Logger.IPrintf(2, "Saving %d consolidated Luccred\n", len(luccredMap))

	// Convert the map of Luccred to a slice for batch saving
	luccredList := make([]*domain_target.Luccred, 0, 2000)
	count := 0

	// Iterate over the luccredMap and append each Luccred to the luccredList slice
	for _, luccred := range luccredMap {
		luccredList = append(luccredList, luccred)
		count++
		if count%2000 == 0 {
			if err := s.Repository.SaveLuccred(luccredList); err != nil {
				return err
			}
			s.Logger.IPrintf(2, "Saved batch of 2000 consolidated Luccred\n")
			luccredList = make([]*domain_target.Luccred, 0, 2000)
		}
	}

	// Save any remaining Luccred records that were not saved in the batch loop
	if err := s.Repository.SaveLuccred(luccredList); err != nil {
		return err
	}

	// Log the completion of saving consolidated Luccred records to the database
	s.Logger.IPrintf(2, "Saved  %d consolidated Luccred\n", len(luccredMap))
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
