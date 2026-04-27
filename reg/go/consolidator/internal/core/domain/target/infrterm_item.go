package target_domain

import (
	"fmt"
	"time"

	source_domain "consolidator/internal/core/domain/source"
)

// InfrtermItem represents the infrastructure term data for a specific year, quarter, and federation unit.
type InfrtermItem struct {
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

// NewInfrtermItem creates a new instance of InfrtermItem.
func NewInfrtermItem() *InfrtermItem {
	return &InfrtermItem{}
}

// TableName specifies the table name for InfrtermItem struct
func (i *InfrtermItem) TableName() string {
	return "cadoc_6334_v2.infrterm"
}

// GetFromTerminal returns an InfrtermItem instance populated with data from a given terminal.
func (i *InfrtermItem) GetFromTerminal(year int, quarter int, term *source_domain.Terminal, clientsFU map[int64]string) *InfrtermItem {
	fu, exists := clientsFU[term.GetEstablishmentCode()]
	if !exists {
		return nil
	}
	return &InfrtermItem{
		Year:              year,
		Quarter:           quarter,
		FederationUnit:    fu,
		PosTotalQuantity:  int64(term.GetPOSQuantity()),
		PosSharedQuantity: int64(term.GetPOSSharedQuantity()),
		PosChipQuantity:   int64(term.GetPOSChipQuantity()),
		PdvQuantity:       int64(term.GetPDVQuantity()),
	}
}

// GetKey generates a unique key for the InfrtermItem struct based on its fields.
func (i *InfrtermItem) GetKey() string {
	return fmt.Sprintf("%d-%d-%s", i.Year, i.Quarter, i.FederationUnit)
}
