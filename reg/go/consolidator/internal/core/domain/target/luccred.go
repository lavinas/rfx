package target_domain

import (
	"fmt"
	"time"

	source_domain "consolidator/internal/core/domain/source"
)

// Luccred represents the financial data for a specific year and quarter.
type Luccred struct {
	ID              int       `gorm:"column:id"`
	CreatedAt       time.Time `gorm:"column:created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at"`
	Year            int       `gorm:"column:year"`
	Quarter         int       `gorm:"column:quarter"`
	GrossRevenue    float64   `gorm:"column:gross_revenue"`
	RentalRevenue   float64   `gorm:"column:rental_revenue"`
	OthersRevenue   float64   `gorm:"column:others_revenue"`
	InterchangeCost float64   `gorm:"column:interchange_cost"`
	MarketingCost   float64   `gorm:"column:marketing_cost"`
	BrandAccessCost float64   `gorm:"column:brand_access_cost"`
	RiskCost        float64   `gorm:"column:risk_cost"`
	ProcessingCost  float64   `gorm:"column:processing_cost"`
	OthersCost      float64   `gorm:"column:others_cost"`
}

// NewLuccred creates a new instance of Luccred.
func NewLuccred() *Luccred {
	return &Luccred{}
}

// TableName returns the name of the database table for the Luccred model.
func (i *Luccred) TableName() string {
	return "cadoc_6334_v2.luccred"
}

// GetFromTransaction returns a ConcCred instance populated with data from a given transaction.
func (i *Luccred) GetFromTransaction(transaction *source_domain.Transaction) *Luccred {
	return &Luccred{
		Year:            transaction.GetYear(),
		Quarter:         transaction.GetQuarter(),
		GrossRevenue:    transaction.GetRevenueMDRValue(),
		RentalRevenue:   0,
		OthersRevenue:   0,
		InterchangeCost: transaction.GetInterchangeFee(),
		MarketingCost:   0,
		BrandAccessCost: 0,
		RiskCost:        0,
		ProcessingCost:  0,
		OthersCost:      0,
	}
}

// GetKey generates a unique key for the Luccred struct based on its fields.
func (i *Luccred) GetKey() string {
	return fmt.Sprintf("%d-%d", i.Year, i.Quarter)
}

// AddTransactions processes a slice of transactions and updates the Luccred instance accordingly.
func (i *Luccred) AddTransactions(transactions []*source_domain.Transaction, items map[string]*Luccred) {
	// for each transaction, get the corresponding Luccred instance and update the transaction amount, quantity and establishment counts
	for _, t := range transactions {
		luccred := i.GetFromTransaction(t)
		key := luccred.GetKey()

		// if the key already exists in the items map, update the existing Luccred instance with the new transaction data
		if existing, exists := items[key]; exists {
			existing.GrossRevenue += luccred.GrossRevenue
			existing.InterchangeCost += luccred.InterchangeCost
			items[key] = existing
		} else {
			items[key] = luccred
		}
	}
}
