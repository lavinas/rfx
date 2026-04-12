package domain

import (
	"time"
)

// ConcCred represents the consolidated credit card transactions for a specific year, quarter, brand, and function.
type ConcCred struct {
	ID                             int       `gorm:"column:id;primaryKey"`
	CreatedAt                      time.Time `gorm:"column:created_at"`
	UpdatedAt                      time.Time `gorm:"column:updated_at"`
	Year                           int       `gorm:"column:year"`
	Quarter                        int       `gorm:"column:quarter"`
	Brand                          int       `gorm:"column:brand"`
	Function                       string    `gorm:"column:function"`
	TransactionAmount              float64   `gorm:"column:transaction_amount"`
	TransactionQuantity            int64     `gorm:"column:transaction_quantity"`
	NumberAccreditedEstablishments int64     `gorm:"column:number_accredited_establishments"`
	NumberActiveEstablishments     int64     `gorm:"column:number_active_establishments"`
}

// TableName specifies the table name for ConcCred struct
func (ConcCred) TableName() string {
	return "cadoc_6334_v2.conccred"
}
