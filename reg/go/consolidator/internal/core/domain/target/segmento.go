package target_domain

import (
	"time"
)

// Segment represents the data structure for segments which will be used for fusing data between intercam, management and webservice
type Segmento struct {
	ID          int64     `gorm:"column:id"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
	SegmentName string    `gorm:"column:segment_name"`
	Description string    `gorm:"column:description"`
	SegmentCode int       `gorm:"column:segment_code"`
}

// NewSegmento creates a new instance of Segmento.
func NewSegmento() *Segmento {
	return &Segmento{}
}

// TableName specifies the table name for Segment struct
func (i *Segmento) TableName() string {
	return "cadoc_6334_v2.segmento"
}
