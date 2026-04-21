package target_domain

import (
	"fmt"
	"time"

	source_domain "consolidator/internal/core/domain/source"
)

const (
	EstimatedActiveEstablishmentRatio = 0.76 // Estimated ratio of active establishments to accredited establishments
)

// ConcCred represents the consolidated credit card transactions for a specific year, quarter, brand, and function.
type ConcCred struct {
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

// NewConcCred creates a new instance of ConcCred.
func NewConcCred() *ConcCred {
	return &ConcCred{}
}

// TableName specifies the table name for ConcCred struct
func (i *ConcCred) TableName() string {
	return "cadoc_6334_v2.conccred"
}

// GetKey generates a unique key for the ConcCred struct based on its fields.
func (i *ConcCred) GetKey() string {
	return fmt.Sprintf("%d-%d-%s-%d", i.Year, i.Quarter, i.Function, i.Brand)
}

// AddTransactions processes a slice of transactions and updates the ConcCred instance accordingly.
func (i *ConcCred) AddTransactions(transactions []*source_domain.Transaction, items map[string]*ConcCred) {
	// for each transaction, get the corresponding ConcCred instance and update the transaction amount, quantity and establishment counts
	for _, t := range transactions {
		concCred := i.GetFromTransaction(t)
		key := concCred.GetKey()

		// if the key already exists in the items map, update the existing ConcCred instance with the new transaction data
		if existing, exists := items[key]; exists {
			existing.TransactionAmount += concCred.TransactionAmount
			existing.TransactionQuantity += concCred.TransactionQuantity
			items[key] = existing
		} else {
			items[key] = concCred
		}
	}
}

// AddEstablishments processes a slice of establishments and updates the ConcCred instance with the number of accredited and active establishments.
func (i *ConcCred) AddEstablishments(year int, quarter int, establishments []*source_domain.Establishment, items map[string]*ConcCred) {
	// get new ConcCred instances for each establishment and aggregate the number of accredited and active establishments by brand and function
	estabItems := make(map[string]*ConcCred)
	for _, e := range establishments {
		if !e.IsAccredited(year, quarter) {
			continue
		}
		concCreds := i.GetFromEstablishment(year, quarter, e)
		for _, concCred := range concCreds {
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
		if existing, exists := items[key]; exists {
			existing.NumberAccreditedEstablishments = estabConcCred.NumberAccreditedEstablishments
			// workaround to set the number of active establishments to 76% of the accredited establishments
			// this is a temporary solution until we have the actual number of active establishments from the transactions data
			existing.NumberActiveEstablishments = int64(float64(estabConcCred.NumberAccreditedEstablishments) * EstimatedActiveEstablishmentRatio)
			items[key] = existing
		} else {
			items[key] = estabConcCred
		}
	}
}

// GetFromTransaction returns a ConcCred instance populated with data from a given transaction.
func (i *ConcCred) GetFromTransaction(transaction *source_domain.Transaction) *ConcCred {
	return &ConcCred{
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
func (i *ConcCred) GetFromEstablishment(year int, quarter int, establishment *source_domain.Establishment) []*ConcCred {
	// Get the brand codes and function codes for the establishment
	brands := establishment.GetBrandCodes()
	functions := establishment.GetFunctionCodes()

	var concCreds []*ConcCred

	// Determine the number of active establishments based on the establishment's status for the given year and quarter
	activeCount := int64(0)
	if establishment.IsActive(year, quarter) {
		activeCount = 1
	}

	// Create a ConcCred instance for each combination of brand and function, populating the establishment counts accordingly
	for _, function := range functions {
		for _, brand := range brands {
			concCreds = append(concCreds, &ConcCred{
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
