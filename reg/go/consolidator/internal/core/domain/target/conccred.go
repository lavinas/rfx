package target_domain

import (
	"fmt"
	"maps"
	"slices"
	"time"

	source_domain "consolidator/internal/core/domain/source"
	"consolidator/internal/core/ports"
)

const (
	EstimatedActiveEstablishmentRatio1 = 0.76 // Estimated ratio of active establishments to accredited establishments
)

// ConcCred represents the consolidated credit card transactions for a specific year, quarter, brand, and function.
type ConcCredItem struct {
	ID                             int       `gorm:"column:id;primaryKey"`
	CreatedAt                      time.Time `gorm:"column:created_at"`
	UpdatedAt                      time.Time `gorm:"column:updated_at"`
	Year                           int       `gorm:"column:year"`
	Quarter                        int       `gorm:"column:quarter"`
	Function                       string    `gorm:"column:function"`
	Brand                          int       `gorm:"column:brand"`
	NumberAccreditedEstablishments int64     `gorm:"column:number_accredited_establishments"`
	NumberActiveEstablishments     int64     `gorm:"column:number_active_establishments"`
	TransactionAmount              float64   `gorm:"column:transaction_amount"`
	TransactionQuantity            int64     `gorm:"column:transaction_quantity"`
}

// Conccred represents the consolidated credit card transactions and establishment data for a specific year and quarter.
type Conccred struct {
	DomainBase
	conccred      *ConcCredItem
	consolidation map[string]*ConcCredItem
}

// NewConccred creates a new instance of Conccred.
func NewConccred() *Conccred {
	return &Conccred{
		conccred:      &ConcCredItem{},
		consolidation: make(map[string]*ConcCredItem),
	}
}

// TableName specifies the table name for ConcCredItem struct
func (i *ConcCredItem) TableName() string {
	return "conccred"
}

// GetKey generates a unique key for the ConcCred struct based on its fields.
func (i *ConcCredItem) GetKey() string {
	return fmt.Sprintf("%d-%d-%s-%d", i.Year, i.Quarter, i.Function, i.Brand)
}

// GetFromTransaction returns a ConcCredItem instance populated with data from a given transaction.
func (i *ConcCredItem) GetFromTransaction(transaction *source_domain.Transaction) *ConcCredItem {
	return &ConcCredItem{
		Year:                           transaction.GetYear(),
		Quarter:                        transaction.GetQuarter(),
		Brand:                          transaction.GetBrandCode(),
		Function:                       transaction.GetFunctionCode(),
		NumberAccreditedEstablishments: 0,
		NumberActiveEstablishments:     0,
		TransactionAmount:              transaction.GetTransactionAmount(),
		TransactionQuantity:            1,
	}
}

// GetFromEstablishment returns a ConcCred instance populated with data from a given establishment.
func (i *ConcCredItem) GetFromEstablishment(year int, quarter int, establishment *source_domain.Establishment) []*ConcCredItem {
	// Get the brand codes and function codes for the establishment
	brands := establishment.GetBrandCodes()
	functions := establishment.GetFunctionCodes()

	var concCreds []*ConcCredItem

	// Determine the number of active establishments based on the establishment's status for the given year and quarter
	activeCount := int64(0)
	if establishment.IsActive(year, quarter) {
		activeCount = 1
	}

	// Create a ConcCred instance for each combination of brand and function, populating the establishment counts accordingly
	for _, function := range functions {
		for _, brand := range brands {
			concCreds = append(concCreds, &ConcCredItem{
				Year:                           year,
				Quarter:                        quarter,
				Brand:                          brand,
				Function:                       function,
				NumberAccreditedEstablishments: 1,
				NumberActiveEstablishments:     activeCount,
			})
		}
	}

	return concCreds
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
func (i *Conccred) Save(repository ports.Repository) error {
	if err := repository.Save(slices.Collect(maps.Values(i.consolidation))); err != nil {
		return err
	}
	return nil
}

// AddTransactions processes a slice of transactions and updates the ConcCred instance with the total transaction amount and quantity.
func (i *Conccred) AddTransactions(transactions []*source_domain.Transaction) {
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
