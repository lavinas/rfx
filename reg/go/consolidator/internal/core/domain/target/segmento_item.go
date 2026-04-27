package target_domain

import (
	"time"

)

// SegmentoItem represents the data structure for segments which will be used for fusing data between intercam, management and webservice
type SegmentoItem struct {
	ID          int64     `gorm:"column:id"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
	Year        int       `gorm:"column:year"`
	Quarter     int       `gorm:"column:quarter"`
	SegmentName string    `gorm:"column:segment_name"`
	Description string    `gorm:"column:segment_description"`
	SegmentCode int       `gorm:"column:segment_code"`
}

// NewSegmentoItem creates a new instance of SegmentoItem.
func NewSegmentoItem() *SegmentoItem {
	return &SegmentoItem{}
}

// TableName specifies the table name for SegmentoItem struct
func (i *SegmentoItem) TableName() string {
	return "cadoc_6334_v2.segmento"
}
