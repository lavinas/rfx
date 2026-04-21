package service

import (
	source_domain "consolidator/internal/core/domain/source"
	domain_target "consolidator/internal/core/domain/target"
)

// runOthers executes the consolidation process for other data types (e.g., Infresta, INfreterm, etc) for a specific date.
func (s *ConsolidateService) runOthers(year int, quarter int) error {
	// Log the start of the consolidation process for other data types
	s.Logger.IPrintf(1, "Running consolidation for other data types for year %d and quarter %d\n", year, quarter)

	// delete existing consolidated data for other data types
	if err := s.deleteOthers(year, quarter); err != nil {
		s.Logger.IPrintf(1, "Error deleting existing consolidated data for other data types: %v\n", err)
		return err
	}

	// read establishments from the database
	s.Logger.IPrintf(2, "*Reading establishments from the database\n")
	establishments, err := s.Repository.GetEstablishments()
	if err != nil {
		s.Logger.IPrintf(2, "Error fetching establishments: %v\n", err)
		return err
	}
	s.Logger.IPrintf(2, "*Read %d establishments\n", len(establishments))

	// read terminals from the database
	s.Logger.IPrintf(2, "*Reading terminals from the database\n")
	terminals, err := s.Repository.GetTerminals()
	if err != nil {
		s.Logger.IPrintf(2, "Error fetching terminals: %v\n", err)
		return err
	}
	s.Logger.IPrintf(2, "*Read %d terminals\n", len(terminals))

	// read ConcCred data from the database
	s.Logger.IPrintf(2, "*Reading ConcCred data from the database\n")
	conccredData, err := s.Repository.GetConcCred(year, quarter)
	if err != nil {
		s.Logger.IPrintf(2, "Error fetching ConcCred data: %v\n", err)
		return err
	}
	conccredMap := make(map[string]*domain_target.ConcCred)
	for _, c := range conccredData {
		conccredMap[c.GetKey()] = c
	}
	s.Logger.IPrintf(2, "*Read %d ConcCred records\n", len(conccredData))

	// process others consolidation
	s.Logger.IPrintf(2, "Processing consolidation for other data types\n")
	infrestaMap := make(map[string]*domain_target.Infresta)
	domain_target.NewInfresta().AddEstablishments(year, quarter, establishments, infrestaMap)
	infrtermMap := make(map[string]*domain_target.Infrterm)
	domain_target.NewInfrterm().AddTerminals(year, quarter, s.getEstablishmentFUMap(year, quarter, establishments), terminals, infrtermMap)
	s.Logger.IPrintf(2, "Processed consolidation for other data types\n")
	domain_target.NewConcCred().AddEstablishments(year, quarter, establishments, conccredMap)

	// save consolidated data for other data types to the database
	if err := s.saveOthers(infrestaMap, infrtermMap, conccredMap); err != nil {
		s.Logger.IPrintf(1, "Error saving consolidated data for other data types: %v\n", err)
		return err
	}

	// Log the completion of the consolidation process for other data types
	s.Logger.IPrintf(1, "Completed consolidation for other data types for year %d and quarter %d\n", year, quarter)

	// Placeholder for the actual consolidation logic for other data types (e.g., Infresta, INfreterm, etc) for a specific date.
	return nil

}

// deleteAll deletes all consolidated data for a specific year and quarter from the database.
func (s *ConsolidateService) deleteOthers(year int, quarter int) error {
	s.Logger.IPrintf(2, "Deleting existing consolidated data for year %d and quarter %d\n", year, quarter)

	// delete Infrterm data
	if err := s.Repository.DeleteInfrterm(year, quarter); err != nil {
		s.Logger.IPrintf(2, "Error deleting Infrterm data: %v\n", err)
		return err
	}

	// delete Infresta data
	if err := s.Repository.DeleteInfresta(year, quarter); err != nil {
		s.Logger.IPrintf(2, "Error deleting Infresta data: %v\n", err)
		return err
	}

	// Placeholder for deleting other consolidated data types (e.g., Infresta, INfreterm, etc) for a specific date.
	return nil
}

