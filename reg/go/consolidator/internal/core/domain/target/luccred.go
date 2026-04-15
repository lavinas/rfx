package target_domain

import (
	"time"
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
