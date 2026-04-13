package target_domain

import (
	"time"
)

// Segment represents the data structure for segments which will be used for fusing data between intercam, management and webservice
type Segment struct {
	ID          int64     `gorm:"column:id"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
	SegmentCode int16     `gorm:"column:segment_code"`
	SegmentName string    `gorm:"column:segment_name"`
	Description string    `gorm:"column:description"`
}

// TableName specifies the table name for Segment struct
func (Segment) TableName() string {
	return "cadoc_6334_v2.segments"
}
