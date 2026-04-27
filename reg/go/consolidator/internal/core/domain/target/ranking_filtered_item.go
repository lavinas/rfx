package target_domain

import (
	"fmt"
	"time"
)

// RankingFilteredItem represents the filtered ranking of establishments based on transaction data.
type RankingFilteredItem struct {
	ID                  int64     `gorm:"column:id"`
	CreatedAt           time.Time `gorm:"column:created_at"`
	UpdatedAt           time.Time `gorm:"column:updated_at"`
	Year                int       `gorm:"column:year"`
	Quarter             int       `gorm:"column:quarter"`
	EstablishmentCode   int64     `gorm:"column:establishment_code"`
	Function            string    `gorm:"column:function"`
	Brand               int       `gorm:"column:brand"`
	CaptureMode         int       `gorm:"column:capture_mode"`
	Installments        int       `gorm:"column:installments"`
	SegmentCode         int       `gorm:"column:segment_code"`
	TransactionAmount   float64   `gorm:"column:transaction_amount"`
	TransactionQuantity int       `gorm:"column:transaction_quantity"`
	AvgMccFee           float64   `gorm:"column:avg_mcc_fee"`
}

// NewRankingFilteredItem creates a new instance of RankingFilteredItem.
func NewRankingFilteredItem() *RankingFilteredItem {
	return &RankingFilteredItem{}
}

// GetKey generates a unique key for the RankingFilteredItem struct based on its fields.
func (i *RankingFilteredItem) GetKey() string {
	return fmt.Sprintf("%d-%d-%d-%s-%d-%d-%d-%d", i.Year, i.Quarter, i.EstablishmentCode, i.Function, i.Brand, i.CaptureMode, i.Installments, i.SegmentCode)
}

// TableName returns the name of the database table for the RankingFilteredItem model.
func (i *RankingFilteredItem) TableName() string {
	return "cadoc_6334_v2.ranking_filtered"
}
