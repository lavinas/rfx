package target_domain


import (
	"fmt"
	"time"

	source_domain "consolidator/internal/core/domain/source"
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

// NewInfrestaItem creates a new instance of InfrestaItem.
func NewInfrestaItem() *InfrestaItem {
	return &InfrestaItem{}
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
