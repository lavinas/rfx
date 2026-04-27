package target_domain

import (

	source_domain "consolidator/internal/core/domain/source"
	"consolidator/internal/core/ports"
)

// Infresta represents the consolidated terminal data for a specific year and quarter.
type Infresta struct {
	infreta *InfrestaItem
	consolidation map[string]*InfrestaItem
}

// NewInfrestaConsolidation creates a new instance of Infresta.
func NewInfrestaConsolidation() *Infresta {
	return &Infresta{
		infreta: NewInfrestaItem(),
		consolidation: make(map[string]*InfrestaItem),
	}
}

// Delete removes the consolidated data for a specific year and quarter from the consolidation map.
func (i *Infresta) Delete(year int, quarter int, repository ports.Repository) error {
	// delete the consolidated data for the specified year and quarter from the repository
	if err := repository.Delete(&InfrestaItem{}, year, quarter); err != nil {
		return err
	}
	return nil
}

// Save persists the consolidated data for a specific year and quarter to the repository.
func (i *Infresta) Save(repository ports.Repository) error {
	if err := repository.Save(i.consolidation); err != nil {
		return err
	}
	return nil
}

// AddTransactions processes a slice of transactions and updates the Infresta instance with the total transaction amount and quantity.
func (i *Infresta) AddTransactions(year int, quarter int, transactions []*source_domain.Transaction) {
}

// AddEstablishments processes a slice of establishments and updates the Infresta instance accordingly.
func (i *Infresta) AddEstablishments(year int, quarter int, establishments []*source_domain.Establishment) {
	// iterate over establishments and update the infresta data accordingly
	for _, e := range establishments {

		// only consider accredited establishments for the consolidation
		if !e.IsAccredited(year, quarter) {
			continue
		}

		// consolidate infresta data for the establishment
		infresta := i.infreta.GetFromEstablishment(year, quarter, e)
		key := infresta.GetKey()
		if existing, exists := i.consolidation[key]; exists {
			existing.EstablishmentTotalQuantity += infresta.EstablishmentTotalQuantity
			existing.EstablishmentManualCaptureQuantity += infresta.EstablishmentManualCaptureQuantity
			existing.EstablishmentEletronicCaptureQuantity += infresta.EstablishmentEletronicCaptureQuantity
			existing.EstablishmentRemoteCaptureQuantity += infresta.EstablishmentRemoteCaptureQuantity
			i.consolidation[key] = existing
		} else {
			i.consolidation[key] = infresta
		}
	}
}

// AddTerminals processes a slice of terminals and updates the Infresta instance with the number of accredited and active terminals.
func (i *Infresta) AddTerminals(year int, quarter int, terminals []*source_domain.Terminal, esblishmentMap map[int64]string) {
}