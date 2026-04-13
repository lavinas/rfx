package target_domain

import (
	"time"
)

// Ranking represents the ranking of establishments based on transaction data.
type Ranking struct {
	ID                  int64     `gorm:"column:id"`
	CreatedAt           time.Time `gorm:"column:created_at"`
	UpdatedAt           time.Time `gorm:"column:updated_at"`
	Year                int16     `gorm:"column:year"`
	Quarter             int16     `gorm:"column:quarter"`
	EstablishmentCode   int64     `gorm:"column:establishment_code"`
	Function            string    `gorm:"column:function"`
	Brand               int16     `gorm:"column:brand"`
	CaptureMode         int16     `gorm:"column:capture_mode"`
	Installments        int16     `gorm:"column:installments"`
	SegmentCode         int16     `gorm:"column:segment_code"`
	TransactionAmount   float64   `gorm:"column:transaction_amount"`
	TransactionQuantity int       `gorm:"column:transaction_quantity"`
	AvgMccFee           float64   `gorm:"column:avg_mcc_fee"`
}

// TableName returns the name of the database table for the Ranking model.
func (Ranking) TableName() string {
	return "cadoc_6334_v2.ranking"
}
