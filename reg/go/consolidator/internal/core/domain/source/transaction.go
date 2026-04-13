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
