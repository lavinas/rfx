package target_domain

import (
	source_domain "consolidator/internal/core/domain/source"
	"consolidator/internal/core/ports"
	"math"
)

// Desconto represents the consolidated discount transactions for a specific year and quarter.
type Desconto struct {
	desconto      *DescontoItem
	consolidation map[string]*DescontoItem
}

// NewDesconto creates a new instance of Desconto.
func NewDesconto() *Desconto {
	return &Desconto{
		desconto:      NewDescontoItem(),
		consolidation: make(map[string]*DescontoItem),
	}
}

// Delete removes the consolidated data for a specific year and quarter from the consolidation map.
func (i *Desconto) Delete(year int, quarter int, repository ports.Repository) error {
	// delete the consolidated data for the specified year and quarter from the repository
	if err := repository.Delete(&DescontoItem{}, year, quarter); err != nil {
		return err
	}
	return nil
}

// Save persists the consolidated data for a specific year and quarter to the repository.
func (i *Desconto) Save(repository ports.Repository) error {
	if err := repository.Save(i.consolidation); err != nil {
		return err
	}
	return nil
}

// GetFromTransactions returns a map of Desconto structs for a given list of transactions.
func (i *Desconto) AddTransactions(transactions []*source_domain.Transaction) {
	// for each transaction, get the corresponding Desconto instance and update the transaction amount, quantity and mdr fee statistics
	for _, t := range transactions {
		desconto := i.desconto.GetFromTransaction(t)
		key := desconto.GetKey()

		// if the key already exists in the items map, update the existing Desconto instance with the new transaction data
		if existing, exists := i.consolidation[key]; exists {
			// sum amount and quantity
			existing.TransactionAmount += desconto.TransactionAmount
			existing.TransactionQuantity += desconto.TransactionQuantity
			// update min and max MDR fees
			if desconto.MinMDRFee < existing.MinMDRFee {
				existing.MinMDRFee = desconto.MinMDRFee
			}
			if desconto.MaxMDRFee > existing.MaxMDRFee {
				existing.MaxMDRFee = desconto.MaxMDRFee
			}
			// update average and standard deviation using Welford's online algorithm
			delta := desconto.AvgMDRFee - existing.AvgMDRFee
			existing.AvgMDRFee += delta / float64(existing.TransactionQuantity)
			delta2 := desconto.AvgMDRFee - existing.AvgMDRFee
			existing.SqrdiffMDRFee += delta2 * delta2
			variance := existing.SqrdiffMDRFee / (float64(existing.TransactionQuantity) - 1)
			existing.StdevMDRFee = math.Sqrt(variance)
			i.consolidation[key] = existing
		} else {
			i.consolidation[key] = desconto
		}
	}
}

// AddEstablishments processes a slice of establishments and updates the Desconto instance with the number of accredited and active establishments.
func (i *Desconto) AddEstablishments(year int, quarter int, establishments []*source_domain.Establishment) {
}

// AddTerminals processes a slice of terminals and updates the Desconto instance with the number of accredited and active terminals.
func (i *Desconto) AddTerminals(year int, quarter int, terminals []*source_domain.Terminal, esblishmentMap map[int64]string) {
}
