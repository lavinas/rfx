package domain

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
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

// SetForInsert sets the Key2 field of the transaction based on available data
func (t *Transaction) PrepareForInsert() {
	// Generate Key2 based on available data
	str := strconv.FormatFloat(*t.TransactionAmount, 'f', 2, 64) + "_" + strconv.FormatInt(*t.TransactionInstallments, 10) + "_" + *t.TransactionBrand + "_" + *t.TransactionProduct + "_" + *t.TransactionCapture
	md5Hash := md5.Sum([]byte(str))
	hashString := hex.EncodeToString(md5Hash[:])
	if t.Key2 == nil || *t.Key2 == "" {
		t.Key2 = new(string)
	}
	*t.Key2 = hashString
}
