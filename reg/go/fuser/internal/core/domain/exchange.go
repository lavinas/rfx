package domain

import (
	"strconv"
	"time"
)

// Exchange represents the data structure for exchange raw data
type Exchange struct {
	CdTransacaoFin       string     `gorm:"column:cd_transacao_fin"`
	Key1                 *string    `gorm:"column:key1"`
	FormaCaptura         *string    `gorm:"column:forma_captura"`
	DtProcessamento      *time.Time `gorm:"column:dt_processamento;type:timestamp "`
	ValorTransacoes      *float64   `gorm:"column:valor_transacoes"`
	PercentualDesconto   *float64   `gorm:"column:percentual_desconto"`
	TaxaIntercambioValor *float64   `gorm:"column:taxa_intercambio_valor"`
	Bandeira             *string    `gorm:"column:bandeira"`
	Parcela              *string    `gorm:"column:parcela"`
	TipoCartao           *string    `gorm:"column:tipo_cartao"`
	Segmento             *string    `gorm:"column:segmento"`
	Bin                  *string    `gorm:"column:bin"`
	TransactionNsu       *string    `gorm:"column:transaction_nsu"`
	AuthorizationCode    *string    `gorm:"column:authorization_code"`
	ExtractorID          *int64     `gorm:"column:extractor_execution_id"`
	CardAcceptorID       *string    `gorm:"column:card_acceptor_id"`
}

// TableName specifies the table name for Exchange struct
func (Exchange) TableName() string {
	return "raw_data_v2.intercambio_transaction"
}

// Translate converts an Exchange instance to a Transaction instance
func (i Exchange) Translate() *Transaction {
	return &Transaction{
		ID:                          0, // ID will be set by the database upon insertion
		CreatedAt:                   time.Now(),
		UpdatedAt:                   time.Now(),
		Key1:                        i.GetKey1(),
		EstablishmentCode:           i.GetEstablishmentCode(),
		EstablishmentNature:         nil,
		EstablishmentMCC:            nil,
		EstablishmentTerminalCode:   nil,
		BIN:                         i.GetBIN(),
		AuthorizationCode:           i.GetAuthorizationCode(),
		TransactionNSU:              i.GetTransactionNSU(),
		TransactionDate:             i.GetTransactionDate(),
		TransactionSecondaryDate:    i.GetTransactionDate(),
		TransactionAmount:           i.GetTransactionAmount(),
		TransactionSecondaryAmount:  i.GetTransactionAmount(),
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

// GetKey1 returns the key1 value of the transaction, if available
func (i *Exchange) GetKey1() string {
	if i.Key1 == nil || *i.Key1 == "" {
		i.Key1 = new(string)
		*i.Key1 = "IC_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	}
	return *i.Key1
}

// GetEstablishmentCode returns the establishment code of the transaction, if available
func (i *Exchange) GetEstablishmentCode() *int64 {
	if i.CardAcceptorID != nil {
		if code, err := strconv.ParseInt(*i.CardAcceptorID, 10, 64); err == nil {
			return &code
		}
	}
	return nil
}

// GetBIN returns the BIN of the transaction, if available
func (i *Exchange) GetBIN() *int64 {
	if i.Bin != nil {
		if bin, err := strconv.ParseInt(*i.Bin, 10, 64); err == nil {
			return &bin
		}
	}
	return nil
}

// GetAuthorizationCode returns the authorization code of the transaction, if available
func (i *Exchange) GetAuthorizationCode() *string {
	if i.AuthorizationCode != nil {
		str := new(string)
		*str = *i.AuthorizationCode
		return str
	}
	return nil
}

// GetTransactionNSU returns the transaction NSU of the transaction, if available
func (i *Exchange) GetTransactionNSU() *string {
	if i.TransactionNsu != nil {
		str := new(string)
		*str = *i.TransactionNsu
		return str
	}
	return nil
}

// GetTransactionDate returns the transaction date of the transaction, if available
func (i *Exchange) GetTransactionDate() *time.Time {
	if i.DtProcessamento != nil {
		ret := new(time.Time)
		*ret = *i.DtProcessamento
		return ret
	}
	return nil
}

// GetTransactionAmount returns the transaction amount of the transaction, if available
func (i *Exchange) GetTransactionAmount() *float64 {
	if i.ValorTransacoes != nil {
		amount := new(float64)
		*amount = *i.ValorTransacoes
		return amount
	}
	return nil
}

// GetTransactionInstallments returns the transaction installments of the transaction, if available
func (i *Exchange) GetTransactionInstallments() *int64 {
	if i.Parcela != nil {
		if installments, err := strconv.ParseInt(*i.Parcela, 10, 64); err == nil {
			return &installments
		}
	}
	return nil
}

// Getbrand returns the brand of the transaction, if available
func (i *Exchange) GetBrand() *string {
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
func (i *Exchange) GetProduct() *string {
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
func (i *Exchange) GetCapture() *string {
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
func (i *Exchange) GetCostInterchangeValue() *float64 {
	if i.TaxaIntercambioValor != nil {
		cost := new(float64)
		*cost = *i.TaxaIntercambioValor
		return cost
	}
	return nil
}

// GetHighPriority returns the high source priority of the transaction, if available
func (i *Exchange) GetHighPriority() *int64 {
	priority := int64(20) // Default value for high source priority
	return &priority
}

// GetStatusID returns the status ID of the transaction, if available
func (i *Exchange) GetStatusID() *int64 {
	statusID := int64(0) // Default value for status ID
	return &statusID
}

// GetStatusName returns the status name of the transaction, if available
func (i *Exchange) GetStatusName() *string {
	statusName := "Exchange" // Default value for status name
	return &statusName
}

// GetPeriodClosingID returns the period closing ID of the transaction, if available
func (i *Exchange) GetPeriodClosingID() *int64 {
	periodClosingID := int64(1) // Default value for period closing ID
	return &periodClosingID
}

// GetTransacID returns the transac ID of the transaction, if available
func (i *Exchange) GetTransacID() *string {
	if i.CdTransacaoFin != "" {
		str := new(string)
		*str = i.CdTransacaoFin
		return str
	}
	return nil
}
