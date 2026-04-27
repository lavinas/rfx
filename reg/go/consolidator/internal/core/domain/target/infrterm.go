package target_domain

import (
	source_domain "consolidator/internal/core/domain/source"
	"consolidator/internal/core/ports"
)

// Infrterm represents the consolidated terminal data for a specific year and quarter.
type Infrterm struct {
	infrterm *InfrtermItem
	consolidation map[string]*InfrtermItem
}

// NewInfrterm creates a new instance of Infrterm.
func NewInfrterm() *Infrterm {
	return &Infrterm{
		infrterm: NewInfrtermItem(),
		consolidation: make(map[string]*InfrtermItem),
	}
}

// Delete removes the consolidated data for a specific year and quarter from the consolidation map.
func (i *Infrterm) Delete(year int, quarter int, repository ports.Repository) error {
	// delete the consolidated data for the specified year and quarter from the repository
	if err := repository.Delete(&InfrtermItem{}, year, quarter); err != nil {
		return err
	}
	return nil
}

// Save persists the consolidated data for a specific year and quarter to the repository.
func (i *Infrterm) Save(repository ports.Repository) error {
	if err := repository.Save(i.consolidation); err != nil {
		return err
	}
	return nil
}

// AddTransactions processes a slice of transactions and updates the Infrterm instance with the total transaction amount and quantity.
func (i *Infrterm) AddTransactions(year int, quarter int, transactions []*source_domain.Transaction) {
}

// AddEstablishments processes a slice of establishments and updates the Infrterm instance accordingly.
func (i *Infrterm) AddEstablishments(year int, quarter int, establishments []*source_domain.Establishment) {
}

// AddTerminals processes a slice of terminals and updates the Infrterm instance accordingly.
func (i *Infrterm) AddTerminals(year int, quarter int, terminals []*source_domain.Terminal, esblishmentMap map[int64]string) {
	for _, t := range terminals {
		infrterm := i.infrterm.GetFromTerminal(year, quarter, t, esblishmentMap)
		if infrterm == nil {
			continue
		}
		key := infrterm.GetKey()
		if existing, exists := i.consolidation[key]; exists {
			existing.PosTotalQuantity += infrterm.PosSharedQuantity + infrterm.PosChipQuantity + infrterm.PdvQuantity
			existing.PosSharedQuantity += infrterm.PosSharedQuantity
			existing.PosChipQuantity += infrterm.PosChipQuantity
			existing.PdvQuantity += infrterm.PdvQuantity
			i.consolidation[key] = existing
		} else {
			i.consolidation[key] = infrterm
		}
	}
}
