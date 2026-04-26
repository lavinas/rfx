package service

import (

	source_domain "consolidator/internal/core/domain/source"
	domain_target "consolidator/internal/core/domain/target"
)

// runOthers executes the consolidation process for other data types (e.g., Infresta, INfreterm, etc) for a specific date.
func (s *ConsolidateService) runOthers(year int, quarter int) error {
	// Log the start of the consolidation process for other data types
	s.Logger.IPrintf(1, "Running consolidation for other data types for year %d and quarter %d\n", year, quarter)

	// read establishments from the database
	s.Logger.IPrintf(2, "*Reading establishments from the database\n")
	establishments, err := s.Repository.GetEstablishments()
	if err != nil {
		s.Logger.IPrintf(2, "Error fetching establishments: %v\n", err)
		return err
	}
	s.Logger.IPrintf(2, "Read %d establishments\n", len(establishments))

	// read terminals from the database
	s.Logger.IPrintf(2, "Reading terminals from the database\n")
	terminals, err := s.Repository.GetTerminals()
	if err != nil {
		s.Logger.IPrintf(2, "Error fetching terminals: %v\n", err)
		return err
	}
	s.Logger.IPrintf(2, "Read %d terminals\n", len(terminals))

	// read ConcCred data from the database
	s.Logger.IPrintf(2, "Reading ConcCred data from the database\n")
	conccredData, err := s.Repository.GetConcCred(year, quarter)
	if err != nil {
		s.Logger.IPrintf(2, "Error fetching ConcCred data: %v\n", err)
		return err
	}
	conccredMap := make(map[string]*domain_target.ConcCred)
	for _, c := range conccredData {
		conccredMap[c.GetKey()] = c
	}
	s.Logger.IPrintf(2, "Read %d ConcCred records\n", len(conccredData))

	// process others consolidation
	s.Logger.IPrintf(2, "Processing consolidation for other data types\n")
	
	domain_target.NewInfresta().AddEstablishments(year, quarter, establishments, s.infresta)
	domain_target.NewInfrterm().AddTerminals(year, quarter, s.getEstablishmentFUMap(year, quarter, establishments), terminals, s.infrterm)
	domain_target.NewConcCred().AddEstablishments(year, quarter, establishments, s.concred)

	// Log the completion of the consolidation process for other data types
	s.Logger.IPrintf(1, "Completed consolidation for other data types for year %d and quarter %d\n", year, quarter)

	// Placeholder for the actual consolidation logic for other data types (e.g., Infresta, INfreterm, etc) for a specific date.
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
