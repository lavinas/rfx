package target_domain

import (
	"time"
)

// Desconto represents the data structure for discounts which will be used for fusing data between intercam, management and webservice
type Desconto struct {
	ID                  int64     `gorm:"column:id"`
	CreatedAt           time.Time `gorm:"column:created_at"`
	UpdatedAt           time.Time `gorm:"column:updated_at"`
	Year                int64     `gorm:"column:year"`
	Quarter             int64     `gorm:"column:quarter"`
	Function            string    `gorm:"column:function"`
	Brand               int64     `gorm:"column:brand"`
	FormaCaptura        int64     `gorm:"column:forma_captura"`
	NumeroParcelas      int64     `gorm:"column:numero_parcelas"`
	SegmentCode         int64     `gorm:"column:segment_code"`
	TransactionAmount   float64   `gorm:"column:transaction_amount"`
	TransactionQuantity int64     `gorm:"column:transaction_quantity"`
	AvgMCCFee           float64   `gorm:"column:avg_mcc_fee"`
	MinMCCFee           float64   `gorm:"column:min_mcc_fee"`
	MaxMCCFee           float64   `gorm:"column:max_mcc_fee"`
	StdevMCCFee         float64   `gorm:"column:stdev_mcc_fee"`
	SqrdiffMCCFee       float64   `gorm:"column:sqrdiff_mcc_fee"`
	CaptureMode         int64     `gorm:"column:capture_mode"`
	Installments        int64     `gorm:"column:installments"`
}

// TableName specifies the table name for Desconto struct
func (Desconto) TableName() string {
	return "cadoc_6334_v2.desconto"
}
