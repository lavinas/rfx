package target_domain

import (
	"time"
	"fmt"

	source_domain "consolidator/internal/core/domain/source"
)


// DescontoItem represents the data structure for discounts which will be used for fusing data between intercam, management and webservice
type DescontoItem struct {
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


// TableName specifies the table name for DescontoItem struct
func (i *DescontoItem) TableName() string {
	return "cadoc_6334_v2.desconto"
}

// NewDescontoItem creates a new instance of DescontoItem.
func NewDescontoItem() *DescontoItem {
	return &DescontoItem{}
}

// GetKey generates a unique key for the DescontoItem struct based on its fields.
func (i *DescontoItem) GetKey() string {
	return fmt.Sprintf("%d-%d-%s-%d-%d-%d-%d", i.Year, i.Quarter, i.Function, i.Brand, i.CaptureMode, i.Installments, i.SegmentCode)
}

// GetFromTransaction returns the discount data for a given transaction.
func (i *DescontoItem) GetFromTransaction(transaction *source_domain.Transaction) *DescontoItem {
	return &DescontoItem{
		Year:                transaction.GetYear(),
		Quarter:             transaction.GetQuarter(),
		Function:            transaction.GetFunctionCode(),
		Brand:               transaction.GetBrandCode(),
		CaptureMode:         transaction.GetCaptureModeCode(),
		Installments:        transaction.GetInstallments(),
		SegmentCode:         transaction.GetSegmentCode(),
		AvgMDRFee:           transaction.GetRevenueMDRValueRate(),
		MinMDRFee:           transaction.GetRevenueMDRValueRate(),
		MaxMDRFee:           transaction.GetRevenueMDRValueRate(),
		StdevMDRFee:         0,
		SqrdiffMDRFee:       0,
		TransactionAmount:   transaction.GetTransactionAmount(),
		TransactionQuantity: 1,
	}
}
