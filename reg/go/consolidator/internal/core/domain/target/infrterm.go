package target_domain

import (
	"fmt"
	"maps"
	"slices"
	"time"

	source_domain "consolidator/internal/core/domain/source"
	"consolidator/internal/core/ports"
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

// Infrterm represents the consolidated terminal data for a specific year and quarter.
type Infrterm struct {
	DomainBase
	infrterm      *InfrtermItem
	consolidation map[string]*InfrtermItem
}

// NewInfrterm creates a new instance of Infrterm.
func NewInfrterm() *Infrterm {
	return &Infrterm{
		infrterm:      &InfrtermItem{},
		consolidation: make(map[string]*InfrtermItem),
	}
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

// Delete removes the consolidated data for a specific year and quarter from the consolidation map.
func (i *Infrterm) Delete(year int, quarter int, repository ports.Repository) error {
	// delete the consolidated data for the specified year and quarter from the repository
	if err := repository.Delete(&InfrtermItem{}, year, quarter); err != nil {
		return err
	}
	return nil
}

// Save persists the consolidated data for a specific year and quarter to the repository.
func (i *Infrterm) Save(repository ports.Repository) error {
	if err := repository.Save(slices.Collect(maps.Values(i.consolidation))); err != nil {
		return err
	}
	return nil
}

// AddTerminals processes a slice of terminals and updates the Infrterm instance accordingly.
func (i *Infrterm) AddTerminals(year int, quarter int, terminals []*source_domain.Terminal, esblishmentMap map[int64]string) {
	for _, t := range terminals {
		infrterm := i.infrterm.GetFromTerminal(year, quarter, t, esblishmentMap)
		if infrterm == nil {
			continue
		}
		key := infrterm.GetKey()
		if existing, exists := i.consolidation[key]; exists {
			existing.PosTotalQuantity += infrterm.PosSharedQuantity + infrterm.PosChipQuantity + infrterm.PdvQuantity
			existing.PosSharedQuantity += infrterm.PosSharedQuantity
			existing.PosChipQuantity += infrterm.PosChipQuantity
			existing.PdvQuantity += infrterm.PdvQuantity
			i.consolidation[key] = existing
		} else {
			i.consolidation[key] = infrterm
		}
	}
}
