package source_domain

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
	return "raw_data_v2.establishments"
}

// GetEstablishmentCode returns the establishment code of the establishment.
func (e *Establishment) GetCode() int64 {
	return e.EstablishmentCode
}

// GetFederationUnit returns the federation unit of the establishment.
func (e *Establishment) GetFederationUnit() string {
	if e.FederationUnit != nil {
		return *e.FederationUnit
	}
	return ""
}

// GetManualCaptureQuantity returns 1 if the establishment has manual capture, otherwise returns 0.
func (e *Establishment) GetManualCaptureQuantity() int64 {
	if e.HasManualCapture {
		return 1
	}
	return 0
}

// GetEletronicCaptureQuantity returns 1 if the establishment has eletronic capture, otherwise returns 0.
func (e *Establishment) GetEletronicCaptureQuantity() int64 {
	if e.HasEletronicCapture {
		return 1
	}
	return 0
}

// GetRemoteCaptureQuantity returns 1 if the establishment has remote capture, otherwise returns 0.
func (e *Establishment) GetRemoteCaptureQuantity() int64 {
	if e.HasRemoteCapture {
		return 1
	}
	return 0
}

// GetFunctionCode returns the function code of the establishment based on its capture capabilities.
func (e *Establishment) GetFunctionCodes() []string {
	var functions []string
	if e.HasCredit {
		functions = append(functions, "D")
	}
	if e.HasDebit {
		functions = append(functions, "C")
	}
	return functions
}

// GetBrandCodes returns the brand code of the establishment based on its card acceptance capabilities.
func (e *Establishment) GetBrandCodes() []int {
	var brands []int
	if e.HasVisa {
		brands = append(brands, 1)
	}
	if e.HasMastercard {
		brands = append(brands, 2)
	}
	if e.HasElo {
		brands = append(brands, 8)
	}
	return brands
}

// IsAccredited returns true if the establishment is accredited, otherwise returns false.
func (e *Establishment) IsAccredited(year int, quarter int) bool {
	lastDayOfQuarter := time.Date(year, time.Month(quarter*3), 1, 0, 0, 0, 0, time.UTC).AddDate(0, 1, 0)
	if e.AccreditationDate != nil && e.AccreditationDate.Before(lastDayOfQuarter) {
		return true
	}
	return false
}

// IsActive returns true if the establishment is active, otherwise returns false.
func (e *Establishment) IsActive(year int, quarter int) bool {
	if !e.IsAccredited(year, quarter) {
		return false
	}
	lastDayOfQuarter := time.Date(year, time.Month(quarter*3), 1, 0, 0, 0, 0, time.UTC).AddDate(0, 1, -1)
	backLimitDate := lastDayOfQuarter.AddDate(0, 0, -180)
	if e.AccreditationDate != nil && e.AccreditationDate.After(backLimitDate) {
		return true
	}
	return false
}

// GetBrands returns the brand codes of the establishment based on its card acceptance capabilities.
func (e *Establishment) GetBrands() []int {
	var brands []int
	if e.HasVisa {
		brands = append(brands, BrandMap["V"])
	}
	if e.HasMastercard {
		brands = append(brands, BrandMap["M"])
	}
	if e.HasElo {
		brands = append(brands, BrandMap["E"])
	}
	return brands
}

// GetFunctions returns the function codes of the establishment based on its capture capabilities.
func (e *Establishment) GetFunctions() []string {
	var functions []string
	if e.HasCredit {
		functions = append(functions, ProductMap["CR"])
	}
	if e.HasDebit {
		functions = append(functions, ProductMap["DB"])
	}
	return functions
}
