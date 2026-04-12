package domain

import (
	"time"
)

// Interchange represents the interchange fee for a given transaction.
type Interchange struct {
	ID                  int64     `gorm:"column:id"`
	CreatedAt           time.Time `gorm:"column:created_at"`
	UpdatedAt           time.Time `gorm:"column:updated_at"`
	Year                int       `gorm:"column:year"`
	Quarter             int       `gorm:"column:quarter"`
	ProductCode         int       `gorm:"column:product_code"`
	CardType            string    `gorm:"column:card_type"`
	Function            string    `gorm:"column:function"`
	Brand               int       `gorm:"column:brand"`
	CaptureMode         int       `gorm:"column:capture_mode"`
	Installments        int       `gorm:"column:installments"`
	SegmentCode         string    `gorm:"column:segment_code"`
	InterchangeFee      float64   `gorm:"column:interchange_fee"`
	TransactionAmount   float64   `gorm:"column:transaction_amount"`
	TransactionQuantity int64     `gorm:"column:transaction_quantity"`
}

// TableName returns the name of the table in the database.
func (Interchange) TableName() string {
	return "cadoc_6334_v2.interchange"
}
