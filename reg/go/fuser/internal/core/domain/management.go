package domain

import (
	"time"
	"strconv"
)

// Management represents the data structure for management raw data
type Management struct {
	CdTransacaoFin           string     `gorm:"column:cd_transacao_fin"`
	Key1                     *string    `gorm:"column:key1"`
	DtProcessamento          *time.Time `gorm:"column:dt_processamento;type:timestamp"`
	ValorTransacao           *float64   `gorm:"column:valor_transacao"`
	Bandeira                 *string    `gorm:"column:bandeira"`
	CdPessoaEstabelecimento  *int64     `gorm:"column:cd_pessoa_estabelecimento"`
	Mcc                      *string    `gorm:"column:mcc"`
	FormaCaptura             *string    `gorm:"column:forma_captura"`
	Funcao                   *string    `gorm:"column:funcao"`
	NumeroParcelas           *int64     `gorm:"column:numero_parcelas"`
	DescontoValor            *float64   `gorm:"column:desconto_valor"`
	PercentualDesconto       *float64   `gorm:"column:percentual_desconto"`
	TransacId                *string    `gorm:"column:transac_id"`
	DtInserter               *time.Time `gorm:"column:dt_inserter;type:timestamp"`
	TransactionalStatusID    *int64     `gorm:"column:transactional_status_id"`
	TransactionalStatusDate  *time.Time `gorm:"column:transactional_status_date;type:timestamp"`
	ReconciliationStatusID   *int64     `gorm:"column:reconciliation_status_id"`
	ReconciliationStatusDate *time.Time `gorm:"column:reconciliation_status_date;type:timestamp"`
}

// TableName specifies the table name for Management struct
func (Management) TableName() string {
	return "raw_data_v2.management_transaction"
}

// Translate converts a Management instance to a Transaction instance
func (i Management) Translate() *Transaction {
	return &Transaction{
		ID:                          0, // ID will be set by the database upon insertion
		CreatedAt:                   time.Now(),
		UpdatedAt:                   time.Now(),
		Key1:                        i.GetKey1(),
		EstablishmentCode:           i.GetEstablishmentCode(),
		EstablishmentNature:         i.GetEstablishmentNature(),
		EstablishmentMCC:            i.GetEstablishmentMCC(),
		EstablishmentTerminalCode:   i.GetEstablishmentTerminalCode(),
		BIN:                         nil,
		AuthorizationCode:           nil,
		TransactionNSU:              nil,
		TransactionDate:             i.GetTransactionDate(),
		TransactionAmount:           i.GetTransactionAmount(),
		TransactionInstallments:     i.GetTransactionInstallments(),
		TransactionInstallmentsType: nil,
		TransactionBrand:            i.GetBrand(),
		TransactionProduct:          i.GetProduct(),
		TransactionCapture:          i.GetCapture(),
		RevenueMDRValue:             i.GetRevenueMDRValue(),
		CostInterchangeValue:        nil,
		HighSourcePriority:          i.GetHighPriority(),
		StatusID:                    i.GetStatusID(),
		StatusName:                  i.GetStatusName(),
		StatusCount:                 1,
		PeriodDate:                  i.GetTransactionDate(),
		PeriodClosingID:             i.GetPeriodClosingID(),
		TransacID:                   i.GetTransacID(),
	}
}

