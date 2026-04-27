package target_domain

import (
	"fmt"
	"time"

	source_domain "consolidator/internal/core/domain/source"
)

// RankingItem represents the ranking of establishments based on transaction data.
type RankingItem struct {
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

// NewRankingItem creates a new instance of RankingItem.
func NewRankingItem() *RankingItem {
	return &RankingItem{}
}

// TableName returns the name of the database table for the RankingItem model.
func (i *RankingItem) TableName() string {
	return "cadoc_6334_v2.ranking"
}

// GetFromTransaction returns a RankingItem instance populated with data from a given transaction.
func (i *RankingItem) GetFromTransaction(transaction *source_domain.Transaction) *RankingItem {
	return &RankingItem{
		Year:                transaction.GetYear(),
		Quarter:             transaction.GetQuarter(),
		EstablishmentCode:   transaction.GetEstablishmentCode(),
		Function:            transaction.GetFunctionCode(),
		Brand:               transaction.GetBrandCode(),
		CaptureMode:         transaction.GetCaptureModeCode(),
		Installments:        transaction.GetInstallments(),
		SegmentCode:         transaction.GetSegmentCode(),
		TransactionAmount:   transaction.GetTransactionAmount(),
		TransactionQuantity: 1,
		AvgMccFee:           transaction.GetRevenueMDRValueRate(),
	}
}

// GetKey generates a unique key for the RankingItem struct based on its fields.
func (i *RankingItem) GetKey() string {
	return fmt.Sprintf("%d-%d-%d-%s-%d-%d-%d-%d", i.Year, i.Quarter, i.EstablishmentCode, i.Function, i.Brand, i.CaptureMode, i.Installments, i.SegmentCode)
}