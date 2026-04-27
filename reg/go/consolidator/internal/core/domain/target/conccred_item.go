package target_domain

import (
	"time"
	"fmt"

	source_domain "consolidator/internal/core/domain/source"

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

// NewConcCredItem creates a new instance of ConcCredItem.
func NewConcCredItem() *ConcCredItem {
	return &ConcCredItem{}
}

// TableName specifies the table name for ConcCredItem struct
func (i *ConcCredItem) TableName() string {
	return "cadoc_6334_v2.conccred"
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
