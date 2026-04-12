package domain

import (
	"time"
)

// Establishment represents the data structure for establishments which will be used for fusing data between intercam, management and webservice
type Establishment struct {
	EstablishmentCode        int64      `gorm:"column:establishment_code"`
	AccreditationDate        *time.Time `gorm:"column:accreditation_date;type:timestamp"`
	CompanyName              *string    `gorm:"column:company_name"`
	TradingName              *string    `gorm:"column:trading_name"`
	CNPJ                     *string    `gorm:"column:cnpj"`
	CPF                      *string    `gorm:"column:cpf"`
	MCCCode                  *string    `gorm:"column:mcc_code"`
	Address                  *string    `gorm:"column:address"`
	CEP                      *string    `gorm:"column:cep"`
	CityIBGECode             *int64     `gorm:"column:city_ibge_code"`
	FederationUnit           *string    `gorm:"column:federation_unit"`
	ContactName              *string    `gorm:"column:contact_name"`
	ContactPhone             *int64     `gorm:"column:contact_phone"`
	ContactEmail             *string    `gorm:"column:contact_email"`
	HasDebit                 bool       `gorm:"column:has_debit"`
	HasCredit                bool       `gorm:"column:has_credit"`
	HasVisa                  bool       `gorm:"column:has_visa"`
	HasMastercard            bool       `gorm:"column:has_mastercard"`
	HasElo                   bool       `gorm:"column:has_elo"`
	HasManualCapture         bool       `gorm:"column:has_manual_capture"`
	HasEletronicCapture      bool       `gorm:"column:has_eletronic_capture"`
	HasRemoteCapture         bool       `gorm:"column:has_remote_capture"`
	IsSubacquirer            bool       `gorm:"column:is_subacquirer"`
	DtInserter               time.Time  `gorm:"column:dt_inserter;type:timestamp with time zone;default:now()"`
	TransactionalStatusID    int64      `gorm:"column:transactional_status_id;default:0"`
	TransactionalStatusDate  *time.Time `gorm:"column:transactional_status_date;type:timestamp"`
	ReconciliationStatusID   int64      `gorm:"column:reconciliation_status_id;default:0"`
	ReconciliationStatusDate *time.Time `gorm:"column:reconciliation_status_date;type:timestamp"`
}

// TableName specifies the table name for Establishment struct
func (Establishment) TableName() string {
	return "raw_date_v2.establishments"
}
