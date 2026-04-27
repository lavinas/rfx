package service

import (
	"slices"
	"maps"

	domain_target "consolidator/internal/core/domain/target"
	"consolidator/internal/core/ports"
)

// ConsolidateService is responsible for consolidating data from various sources and storing it in the target database.
type ConsolidateService struct {
	Repository      ports.Repository
	Logger          ports.Logger
	Config          ports.Config
	consolidations map[string]ports.Domain
}

func NewConsolidateService(repository ports.Repository, logger ports.Logger, config ports.Config) *ConsolidateService {
	bins := 

	consolidations := map[string]ports.Domain{
		"Desconto":        domain_target.NewDesconto(),
		"Ranking":         domain_target.NewRanking(),
		"Intercam":        domain_target.NewIntercam(config.GetBins()),
		"ConcCred":        domain_target.NewConccred(),
		"Segmento":        domain_target.NewSegmento(),
		"infrterm":        domain_target.NewInfrterm(),
		"infresta":        domain_target.NewInfresta(),
	}
	return &ConsolidateService{
		Repository:      repository,
		Logger:          logger,
		Config:          config,
	}
}

// Run executes the consolidation process for a specific date.
func (s *ConsolidateService) Run(year int, quarter int, days int) error {
	// Log the start of the consolidation process
	s.Logger.IPrintf(0, "Starting consolidation process for year: %d, quarter: %d\n", year, quarter)


	// delete existing consolidated data for transactionsif the delete flag is set
	if err := s.delete(year, quarter); err != nil {
		s.Logger.IPrintf(1, "Error deleting existing consolidated data for transactions: %v\n", err)
		return err
	}

	// running the consolidation process for transactions dependencies
	if err := s.runTransaction(year, quarter, days); err != nil {
		s.Logger.IPrintf(1, "Error running consolidation transaction: %v\n", err)
		return err
	}

	// running the consolidation process for other dependencies
	if err := s.runOthers(year, quarter); err != nil {
		s.Logger.IPrintf(1, "Error running consolidation others: %v\n", err)
		return err
	}

	//save consolidated data to the database
	if err := s.save(year, quarter); err != nil {
		s.Logger.IPrintf(1, "Error saving consolidated data: %v\n", err)
		return err
	}

	// Log the completion of the consolidation process
	s.Logger.IPrintf(0, "Consolidation process completed successfully for year: %d, quarter: %d\n", year, quarter)

	// If all processes ran successfully, return nil
	return nil
}

// delete deletes existing consolidated data from the database for the specified year and quarter
func (s *ConsolidateService) delete(year int, quarter int) error {
	// Log the start of deleting existing consolidated data for the specified year and quarter
	s.Logger.IPrintf(1, "Deleting existing consolidated data for year: %d, quarter: %d\n", year, quarter)
	delMap := map[string]interface{}{
		"Desconto":        domain_target.Desconto{},
		"Ranking":         domain_target.Ranking{},
		"RankingFiltered": domain_target.RankingFiltered{},
		"Intercam":        domain_target.Intercam{},
		"ConcCred":        domain_target.Conccred{},
		"Segmento":        domain_target.Segmento{},
		"infrterm":        domain_target.Infrterm{},
		"infresta":        domain_target.Infresta{},
	}	

	// Log the start of deleting existing consolidated data for the specified year and quarter
	s.Logger.IPrintf(1, "Deleting existing consolidated data for year: %d, quarter: %d\n", year, quarter)

	for _, model := range delMap {
		if err := s.Repository.Delete(model, year, quarter); err != nil {
			return err
		}
	}

	// Log the completion of deleting existing consolidated data for the specified year and quarter
	s.Logger.IPrintf(1, "Deleted existing consolidated data for year: %d, quarter: %d\n", year, quarter)
	return nil
}


// save saves the consolidated data for transactions to the database.
func (s *ConsolidateService) save(year, quarter int) error {
	s.Logger.IPrintf(1, "Saving consolidated data for transactions for year %d quarter %d\n", year, quarter)
	if err := s.Repository.Save(slices.Collect(maps.Values(s.desconto))); err != nil {
		return err
	}
	if err := s.Repository.Save(slices.Collect(maps.Values(s.ranking))); err != nil {
		return err
	}
	if err := s.Repository.Save(slices.Collect(maps.Values(s.rankingFiltered))); err != nil {
		return err
	}
	if err := s.Repository.Save(slices.Collect(maps.Values(s.intercam))); err != nil {
		return err
	}
	if err := s.Repository.Save(slices.Collect(maps.Values(s.concred))); err != nil {
		return err
	}
	if err := s.Repository.Save(slices.Collect(maps.Values(s.segmento))); err != nil {
		return err
	}
	if err := s.Repository.Save(slices.Collect(maps.Values(s.infrterm))); err != nil {
		return err
	}
	if err := s.Repository.Save(slices.Collect(maps.Values(s.infresta))); err != nil {
		return err
	}
	s.Logger.IPrintf(1, "Saved consolidated data for transactions for year %d quarter %d\n", year, quarter)
	return nil

}