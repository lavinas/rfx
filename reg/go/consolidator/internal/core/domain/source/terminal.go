package source_domain

// Terminal represents the terminal data for a specific year, quarter, and establishment.
type Terminal struct {
	TerminalCode      string `gorm:"column:terminal_code"`
	EstablishmentCode int64  `gorm:"column:establishment_code"`
	TerminalType      string `gorm:"column:terminal_type"`
}

// GetPOSQuantity returns the total quantity of POS terminals based on the TerminalType field.
func (t *Terminal) GetPOSQuantity() int {
	if t.TerminalType == "POS" {
		return 1
	}
	return 0
}

// GetPOSSharedQuantity returns the quantity of shared POS terminals based on the TerminalType field.
func (t *Terminal) GetPOSSharedQuantity() int {
	return 0
}

// GetPOSChipQuantity returns the quantity of chip POS terminals based on the TerminalType field.
func (t *Terminal) GetPOSChipQuantity() int {
	if t.TerminalType == "POS" {
		return 1
	}
	return 0
}

// GetPDVQuantity returns the quantity of PDV terminals based on the TerminalType field.
func (t *Terminal) GetPDVQuantity() int {
	if t.TerminalType == "PDV" {
		return 1
	}
	return 0
}