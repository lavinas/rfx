package target_domain

import (
	source_domain "consolidator/internal/core/domain/source"
	"consolidator/internal/core/ports"
)

// Conccred represents the consolidated credit card transactions and establishment data for a specific year and quarter.
type Conccred struct {
	conccred *ConcCredItem
	consolidation map[string]*ConcCredItem
}

// NewConccred creates a new instance of Conccred.
func NewConccred() *Conccred {
	return &Conccred{
		conccred: NewConcCredItem(),
		consolidation: make(map[string]*ConcCredItem),
	}
}

// Delete removes the consolidated data for a specific year and quarter from the consolidation map.
func (i *Conccred) Delete(year int, quarter int, repository ports.Repository) error {
	// delete the consolidated data for the specified year and quarter from the repository
	if err := repository.Delete(&ConcCredItem{}, year, quarter); err != nil {
		return err
	}
	return nil
}

// Save persists the consolidated data for a specific year and quarter to the repository.
func (i *Conccred) Save(repository ports.Repository) error{
	if err := repository.Save(i.consolidation); err != nil {
		return err
	}
	return nil
}

// AddTransactions processes a slice of transactions and updates the ConcCred instance with the total transaction amount and quantity.
func (i *Conccred) AddTransactions(year int, quarter int, transactions []*source_domain.Transaction) {
	// get new ConcCred instances for each transaction and aggregate the transaction amount and quantity by brand and function
	for _, t := range transactions {
		concCred := i.conccred.GetFromTransaction(t)
		key := concCred.GetKey()
		if existing, exists := i.consolidation[key]; exists {
			existing.TransactionAmount += concCred.TransactionAmount
			existing.TransactionQuantity += concCred.TransactionQuantity
			i.consolidation[key] = existing
		} else {
			i.consolidation[key] = concCred
		}
	}
}

// AddEstablishments processes a slice of establishments and updates the ConcCred instance with the number of accredited and active establishments.
func (i *Conccred) AddEstablishments(year int, quarter int, establishments []*source_domain.Establishment) {
	// get new ConcCred instances for each establishment and aggregate the number of accredited and active establishments by brand and function
	estabItems := make(map[string]*ConcCredItem)
	for _, e := range establishments {
		if !e.IsAccredited(year, quarter) {
			continue
		}
		Conccred := i.conccred.GetFromEstablishment(year, quarter, e)
		for _, concCred := range Conccred {
			key := concCred.GetKey()
			if existing, exists := estabItems[key]; exists {
				existing.NumberAccreditedEstablishments += concCred.NumberAccreditedEstablishments
				existing.NumberActiveEstablishments += concCred.NumberActiveEstablishments
				estabItems[key] = existing
			} else {
				estabItems[key] = concCred
			}
		}
	}
	// merge the establishment data into the main items map
	for key, estabConcCred := range estabItems {
		if existing, exists := i.consolidation[key]; exists {
			existing.NumberAccreditedEstablishments = estabConcCred.NumberAccreditedEstablishments
			// workaround to set the number of active establishments to 76% of the accredited establishments
			// this is a temporary solution until we have the actual number of active establishments from the transactions data
			existing.NumberActiveEstablishments = int64(float64(estabConcCred.NumberAccreditedEstablishments) * EstimatedActiveEstablishmentRatio1)
			i.consolidation[key] = existing
		} else {
			i.consolidation[key] = estabConcCred
		}
	}
}

// AddTerminals processes a slice of terminals and updates the ConcCred instance with the number of accredited and active terminals.
func (i *Conccred) AddTerminals(year int, quarter int, terminals []*source_domain.Terminal, esblishmentMap map[int64]string) {
}