package target_domain

import (
	"time"
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

// TableName specifies the table name for Infresta struct
func (Infresta) TableName() string {
	return "cadoc_6334_v2.infresta"
}
