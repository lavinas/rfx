package domain

import (
	"time"
)

// Webservice represents the data structure for webservice raw data
type Webservice struct {
	RefNumBnd                 *string    `gorm:"column:ref_num_bnd"`
	Key1                      *string    `gorm:"column:key1"`
	RefNumFis                 string     `gorm:"column:ref_num_fis"`
	TransactionBrand          *string    `gorm:"column:transaction_brand"`
	TransactionProduct        *string    `gorm:"column:transaction_product"`
	TransactionDate           *time.Time `gorm:"column:transaction_date"`
	DtPos                     *time.Time `gorm:"column:dt_pos"`
	EstablishmentTerminalCode *string    `gorm:"column:establishment_terminal_code"`
	TermID                    *string    `gorm:"column:term_id"`
	TransactionAmount         *float64   `gorm:"column:transaction_amount"`
	QtdParc                   *int       `gorm:"column:qtd_parc"`
	Bin                       *string    `gorm:"column:bin"`
	DtInserter                *time.Time `gorm:"column:dt_inserter"`
	TransactionalStatusID     int64      `gorm:"column:transactional_status_id"`
	TransactionalStatusDate   *time.Time `gorm:"column:transactional_status_date"`
	ReconciliationStatusID    int64      `gorm:"column:reconciliation_status_id"`
	ReconciliationStatusDate  *time.Time `gorm:"column:reconciliation_status_date"`
}

// TableName specifies the table name for Webservice struct
func (Webservice) TableName() string {
	return "raw_data.webservice_transaction"
}