// saveOthers saves the consolidated data for other data types (e.g., Infresta, INfreterm, etc) to the database.
func (s *ConsolidateService) saveOthers(infrestaMap map[string]*domain_target.Infresta, infrtermMap map[string]*domain_target.Infrterm,
	conccredMap map[string]*domain_target.ConcCred) error {

	// save Infresta
	if err := s.saveInfresta(infrestaMap); err != nil {
		s.Logger.IPrintf(2, "Error saving Infresta data: %v\n", err)
		return err
	}

	// save Infrterm
	if err := s.saveInfrterm(infrtermMap); err != nil {
		s.Logger.IPrintf(2, "Error saving Infrterm data: %v\n", err)
		return err
	}

	// save ConcCred
	if err := s.saveConcCred(conccredMap); err != nil {
		s.Logger.IPrintf(2, "Error saving ConcCred data: %v\n", err)
		return err
	}

	return nil
}

// saveInfresta saves the consolidated Infresta data to the database.
func (s *ConsolidateService) saveInfresta(infrestaMap map[string]*domain_target.Infresta) error {
	// Log the number of consolidated Infresta records being saved
	s.Logger.IPrintf(2, "Saving %d consolidated Infresta\n", len(infrestaMap))

	// Convert the map of Infresta to a slice for batch saving
	infrestaList := make([]*domain_target.Infresta, 0, 2000)
	count := 0

	// Iterate over the infrestaMap and append each Infresta to the infrestaList slice
	for _, infresta := range infrestaMap {
		infrestaList = append(infrestaList, infresta)
		count++
		if count%2000 == 0 {
			if err := s.Repository.SaveInfresta(infrestaList); err != nil {
				return err
			}
			s.Logger.IPrintf(2, "Saved batch of 2000 consolidated Infresta\n")
			infrestaList = make([]*domain_target.Infresta, 0, 2000)
		}
	}

	// Save any remaining Infresta records that didn't make up a full batch
	if len(infrestaList) > 0 {
		if err := s.Repository.SaveInfresta(infrestaList); err != nil {
			return err
		}
		s.Logger.IPrintf(2, "Saved remaining %d consolidated Infresta\n", len(infrestaList))
	}
	return nil
}

// saveInfrterm saves the consolidated Infrterm data to the database.
func (s *ConsolidateService) saveInfrterm(infrtermMap map[string]*domain_target.Infrterm) error {
	// Log the number of consolidated Infrterm records being saved
	s.Logger.IPrintf(2, "Saving %d consolidated Infrterm\n", len(infrtermMap))

	// Convert the map of Infrterm to a slice for batch saving
	infrtermList := make([]*domain_target.Infrterm, 0, 2000)
	count := 0

	// Iterate over the infrtermMap and append each Infrterm to the infrtermList slice
	for _, infrterm := range infrtermMap {
		infrtermList = append(infrtermList, infrterm)
		count++
		if count%2000 == 0 {
			if err := s.Repository.SaveInfrterm(infrtermList); err != nil {
				return err
			}
			s.Logger.IPrintf(2, "Saved batch of 2000 consolidated Infrterm\n")
			infrtermList = make([]*domain_target.Infrterm, 0, 2000)
		}
	}

	// Save any remaining Infrterm records that didn't make up a full batch
	if len(infrtermList) > 0 {
		if err := s.Repository.SaveInfrterm(infrtermList); err != nil {
			return err
		}
		s.Logger.IPrintf(2, "Saved remaining %d consolidated Infrterm\n", len(infrtermList))
	}
	return nil
}

// getEstablishmentFUMap creates a map of establishment codes to their corresponding federation units.
func (s *ConsolidateService) getEstablishmentFUMap(year int, quarter int, establishments []*source_domain.Establishment) map[int64]string {
	establishmentFUMap := make(map[int64]string)
	for _, e := range establishments {
		if !e.IsAccredited(year, quarter) {
			continue
		}
		establishmentFUMap[e.GetCode()] = e.GetFederationUnit()
	}
	return establishmentFUMap
}
