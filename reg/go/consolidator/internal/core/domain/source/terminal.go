package source_domain

// Terminal represents the terminal data for a specific year, quarter, and establishment.
type Terminal struct {
	TerminalCode      string `gorm:"column:terminal_code"`
	EstablishmentCode int64  `gorm:"column:establishment_code"`
	TerminalType      string `gorm:"column:terminal_type"`
}
