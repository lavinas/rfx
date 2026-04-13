package target_domain

import (
	"time"
)

// Contact represents a contact associated with a transaction, such as a buyer or seller.
type Contact struct {
	ID          int       `gorm:"primaryKey"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
	Year        int       `gorm:"column:year"`
	Quarter     int       `gorm:"column:quarter"`
	ContactType string    `gorm:"column:contact_type"`
	Name        string    `gorm:"column:name"`
	Position    string    `gorm:"column:position"`
	PhoneNumber string    `gorm:"column:phone_number"`
	Email       string    `gorm:"column:email"`
}

// TableName specifies the table name for Contact struct
func (Contact) TableName() string {
	return "cadoc_6334_v2.contact"
}
