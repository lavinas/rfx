package domain

import (
	"time"
)

// Intercam represents the data structure for intercam raw data
type Intercam struct {
	CdTransacaoFin           string     `gorm:"column:cd_transacao_fin"`
	Key1                     *string    `gorm:"column:key1"`
	FormaCaptura             *string    `gorm:"column:forma_captura"`
	DtProcessamento          *time.Time `gorm:"column:dt_processamento"`
	ValorTransacoes          *float64   `gorm:"column:valor_transacoes"`
	PercentualDesconto       *float64   `gorm:"column:percentual_desconto"`
	TaxaIntercambioValor     *float64   `gorm:"column:taxa_intercambio_valor"`
	Bandeira                 *string    `gorm:"column:bandeira"`
	Parcela                  *string    `gorm:"column:parcela"`
	TipoCartao               *string    `gorm:"column:tipo_cartao"`
	Segmento                 *string    `gorm:"column:segmento"`
	Bin                      *string    `gorm:"column:bin"`
	TransactionNsu           *string    `gorm:"column:transaction_nsu"`
	AuthorizationCode        *string    `gorm:"column:authorization_code"`
	DtInserter               *time.Time `gorm:"column:dt_inserter"`
	TransactionalStatusID    *int64     `gorm:"column:transactional_status_id"`
	TransactionalStatusDate  *time.Time `gorm:"column:transactional_status_date"`
	ReconciliationStatusID   *int64     `gorm:"column:reconciliation_status_id"`
	ReconciliationStatusDate *time.Time `gorm:"column:reconciliation_status_date"`
}

// TableName specifies the table name for Intercam struct
func (Intercam) TableName() string {
	return "raw_data.intercambio_transaction"
}
