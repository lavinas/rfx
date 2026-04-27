package target_domain

import (
	source_domain "consolidator/internal/core/domain/source"
	"consolidator/internal/core/ports"
)

// Intercam represents the consolidated interchange fee transactions for a specific year and quarter.
type Intercam struct {
	Intercam *IntercamItem
	consolidation map[string]*IntercamItem
}

// NewIntercam creates a new instance of Intercam.
func NewIntercam(bins map[int64]*source_domain.Bin) *Intercam {
	return &Intercam{
		Intercam: NewIntercamItem(bins),
		consolidation: make(map[string]*IntercamItem),
	}
}

// Delete removes the consolidated data for a specific year and quarter from the consolidation map.
func (i *Intercam) Delete(year int, quarter int, repository ports.Repository) error {
	// delete the consolidated data for the specified year and quarter from the repository
	if err := repository.Delete(&IntercamItem{}, year, quarter); err != nil {
		return err
	}
	return nil
}

// Save persists the consolidated data for a specific year and quarter to the repository.
func (i *Intercam) Save(repository ports.Repository) error {
	if err := repository.Save(i.consolidation); err != nil {
		return err
	}
	return nil
}


// Translate transforms the Intercam struct into a format suitable for database storage, if necessary.
func (i *Intercam) AddTransactions(transactions []*source_domain.Transaction) {
	for _, t := range transactions {
		interchange := i.Intercam.GetFromTransaction(t)
		key := interchange.GetKey()
		if existing, exists := i.consolidation[key]; exists {
			existing.TransactionAmount += interchange.TransactionAmount
			existing.TransactionQuantity += interchange.TransactionQuantity
			delta := interchange.InterchangeFee - existing.InterchangeFee
			existing.InterchangeFee += delta / float64(existing.TransactionQuantity)
			i.consolidation[key] = existing
		} else {
			i.consolidation[key] = interchange
		}
	}
}

// AddEstablishments processes a slice of establishments and updates the Intercam instance with the number of accredited and active establishments.
func (i *Intercam) AddEstablishments(year int, quarter int, establishments []*source_domain.Establishment) {
}

// AddTerminals processes a slice of terminals and updates the Intercam instance with the number of accredited and active terminals.
func (i *Intercam)AddTerminals(year int, quarter int, terminals []*source_domain.Terminal, esblishmentMap map[int64]string) {
}


