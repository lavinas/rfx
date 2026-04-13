package target_domain

import (
	"time"
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
func (Infrterm) TableName() string {
	return "cadoc_6334_v2.infrterm"
}
