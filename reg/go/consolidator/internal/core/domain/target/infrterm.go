package target_domain

import (
	"time"
	"fmt"

	source_domain "consolidator/internal/core/domain/source"
)

// Infrterm represents the infrastructure term data for a specific year, quarter, and federation unit.
type Infrterm struct {
	ID                int64     `gorm:"primaryKey"`
	CreatedAt         time.Time `gorm:"autoCreateTime"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime"`
	Year              int       `gorm:"column:year"`
	Quarter           int       `gorm:"column:quarter"`
	FederationUnit    string    `gorm:"column:federation_unit"`
	PosTotalQuantity  int64     `gorm:"column:pos_total_quantity"`
	PosSharedQuantity int64     `gorm:"column:pos_shared_quantity"`
	PosChipQuantity   int64     `gorm:"column:pos_chip_quantity"`
	PdvQuantity       int64     `gorm:"column:pdv_quantity"`
}

// TableName specifies the table name for Infrterm struct
func (i *Infrterm) TableName() string {
	return "cadoc_6334_v2.infrterm"
}

// GetFromTerminal returns an Infrterm instance populated with data from a given terminal.
func (i *Infrterm) GetFromTerminal(year int, quarter int, term *source_domain.Terminal) *Infrterm {
	return &Infrterm{
		Year:              year,
		Quarter:           quarter,
		FederationUnit:    "SP", // DO IT
		PosTotalQuantity:  int64(term.GetPOSQuantity()),
		PosSharedQuantity: int64(term.GetPOSSharedQuantity()),
		PosChipQuantity:   int64(term.GetPOSChipQuantity()),
		PdvQuantity:       int64(term.GetPDVQuantity()),
	}
}

// GetKey generates a unique key for the Infrterm struct based on its fields.
func (i *Infrterm) GetKey() string {
	return fmt.Sprintf("%d-%d-%s", i.Year, i.Quarter, i.FederationUnit)
}

// AddTerminals processes a slice of terminals and updates the Infrterm instance accordingly.
func (i *Infrterm) AddTerminals(terminals []*source_domain.Terminal, items map[string]*Infrterm) {
	for _, t := range terminals {
		infrterm := i.GetFromTerminal(i.Year, i.Quarter, t)
		key := infrterm.GetKey()
		if existing, exists := items[key]; exists {
			existing.PosTotalQuantity += infrterm.PosTotalQuantity
			existing.PosSharedQuantity += infrterm.PosSharedQuantity
			existing.PosChipQuantity += infrterm.PosChipQuantity
			existing.PdvQuantity += infrterm.PdvQuantity
			items[key] = existing
		} else {
			items[key] = infrterm
		}
	}
}
