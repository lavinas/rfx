package target_domain

import (
	"time"
	"fmt"
	"math"

	source_domain "consolidator/internal/core/domain/source"
)

// Desconto represents the data structure for discounts which will be used for fusing data between intercam, management and webservice
type Desconto struct {
	ID                  int64     `gorm:"column:id"`
	CreatedAt           time.Time `gorm:"column:created_at"`
	UpdatedAt           time.Time `gorm:"column:updated_at"`
	Year                int       `gorm:"column:year"`
	Quarter             int       `gorm:"column:quarter"`
	Function            string    `gorm:"column:function"`
	Brand               int       `gorm:"column:brand"`
	CaptureMode         int       `gorm:"column:capture_mode"`
	Installments        int       `gorm:"column:installments"`
	SegmentCode         int       `gorm:"column:segment_code"`
	AvgMDRFee           float64   `gorm:"column:avg_mdr_fee"`
	MinMDRFee           float64   `gorm:"column:min_mdr_fee"`
	MaxMDRFee           float64   `gorm:"column:max_mdr_fee"`
	StdevMDRFee         float64   `gorm:"column:stdev_mdr_fee"`
	SqrdiffMDRFee       float64   `gorm:"column:sqrdiff_mdr_fee"`
	TransactionAmount   float64   `gorm:"column:transaction_amount"`
	TransactionQuantity int64     `gorm:"column:transaction_quantity"`
}

// TableName specifies the table name for Desconto struct
func (i *Desconto) TableName() string {
	return "cadoc_6334_v2.desconto"
}

// GetFromTransaction returns the discount data for a given transaction.
func (i *Desconto) GetFromTransaction(transaction *source_domain.Transaction) *Desconto {
	return &Desconto{
		Year:                transaction.GetYear(),
		Quarter:             transaction.GetQuarter(),
		Function:            transaction.GetFunctionCode(),
		Brand:               transaction.GetBrandCode(),
		CaptureMode:         transaction.GetCaptureModeCode(),
		Installments:        transaction.GetInstallments(),
		SegmentCode:         transaction.GetSegmentCode(),
		AvgMDRFee:           transaction.GetRevenueMDRValue(),
		MinMDRFee:           transaction.GetRevenueMDRValue(),
		MaxMDRFee:           transaction.GetRevenueMDRValue(),
		StdevMDRFee:         0,
		SqrdiffMDRFee:       0,
		TransactionAmount:   transaction.GetTransactionAmount(),
		TransactionQuantity: 1,
	}
}

// GetKey generates a unique key for the Desconto struct based on its fields.
func (i *Desconto) GetKey() string {
	return fmt.Sprintf("%d-%d-%s-%d-%d-%d-%d", i.Year, i.Quarter, i.Function, i.Brand, i.CaptureMode, i.Installments, i.SegmentCode)
}

// GetFromTransactions returns a map of Desconto structs for a given list of transactions.
func (i *Desconto) AddTransactions(transactions []*source_domain.Transaction, items map[string]*Desconto) {
	for _, t := range transactions {
		desconto := i.GetFromTransaction(t)
		key := desconto.GetKey()
		if existing, exists := items[key]; exists {
			// sum amount and quantity
			existing.TransactionAmount += desconto.TransactionAmount
			existing.TransactionQuantity += desconto.TransactionQuantity
			// update min and max MDR fees
			if desconto.MinMDRFee < existing.MinMDRFee {
				existing.MinMDRFee = desconto.MinMDRFee
			}
			if desconto.MaxMDRFee > existing.MaxMDRFee {
				existing.MaxMDRFee = desconto.MaxMDRFee
			}
			// update average and standard deviation using Welford's online algorithm
			delta := desconto.AvgMDRFee - existing.AvgMDRFee
			existing.AvgMDRFee += delta / float64(existing.TransactionQuantity)
			delta2 := desconto.StdevMDRFee - existing.StdevMDRFee
			existing.SqrdiffMDRFee += delta2 * delta2
			variance := existing.SqrdiffMDRFee / (float64(existing.TransactionQuantity) - 1)
			existing.StdevMDRFee = math.Sqrt(variance)
			items[key] = existing
		} else {
			items[key] = desconto
		}
	}
}