// GetKey1 returns the key1 value of the transaction, if available
func (i Management) GetKey1() string {
	if i.Key1 == nil || *i.Key1 == "" {
		i.Key1 = new(string)
		*i.Key1 = "MG_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	}
	return *i.Key1
}

// GetEstablishmentCode returns the establishment code of the transaction, if available
func (i Management) GetEstablishmentCode() *int64 {
	if i.CdPessoaEstabelecimento != nil {
		return i.CdPessoaEstabelecimento
	}
	return nil
}

// GetEstablishmentNature returns the establishment nature of the transaction, if available
func (i Management) GetEstablishmentNature() *int64 {
	var nature int64 = 1 // Placeholder for actual implementation
	return &nature
}

// GetStatusID returns the status ID of the transaction, if available
func (i Management) GetStatusID() *int64 {
	statusID := int64(1) // Default value for status ID
	return &statusID
}

// GetStatusName returns the status name of the transaction, if available
func (i Management) GetStatusName() *string {
	statusName := "Management" // Default value for status name
	return &statusName
}

// GetPeriodClosingID returns the period closing ID of the transaction, if available
func (i Management) GetPeriodClosingID() *int64 {
	periodClosingID := int64(1) // Default value for period closing ID
	return &periodClosingID
}

// GetTransacID returns the transac ID of the transaction, if available
func (i Management) GetTransacID() *string {
	if i.CdTransacaoFin != "" {
		str := new(string)
		*str = i.CdTransacaoFin
		return str
	}
	return nil
}

// GetEstablishmentMCC returns the establishment MCC of the transaction, if available
func (i Management) GetEstablishmentMCC() *int64 {
	if i.Mcc != nil {
		if mcc, err := strconv.ParseInt(*i.Mcc, 10, 64); err == nil {
			return &mcc
		}
	}
	return nil
}

// GetEEstablishmentTerminalCode returns the establishment terminal code of the transaction, if available
func (i Management) GetEstablishmentTerminalCode() *int64 {
	if i.FormaCaptura != nil {
		terminalCode := int64(0) // Default value for terminal code
		return &terminalCode
	}
	return nil
}

// GetTransactionDate returns the transaction date, if available
func (i Management) GetTransactionDate() *time.Time {
	if i.DtProcessamento != nil {
		return i.DtProcessamento
	}
	return nil
}

// GetTransactionAmount returns the transaction amount, if available
func (i Management) GetTransactionAmount() *float64 {
	if i.ValorTransacao != nil {
		return i.ValorTransacao
	}
	return nil
}

// GetTransactionInstallments returns the number of installments for the transaction, if available
func (i Management) GetTransactionInstallments() *int64 {
	if i.NumeroParcelas != nil {
		return i.NumeroParcelas
	}
	return nil
}

// GetBrand returns the brand of the transaction, if available
func (i Management) GetBrand() *string {
	brandMap := map[string]string{
		"31": "E",
		"32": "E",
		"26": "M",
		"29": "M",
		"27": "V",
		"30": "V",
	}
	if i.Bandeira != nil {
		if brand, ok := brandMap[*i.Bandeira]; ok {
			return &brand
		}
	}
	return nil
}

// GetProduct returns the product of the transaction, if available
func (i Management) GetProduct() *string {
	productMap := map[string]string{
		"2": "CR",
		"1": "DB",
	}
	if i.Funcao != nil {
		if product, ok := productMap[*i.Funcao]; ok {
			return &product
		}
	}
	return nil
}

// GetCapture returns the capture method of the transaction, if available
func (i Management) GetCapture() *string {
	mapping := map[string]string{
		"07": "CTC",
		"02": "CHP",
	}
	if i.FormaCaptura != nil {
		if capture, exists := mapping[*i.FormaCaptura]; exists {
			ret := new(string)
			*ret = capture
			return ret
		} else {
			ret := new(string)
			*ret = "TAR" // Default value for unknown capture methods
			return ret
		}
	}
	return nil
}

// GetRevenueMDRValue returns the revenue MDR value of the transaction, if available
func (i Management) GetRevenueMDRValue() *float64 {
	if i.DescontoValor != nil {
		return i.DescontoValor
	}
	return nil
}

// GetHighPriority returns the high source priority of the transaction, if available
func (i Management) GetHighPriority() *int64 {
	priority := int64(10) // Default value for high source priority
	return &priority
}

// MergeManagement merges two  transactions into one, prioritizing Intercam values over Management values when both are available
func MergeManagement(interTransaction *Transaction, repoTransaction *Transaction) {
	if repoTransaction.EstablishmentCode == nil {
		repoTransaction.EstablishmentCode = interTransaction.EstablishmentCode
	}
	if repoTransaction.EstablishmentNature == nil {
		repoTransaction.EstablishmentNature = interTransaction.EstablishmentNature
	}
	if repoTransaction.EstablishmentMCC == nil {
		repoTransaction.EstablishmentMCC = interTransaction.EstablishmentMCC
	}
	if repoTransaction.EstablishmentTerminalCode == nil {
		repoTransaction.EstablishmentTerminalCode = interTransaction.EstablishmentTerminalCode
	}
	if repoTransaction.BIN == nil {
		repoTransaction.BIN = interTransaction.BIN
	}
	if repoTransaction.AuthorizationCode == nil {
		repoTransaction.AuthorizationCode = interTransaction.AuthorizationCode
	}
	if repoTransaction.TransactionNSU == nil {
		repoTransaction.TransactionNSU = interTransaction.TransactionNSU
	}
	if repoTransaction.TransactionDate == nil {
		repoTransaction.TransactionDate = interTransaction.TransactionDate
	}
	if repoTransaction.TransactionAmount == nil {
		repoTransaction.TransactionAmount = interTransaction.TransactionAmount
	}
	if repoTransaction.TransactionInstallments == nil {
		repoTransaction.TransactionInstallments = interTransaction.TransactionInstallments
	}
	if repoTransaction.TransactionBrand == nil {
		repoTransaction.TransactionBrand = interTransaction.TransactionBrand
	}
	if repoTransaction.TransactionProduct == nil {
		repoTransaction.TransactionProduct = interTransaction.TransactionProduct
	}
	if repoTransaction.TransactionCapture == nil {
		repoTransaction.TransactionCapture = interTransaction.TransactionCapture
	}
	if repoTransaction.CostInterchangeValue == nil {
		repoTransaction.CostInterchangeValue = interTransaction.CostInterchangeValue
	}
	if repoTransaction.HighSourcePriority == nil {
		repoTransaction.HighSourcePriority = interTransaction.HighSourcePriority
	}
	if repoTransaction.PeriodDate == nil {
		repoTransaction.PeriodDate = interTransaction.PeriodDate
	}
	if repoTransaction.PeriodClosingID == nil {
		repoTransaction.PeriodClosingID = interTransaction.PeriodClosingID
	}
	if repoTransaction.TransacID == nil {
		repoTransaction.TransacID = interTransaction.TransacID
	}
	if repoTransaction.RevenueMDRValue == nil {
		repoTransaction.RevenueMDRValue = interTransaction.RevenueMDRValue
	}
	// Calculate status
	if *repoTransaction.StatusID == 0 {
		repoTransaction.StatusCount = 0
		*repoTransaction.StatusID = 2
		*repoTransaction.StatusName = "Pronto" 
	}
}