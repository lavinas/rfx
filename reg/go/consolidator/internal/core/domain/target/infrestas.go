package target_domain

import (
	"fmt"
	"maps"
	"slices"
	"time"

	source_domain "consolidator/internal/core/domain/source"
	"consolidator/internal/core/ports"
)

// Infresta represents the data structure for infresta which will be used for fusing data between intercam, management and webservice
type InfrestaItem struct {
	ID                                    int64     `gorm:"column:id"`
	CreatedAt                             time.Time `gorm:"column:created_at;type:timestamp"`
	UpdatedAt                             time.Time `gorm:"column:updated_at;type:timestamp"`
	Year                                  int       `gorm:"column:year"`
	Quarter                               int       `gorm:"column:quarter"`
	FederationUnit                        string    `gorm:"column:federation_unit"`
	EstablishmentTotalQuantity            int64     `gorm:"column:establishment_total_quantity"`
	EstablishmentManualCaptureQuantity    int64     `gorm:"column:establishment_manual_capture_quantity"`
	EstablishmentEletronicCaptureQuantity int64     `gorm:"column:establishment_eletronic_capture_quantity"`
	EstablishmentRemoteCaptureQuantity    int64     `gorm:"column:establishment_remote_capture_quantity"`
}

// Infresta represents the consolidated terminal data for a specific year and quarter.
type Infresta struct {
	DomainBase
	infresta      *InfrestaItem
	consolidation map[string]*InfrestaItem
}

// NewInfrestaConsolidation creates a new instance of Infresta.
func NewInfresta() *Infresta {
	return &Infresta{
		infresta:      &InfrestaItem{},
		consolidation: make(map[string]*InfrestaItem),
	}
}

// TableName specifies the table name for InfrestaItem struct
func (i *InfrestaItem) TableName() string {
	return "cadoc_6334_v2.infresta"
}

// GetFromClient returns the infresta data for a given transaction.
func (i *InfrestaItem) GetFromEstablishment(year int, quarter int, establishment *source_domain.Establishment) *InfrestaItem {
	return &InfrestaItem{
		Year:                                  year,
		Quarter:                               quarter,
		FederationUnit:                        establishment.GetFederationUnit(),
		EstablishmentTotalQuantity:            1,
		EstablishmentManualCaptureQuantity:    establishment.GetManualCaptureQuantity(),
		EstablishmentEletronicCaptureQuantity: establishment.GetEletronicCaptureQuantity(),
		EstablishmentRemoteCaptureQuantity:    establishment.GetRemoteCaptureQuantity(),
	}
}

// GetKey generates a unique key for the InfrestaItem struct based on its fields.
func (i *InfrestaItem) GetKey() string {
	return fmt.Sprintf("%d-%d-%s", i.Year, i.Quarter, i.FederationUnit)
}

// Delete removes the consolidated data for a specific year and quarter from the consolidation map.
func (i *Infresta) Delete(year int, quarter int, repository ports.Repository) error {
	// delete the consolidated data for the specified year and quarter from the repository
	if err := repository.Delete(&InfrestaItem{}, year, quarter); err != nil {
		return err
	}
	return nil
}

// Save persists the consolidated data for a specific year and quarter to the repository.
func (i *Infresta) Save(repository ports.Repository) error {
	if err := repository.Save(slices.Collect(maps.Values(i.consolidation))); err != nil {
		return err
	}
	return nil
}

// AddEstablishments processes a slice of establishments and updates the Infresta instance accordingly.
func (i *Infresta) AddEstablishments(year int, quarter int, establishments []*source_domain.Establishment) {
	// iterate over establishments and update the infresta data accordingly
	for _, e := range establishments {

		// only consider accredited establishments for the consolidation
		if !e.IsAccredited(year, quarter) {
			continue
		}

		// consolidate infresta data for the establishment
		infresta := i.infresta.GetFromEstablishment(year, quarter, e)
		key := infresta.GetKey()
		if existing, exists := i.consolidation[key]; exists {
			existing.EstablishmentTotalQuantity += infresta.EstablishmentTotalQuantity
			existing.EstablishmentManualCaptureQuantity += infresta.EstablishmentManualCaptureQuantity
			existing.EstablishmentEletronicCaptureQuantity += infresta.EstablishmentEletronicCaptureQuantity
			existing.EstablishmentRemoteCaptureQuantity += infresta.EstablishmentRemoteCaptureQuantity
			i.consolidation[key] = existing
		} else {
			i.consolidation[key] = infresta
		}
	}
}
