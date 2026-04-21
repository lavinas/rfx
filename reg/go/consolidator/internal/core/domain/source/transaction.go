package source_domain

import (
	"time"
)

const (
	DefaultSegmentCode = 423 // Default segment code for transactions that do not match any specific segment
)

// Transaction represents the data structure for transactions which will be used for fusing data between intercam, management and webservice
type Transaction struct {
	ID                      int64      `gorm:"column:id"`
	EstablishmentCode       *int64     `gorm:"column:establishment_code"`
	EstablishmentMCC        *int64     `gorm:"column:establishment_mcc"`
	BIN                     *int64     `gorm:"column:bin"`
	TransactionAmount       *float64   `gorm:"column:transaction_amount"`
	TransactionInstallments *int64     `gorm:"column:transaction_installments"`
	TransactionBrand        *string    `gorm:"column:transaction_brand"`
	TransactionProduct      *string    `gorm:"column:transaction_product"`
	TransactionCapture      *string    `gorm:"column:transaction_capture"`
	RevenueMDRValue         *float64   `gorm:"column:revenue_mdr_value"`
	CostInterchangeValue    *float64   `gorm:"column:cost_interchange_value"`
	PeriodDate              *time.Time `gorm:"column:period_date;type:timestamp"`
}

// TableName specifies the table name for Transaction struct
func (Transaction) TableName() string {
	return "transaction_v4.transaction"
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

// GetMccCode returns the Merchant Category Code (MCC) of the transaction based on the EstablishmentMCC field.
func (t *Transaction) GetMccCode() int {
	if t.EstablishmentMCC != nil {
		return int(*t.EstablishmentMCC)
	}
	return 0
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
	// If the EstablishmentMCC does not match any of the specified ranges or mappings, return the default segment code.
	return DefaultSegmentCode
}

// GetSegmentName returns the segment name of the transaction based on the EstablishmentMCC field.
func (t *Transaction) GetSegmentName() string {
	code := t.GetSegmentCode()
	if code == 0 {
		return ""
	}
	return SegmentNameMap[code]
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

// GetInterchangeFee returns the interchange fee of the transaction based on the CostInterchangeValue field.
func (t *Transaction) GetInterchangeFeeRate() float64 {
	if t.CostInterchangeValue != nil && t.TransactionAmount != nil && *t.TransactionAmount != 0 {
		return *t.CostInterchangeValue / *t.TransactionAmount * 100
	}
	return 0
}

// GetRevenueMDRValueRate returns the revenue MDR value rate of the transaction based on the RevenueMDRValue field.
func (t *Transaction) GetRevenueMDRValueRate() float64 {
	if t.RevenueMDRValue != nil && t.TransactionAmount != nil && *t.TransactionAmount != 0 {
		return *t.RevenueMDRValue / *t.TransactionAmount * 100
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
