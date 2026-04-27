package target_domain

import (
	source_domain "consolidator/internal/core/domain/source"
	"consolidator/internal/core/ports"
)

// 
type Ranking struct {
	Ranking *RankingItem
	consolidation map[string]*RankingItem
}

// NewRanking creates a new instance of Ranking.
func NewRanking() *Ranking {
	return &Ranking{
		Ranking: NewRankingItem(),
		consolidation: make(map[string]*RankingItem),
	}
}

// Delete removes the consolidated data for a specific year and quarter from the consolidation map.
func (i *Ranking) Delete(year int, quarter int, repository ports.Repository) error {
	// delete the consolidated data for the specified year and quarter from the repository
	if err := repository.Delete(&RankingItem{}, year, quarter); err != nil {
		return err
	}
	return nil
}

// Save persists the consolidated data for a specific year and quarter to the repository.
func (i *Ranking) Save(repository ports.Repository) error {
	if err := repository.Save(i.consolidation); err != nil {
		return err
	}
	return nil
}

// GetFromTransactions processes a slice of transactions and returns a map of Ranking instances keyed by their unique keys.
func (i *Ranking) AddTransactions(transactions []*source_domain.Transaction) {
	// for each transaction, get the corresponding Ranking instance and update the transaction amount, quantity and average mcc fee
	for _, t := range transactions {
		ranking := i.Ranking.GetFromTransaction(t)
		key := ranking.GetKey()
		if existing, exists := i.consolidation[key]; exists {
			existing.TransactionAmount += ranking.TransactionAmount
			existing.TransactionQuantity += ranking.TransactionQuantity
			delta := ranking.AvgMccFee - existing.AvgMccFee
			existing.AvgMccFee += delta / float64(existing.TransactionQuantity)
		} else {
			i.consolidation[key] = ranking
		}
	}
}

// AddEstablishments processes a slice of establishments and updates the Ranking instance with the number of accredited and active establishments.
func (i *Ranking) AddEstablishments(year int, quarter int, establishments []*source_domain.Establishment) {
}

// AddTerminals processes a slice of terminals and updates the Ranking instance with the number of accredited and active terminals.
func (i *Ranking) AddTerminals(year int, quarter int, terminals []*source_domain.Terminal, esblishmentMap map[int64]string) {
}