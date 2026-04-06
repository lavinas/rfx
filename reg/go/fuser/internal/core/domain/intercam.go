package domain

import (
	"strconv"
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

// Translate converts an Intercam instance to a Transaction instance
func (i Intercam) Translate() *Transaction {
	return &Transaction{
		ID:                          0, // ID will be set by the database upon insertion
		CreatedAt:                   time.Now(),
		UpdatedAt:                   time.Now(),
		Key1:                        *i.Key1,
		EstablishmentCode:           nil, // This field is not present in Intercam, set to nil or default value
		EstablishmentNature:         nil, // This field is not present in Intercam, set to nil or default value
		EstablishmentMCC:            nil, // This field is not present in Intercam, set to nil or default value
		EstablishmentTerminalCode:   nil, // This field is not present in Intercam, set to nil or default value
		BIN:                         i.GetBIN(),
		AuthorizationCode:           i.GetAuthorizationCode(),
		TransactionNSU:              i.GetTransactionNSU(),
		TransactionDate:             i.GetTransactionDate(),
		TransactionAmount:           i.GetTransactionAmount(),
		TransactionInstallments:     i.GetTransactionInstallments(),
		TransactionInstallmentsType: nil,
		TransactionBrand:            i.GetBrand(),
		TransactionProduct:          i.GetProduct(),
		TransactionCapture:          i.GetCapture(),
		RevenueMDRValue:             nil,
		CostInterchangeValue:        i.GetCostInterchangeValue(),
		HighSourcePriority:          i.GetHighPriority(),
		StatusID:                    i.GetStatusID(),
		StatusName:                  i.GetStatusName(),
		StatusCount:                 1,
		PeriodDate:                  i.GetTransactionDate(),
		PeriodClosingID:             i.GetPeriodClosingID(),
		TransacID:                   i.GetTransacID(),
	}
}

// GetBIN returns the BIN of the transaction, if available
func (i Intercam) GetBIN() *int64 {
	if i.Bin != nil {
		if bin, err := strconv.ParseInt(*i.Bin, 10, 64); err == nil {
			return &bin
		}
	}
	return nil
}

// GetAuthorizationCode returns the authorization code of the transaction, if available
func (i Intercam) GetAuthorizationCode() *string {
	if i.AuthorizationCode != nil {
		str := new(string)
		*str = *i.AuthorizationCode
		return str
	}
	return nil
}

// GetTransactionNSU returns the transaction NSU of the transaction, if available
func (i Intercam) GetTransactionNSU() *string {
	if i.TransactionNsu != nil {
		str := new(string)
		*str = *i.TransactionNsu
		return str
	}
	return nil
}

// GetTransactionDate returns the transaction date of the transaction, if available
func (i Intercam) GetTransactionDate() *time.Time {
	if i.DtProcessamento != nil {
		ret := new(time.Time)
		*ret = *i.DtProcessamento
		return ret
	}
	return nil
}

// GetTransactionAmount returns the transaction amount of the transaction, if available
func (i Intercam) GetTransactionAmount() *float64 {
	if i.ValorTransacoes != nil {
		amount := new(float64)
		*amount = *i.ValorTransacoes
		return amount
	}
	return nil
}

// GetTransactionInstallments returns the transaction installments of the transaction, if available
func (i Intercam) GetTransactionInstallments() *int64 {
	if i.Parcela != nil {
		if installments, err := strconv.ParseInt(*i.Parcela, 10, 64); err == nil {
			return &installments
		}
	}
	return nil
}

// Getbrand returns the brand of the transaction, if available
func (i Intercam) GetBrand() *string {
	mapping := map[string]string{
		"VISA":       "V",
		"MASTERCARD": "M",
		"ELO":        "E",
	}
	if i.Bandeira != nil {
		if brand, exists := mapping[*i.Bandeira]; exists {
			ret := new(string)
			*ret = brand
			return ret
		}
	}
	return nil
}

// GetProduct returns the product of the transaction, if available
func (i Intercam) GetProduct() *string {
	mapping := map[string]string{
		"CREDIT": "CR",
		"DEBIT":  "DB",
	}
	if i.TipoCartao != nil {
		if product, exists := mapping[*i.TipoCartao]; exists {
			ret := new(string)
			*ret = product
			return ret
		}
	}
	return nil
}

// GetCapture returns the capture method of the transaction, if available
func (i Intercam) GetCapture() *string {
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

// GetCostInterchangeValue returns the cost interchange value of the transaction, if available
func (i Intercam) GetCostInterchangeValue() *float64 {
	if i.TaxaIntercambioValor != nil {
		cost := new(float64)
		*cost = *i.TaxaIntercambioValor
		return cost
	}
	return nil
}

// GetHighPriority returns the high source priority of the transaction, if available
func (i Intercam) GetHighPriority() *int64 {
	priority := int64(20) // Default value for high source priority
	return &priority
}

// GetStatusID returns the status ID of the transaction, if available
func (i Intercam) GetStatusID() *int64 {
	statusID := int64(1) // Default value for status ID
	return &statusID
}

// GetStatusName returns the status name of the transaction, if available
func (i Intercam) GetStatusName() *string {
	statusName := "Pendente" // Default value for status name
	return &statusName
}

// GetPeriodClosingID returns the period closing ID of the transaction, if available
func (i Intercam) GetPeriodClosingID() *int64 {
	periodClosingID := int64(1) // Default value for period closing ID
	return &periodClosingID
}

// GetTransacID returns the transac ID of the transaction, if available
func (i Intercam) GetTransacID() *string {
	if i.CdTransacaoFin != "" {
		str := new(string)
		*str = i.CdTransacaoFin
		return str
	}
	return nil
}

// MergeTransactions merges two Intercam transactions into one, prioritizing non-nil values from the second transaction
func MergeIntercam(interTransaction *Transaction, repoTransaction *Transaction) {
	repoTransaction.BIN = interTransaction.BIN
	repoTransaction.AuthorizationCode = interTransaction.AuthorizationCode
	repoTransaction.TransactionNSU = interTransaction.TransactionNSU
	repoTransaction.TransactionDate = interTransaction.TransactionDate
	repoTransaction.TransactionAmount = interTransaction.TransactionAmount
	repoTransaction.TransactionInstallments = interTransaction.TransactionInstallments
	repoTransaction.TransactionBrand = interTransaction.TransactionBrand
	repoTransaction.TransactionProduct = interTransaction.TransactionProduct
	repoTransaction.TransactionCapture = interTransaction.TransactionCapture
	repoTransaction.CostInterchangeValue = interTransaction.CostInterchangeValue
	repoTransaction.HighSourcePriority = interTransaction.HighSourcePriority
	repoTransaction.PeriodDate = interTransaction.PeriodDate
	repoTransaction.PeriodClosingID = interTransaction.PeriodClosingID
	repoTransaction.TransacID = interTransaction.TransacID
	// Calculate status
	if repoTransaction.StatusCount == 2 {
		repoTransaction.StatusCount = 0
		*repoTransaction.StatusID = 2
		*repoTransaction.StatusName = "Pronto"
	}

}
