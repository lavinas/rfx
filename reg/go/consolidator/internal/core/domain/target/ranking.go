package target_domain

import (
	"time"
	"fmt"

	source_domain "consolidator/internal/core/domain/source"
)

// Ranking represents the ranking of establishments based on transaction data.
type Ranking struct {
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

// TableName returns the name of the database table for the Ranking model.
func (i *Ranking) TableName() string {
	return "cadoc_6334_v2.ranking"
}

// GetFromTransaction returns a Ranking instance populated with data from a given transaction.
func (i *Ranking) GetFromTransaction(transaction *source_domain.Transaction) *Ranking {
	return &Ranking{
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
		AvgMccFee:           transaction.GetRevenueMDRValue(),
	}
}

// GetKey generates a unique key for the Ranking struct based on its fields.
func (i *Ranking) GetKey() string {
	return fmt.Sprintf("%d-%d-%d-%s-%d-%d-%d-%d", i.Year, i.Quarter, i.EstablishmentCode, i.Function, i.Brand, i.CaptureMode, i.Installments, i.SegmentCode)
}

// GetFromTransactions processes a slice of transactions and returns a map of Ranking instances keyed by their unique keys.
func (i *Ranking) AddTransactions(transactions []*source_domain.Transaction, items map[string]*Ranking) {
	for _, t := range transactions {
		ranking := i.GetFromTransaction(t)
		key := ranking.GetKey()
		if existing, exists := items[key]; exists {
			existing.TransactionAmount += ranking.TransactionAmount
			existing.TransactionQuantity += ranking.TransactionQuantity
			delta := ranking.AvgMccFee - existing.AvgMccFee
			existing.AvgMccFee += delta / float64(existing.TransactionQuantity)
		} else {
			items[key] = ranking
		}
	}
}