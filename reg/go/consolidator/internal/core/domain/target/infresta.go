package target_domain

import (
	"fmt"
	"time"

	source_domain "consolidator/internal/core/domain/source"
)

// Infresta represents the data structure for infresta which will be used for fusing data between intercam, management and webservice
type Infresta struct {
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

// NewInfresta creates a new instance of Infresta.
func NewInfresta() *Infresta {
	return &Infresta{}
}

// TableName specifies the table name for Infresta struct
func (i *Infresta) TableName() string {
	return "cadoc_6334_v2.infresta"
}

// GetFromClient returns the infresta data for a given transaction.
func (i *Infresta) GetFromEstablishment(year int, quarter int, establishment *source_domain.Establishment) *Infresta {
	return &Infresta{
		Year:                                  year,
		Quarter:                               quarter,
		FederationUnit:                        establishment.GetFederationUnit(),
		EstablishmentTotalQuantity:            1,
		EstablishmentManualCaptureQuantity:    establishment.GetManualCaptureQuantity(),
		EstablishmentEletronicCaptureQuantity: establishment.GetEletronicCaptureQuantity(),
		EstablishmentRemoteCaptureQuantity:    establishment.GetRemoteCaptureQuantity(),
	}
}

// GetKey generates a unique key for the Infresta struct based on its fields.
func (i *Infresta) GetKey() string {
	return fmt.Sprintf("%d-%d-%s", i.Year, i.Quarter, i.FederationUnit)
}

// AddEstablishments processes a slice of establishments and updates the Infresta instance accordingly.
func (i *Infresta) AddEstablishments(year int, quarter int, establishments []*source_domain.Establishment, items map[string]*Infresta) {
	// iterate over establishments and update the infresta data accordingly
	for _, e := range establishments {

		// only consider accredited establishments for the consolidation
		if !e.IsAccredited(year, quarter) {
			continue
		}

		// consolidate infresta data for the establishment
		infresta := i.GetFromEstablishment(year, quarter, e)
		key := infresta.GetKey()
		if existing, exists := items[key]; exists {
			existing.EstablishmentTotalQuantity += infresta.EstablishmentTotalQuantity
			existing.EstablishmentManualCaptureQuantity += infresta.EstablishmentManualCaptureQuantity
			existing.EstablishmentEletronicCaptureQuantity += infresta.EstablishmentEletronicCaptureQuantity
			existing.EstablishmentRemoteCaptureQuantity += infresta.EstablishmentRemoteCaptureQuantity
			items[key] = existing
		} else {
			items[key] = infresta
		}
	}
}
