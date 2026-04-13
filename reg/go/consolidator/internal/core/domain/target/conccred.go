package target_domain

import (
	"time"
	"fmt"

	source_domain "consolidator/internal/core/domain/source"
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

// TableName specifies the table name for ConcCred struct
func (i *ConcCred) TableName() string {
	return "cadoc_6334_v2.conccred"
}

// GetFromTransaction returns a ConcCred instance populated with data from a given transaction.
func (i *ConcCred) GetFromTransaction(transaction *source_domain.Transaction) *ConcCred {
	return &ConcCred{
		Year:                           transaction.GetYear(),
		Quarter:                        transaction.GetQuarter(),
		Brand:                          transaction.GetBrandCode(),
		Function:                       transaction.GetFunctionCode(),
		NumberAccreditedEstablishments: 0, // This would require additional logic to determine the number of accredited establishments
		NumberActiveEstablishments:     0, // This would require additional logic to determine the number of active establishments
		TransactionAmount:              transaction.GetTransactionAmount(),
		TransactionQuantity:            1,
	}
}

// GetKey generates a unique key for the ConcCred struct based on its fields.
func (i *ConcCred) GetKey() string {
	return fmt.Sprintf("%d-%d-%s-%d", i.Year, i.Quarter, i.Function, i.Brand)
}

// AddTransactions processes a slice of transactions and updates the ConcCred instance accordingly.
func (i *ConcCred) AddTransactions(transactions []*source_domain.Transaction, items map[string]*ConcCred) {
	for _, t := range transactions {
		concCred := i.GetFromTransaction(t)
		key := concCred.GetKey()
		if existing, exists := items[key]; exists {
			existing.TransactionAmount += concCred.TransactionAmount
			existing.TransactionQuantity += concCred.TransactionQuantity
			items[key] = existing
		} else {
			items[key] = concCred
		}
	}
}

// GetFromEstablishment returns a ConcCred instance populated with data from a given establishment.
func (i *ConcCred) GetFromEstablishment(year int, quarter int, establishment *source_domain.Establishment) []*ConcCred {
	var concCreds []*ConcCred
	for _, functions := range establishment.GetFunctionCodes() {
		for _, brands := range establishment.GetBrandCodes() {
				concCred := &ConcCred{
					Year:                           year,
					Quarter:                        quarter,
					Brand:                          brands,
					Function:                       functions,
					NumberAccreditedEstablishments: int64(establishment.IsAccredited(year, quarter)),
					NumberActiveEstablishments:     int64(establishment.IsActive(year, quarter)),
					TransactionAmount:              0, // This would require additional logic to determine the transaction amount for the establishment
					TransactionQuantity:            0, // This would require additional logic to determine the transaction quantity for the establishment
				}
				concCreds = append(concCreds, concCred)
		}
	}
	return concCreds
}

// AddEstablishments processes a slice of establishments and updates the ConcCred instance accordingly.
func (i *ConcCred) AddEstablishments(establishments []*source_domain.Establishment, items map[string]*ConcCred) {
	for _, e := range establishments {
		concCreds := i.GetFromEstablishment(i.Year, i.Quarter, e)
		for _, concCred := range concCreds {		
			key := concCred.GetKey()
			if existing, exists := items[key]; exists {
				existing.NumberAccreditedEstablishments += concCred.NumberAccreditedEstablishments
				existing.NumberActiveEstablishments += concCred.NumberActiveEstablishments
				items[key] = existing
			} else {
				items[key] = concCred
			}
		}
	}
}

