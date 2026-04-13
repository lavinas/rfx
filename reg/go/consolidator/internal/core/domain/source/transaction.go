package source_domain

import (
	"time"
)

// Transaction represents the data structure for transactions which will be used for fusing data between intercam, management and webservice
type Transaction struct {
	ID                          int64      `gorm:"column:id"`
	CreatedAt                   time.Time  `gorm:"column:created_at;type:timestamp"`
	UpdatedAt                   time.Time  `gorm:"column:updated_at;type:timestamp"`
	Key1                        string     `gorm:"column:key1"`
	Key2                        *string    `gorm:"column:key2"`
	EstablishmentCode           *int64     `gorm:"column:establishment_code"`
	EstablishmentNature         *int64     `gorm:"column:establishment_nature"`
	EstablishmentMCC            *int64     `gorm:"column:establishment_mcc"`
	EstablishmentTerminalCode   *int64     `gorm:"column:establishment_terminal_code"`
	BIN                         *int64     `gorm:"column:bin"`
	AuthorizationCode           *string    `gorm:"column:authorization_code"`
	TransactionNSU              *string    `gorm:"column:transaction_nsu"`
	TransactionDate             *time.Time `gorm:"column:transaction_date;type:timestamp"`
	TransactionSecondaryDate    *time.Time `gorm:"column:transaction_secondary_date;type:timestamp"`
	TransactionAmount           *float64   `gorm:"column:transaction_amount"`
	TransactionInstallments     *int64     `gorm:"column:transaction_installments"`
	TransactionInstallmentsType *string    `gorm:"column:transaction_installments_type"`
	TransactionBrand            *string    `gorm:"column:transaction_brand"`
	TransactionProduct          *string    `gorm:"column:transaction_product"`
	TransactionCapture          *string    `gorm:"column:transaction_capture"`
	RevenueMDRValue             *float64   `gorm:"column:revenue_mdr_value"`
	CostInterchangeValue        *float64   `gorm:"column:cost_interchange_value"`
	HighSourcePriority          *int64     `gorm:"column:high_source_priority"`
	StatusID                    *int64     `gorm:"column:status_id"`
	StatusName                  *string    `gorm:"column:status_name"`
	StatusCount                 int64      `gorm:"column:status_count"`
	PeriodDate                  *time.Time `gorm:"column:period_date;type:timestamp"`
	PeriodClosingID             *int64     `gorm:"column:period_closing_id"`
	TransacID                   *string    `gorm:"column:transac_id"`
}

// TableName specifies the table name for Transaction struct
func (Transaction) TableName() string {
	return "transaction_v2.transaction"
}

// GetYear returns the year of the transaction based on the PeriodDate field.
func (t *Transaction) GetYear() int {
	if t.PeriodDate != nil {
		return t.PeriodDate.Year()
	}
	return 0
}

// GetQuarter returns the quarter of the transaction based on the PeriodDate field.
func (t *Transaction) GetQuarter() int {
	if t.PeriodDate != nil {
		return (int(t.PeriodDate.Month())-1)/3 + 1
	}
	return 0
}

// GetEstablishmentCode returns the establishment code of the transaction based on the EstablishmentCode field.
func (t *Transaction) GetEstablishmentCode() int64 {
	if t.EstablishmentCode != nil {
		return *t.EstablishmentCode
	}
	return 0
}

// GetBin returns the Bank Identification Number (BIN) of the transaction based on the BIN field.
func (t *Transaction) GetBin() int64 {
	if t.BIN != nil {
		return *t.BIN
	}
	return 0
}

// GetFunctionCode returns the function code of the transaction based on the TransactionProduct field.
func (t *Transaction) GetFunctionCode() string {
	if t.TransactionProduct != nil {
		if productCode, ok := ProductMap[*t.TransactionProduct]; ok {
			return productCode
		}
		return ""
	}
	return ""
}

// GetBrandCode returns the brand code of the transaction based on the TransactionBrand field.
func (t *Transaction) GetBrandCode() int {
	if t.TransactionBrand != nil {
		if brand, ok := BrandMap[*t.TransactionBrand]; ok {
			return brand
		}
		return 0
	}
	return 0
}

// GetCaptureModeCode returns the capture mode code of the transaction based on the TransactionCapture field.
func (t *Transaction) GetCaptureModeCode() int {
	if t.TransactionCapture != nil {
		if captureMode, ok := CaptureModeMap[*t.TransactionCapture]; ok {
			return captureMode
		}
		return 0
	}
	return 0
}

// GetInstallmentsTypeCode returns the installments type code of the transaction based on the TransactionInstallmentsType field.
func (t *Transaction) GetInstallments() int {

	if t.TransactionInstallments == nil {
		return 0
	}
	if *t.TransactionInstallments == 0 {
		return 1 // Assuming 1 represents a specific installments type when the value is 0
	}
	return int(*t.TransactionInstallments)
}

// GetSegmentCode returns the segment code of the transaction based on the EstablishmentNature field.
func (t *Transaction) GetSegmentCode() int {
	if t.EstablishmentMCC == nil {
		return 0
	}
	if *t.EstablishmentMCC >= 3000 && *t.EstablishmentMCC <= 3350 {
		return 403
	}
	if *t.EstablishmentMCC >= 3501 && *t.EstablishmentMCC <= 3999 {
		return 418
	}
	if *t.EstablishmentMCC >= 3351 && *t.EstablishmentMCC <= 3500 {
		return 418
	}
	if segment, ok := SegmentMap[int(*t.EstablishmentMCC)]; ok {
		return segment
	}
	return 0
}

// GetTransactionAmount returns the transaction amount of the transaction based on the TransactionAmount field.
func (t *Transaction) GetTransactionAmount() float64 {
	if t.TransactionAmount != nil {
		return *t.TransactionAmount
	}
	return 0
}

// GetInterchangeFee returns the interchange fee of the transaction based on the CostInterchangeValue field.
func (t *Transaction) GetInterchangeFee() float64 {
	if t.CostInterchangeValue != nil {
		return *t.CostInterchangeValue
	}
	return 0
}

// GetRevenueMDRValue returns the revenue MDR value of the transaction based on the RevenueMDRValue field.
func (t *Transaction) GetRevenueMDRValue() float64 {
	if t.RevenueMDRValue != nil {
		return *t.RevenueMDRValue
	}
	return 0
}
