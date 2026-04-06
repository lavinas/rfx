package domain

import (
	"time"
)

// Management represents the data structure for management raw data
type Management struct {
	CdTransacaoFin           string     `gorm:"column:cd_transacao_fin"`
	Key1                     *string    `gorm:"column:key1"`
	DtProcessamento          *time.Time `gorm:"column:dt_processamento"`
	ValorTransacao           *float64   `gorm:"column:valor_transacao"`
	Bandeira                 *string    `gorm:"column:bandeira"`
	CdPessoaEstabelecimento  *int64     `gorm:"column:cd_pessoa_estabelecimento"`
	Mcc                      *string    `gorm:"column:mcc"`
	FormaCaptura             *string    `gorm:"column:forma_captura"`
	Funcao                   *string    `gorm:"column:funcao"`
	NumeroParcelas           *int       `gorm:"column:numero_parcelas"`
	DescontoValor            *float64   `gorm:"column:desconto_valor"`
	PercentualDesconto       *float64   `gorm:"column:percentual_desconto"`
	TransacId                *string    `gorm:"column:transac_id"`
	DtInserter               *time.Time `gorm:"column:dt_inserter"`
	TransactionalStatusID    *int64     `gorm:"column:transactional_status_id"`
	TransactionalStatusDate  *time.Time `gorm:"column:transactional_status_date"`
	ReconciliationStatusID   *int64     `gorm:"column:reconciliation_status_id"`
	ReconciliationStatusDate *time.Time `gorm:"column:reconciliation_status_date"`
}

// TableName specifies the table name for Management struct
func (Management) TableName() string {
	return "raw_data.management_transaction"
}


