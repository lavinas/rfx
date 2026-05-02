package domain

import (
	"bufio"
	"fmt"

	"validator/internal/port"

	"github.com/ianlopshire/go-fixedwidth"
)

// Contact represents the Contact data model.
type Contact struct {
	Year        int64  `fixed:"1,4" gorm:"column:year"`
	Quarter     int64  `fixed:"5,5" gorm:"column:quarter"`
	ContactType string `fixed:"6,6" gorm:"column:contact_type"`
	Name        string `fixed:"7,56" gorm:"column:name"`
	Position    string `fixed:"57,106" gorm:"column:position"`
	Phone       string `fixed:"107,156" gorm:"column:phone_number"`
	Email       string `fixed:"157,206" gorm:"column:email"`
}

// TableName specifies the table name for Contact struct
func (c *Contact) TableName() string {
	return "contatos"
}

// NewContact creates a new Contact instance.
func NewContact() *Contact {
	return &Contact{}
}

// GetName gets name of the report
func (c *Contact) GetName() string {
	return "CONTATOS"
}

// Format marshals the Contact struct into a fixed-width format.
func (c *Contact) Format() string {
	ret := ""
	ret += fmt.Sprintf("%04d", c.Year)
	ret += fmt.Sprintf("%01d", c.Quarter)
	ret += fmt.Sprintf("%01s", c.ContactType)
	ret += fmt.Sprintf("%-50s", c.Name)
	ret += fmt.Sprintf("%-50s", c.Position)
	ret += fmt.Sprintf("%-50s", c.Phone)
	ret += fmt.Sprintf("%-50s", c.Email)
	return ret
}

// Validate validates the Contact header information.
func (c *Contact) Validate() error {
	if c.Year <= 0 {
		return fmt.Errorf("invalid year in header")
	}
	if c.Quarter <= 0 {
		return fmt.Errorf("invalid quarter in header")
	}
	if c.ContactType == "" {
		return fmt.Errorf("invalid contact type in header")
	}
	if c.Email == "" {
		return fmt.Errorf("invalid email in header")
	}
	return nil
}

// GetKey returns the unique key for the Contact record.
func (c *Contact) GetKey() string {
	return fmt.Sprintf("%d|%d|%s|%s", c.Year, c.Quarter, c.ContactType, c.Name)
}

// GetDB returns the database connection.
func (c *Contact) GetDB(repo port.Repository, year int, quarter int) (map[string]port.Report, error) {
	var records []*Contact
	err := repo.FindAll(&records, 0, 0, "", "year = ? AND quarter = ?", year, quarter)
	if err != nil {
		return nil, err
	}
	ret := make(map[string]port.Report)
	for _, r := range records {
		ret[r.GetKey()] = r
	}
	return ret, nil
}

// Parse parses the Contact data from a fixed-width file.
func (c *Contact) Parse(line string) (*Contact, error) {
	err := fixedwidth.Unmarshal([]byte(line), c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// String returns the string representation of the Contact.
func (c *Contact) String() string {
	return fmt.Sprintf("Contact{Year: %d, Quarter: %d, ContactType: %s, Name: %s, Position: %s, Phone: %s, Email: %s}",
		c.Year, c.Quarter, c.ContactType, c.Name, c.Position, c.Phone, c.Email)
}

// ParseContactFile parses a file containing Contact records.
func (c *Contact) ParseContactFile(file *bufio.Scanner) ([]*Contact, error) {
	// read header
	if !file.Scan() {
		return nil, fmt.Errorf("file is empty")
	}
	headerLine := file.Text()
	header := &RankingHeader{}
	if _, err := header.Parse(headerLine); err != nil {
		return nil, fmt.Errorf("error parsing header: %w", err)
	}
	// read contacts
	count := int64(0)
	contacts := make([]*Contact, 0)
	for file.Scan() {
		line := file.Text()
		contact := NewContact()
		err := fixedwidth.Unmarshal([]byte(line), contact)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, contact)
		count++
	}
	if err := file.Err(); err != nil {
		return nil, err
	}
	if err := header.Validate("CONTATOS", count); err != nil {
		return nil, err
	}
	return contacts, nil
}

// GetParsedFile retrieves and maps Conccred records from a file.
func (c *Contact) GetParsedFile(file *bufio.Scanner) (map[string]port.Report, error) {
	fileConccred, err := c.ParseContactFile(file)
	if err != nil {
		return nil, err
	}
	mapConccred := make(map[string]port.Report)
	for _, conc := range fileConccred {
		mapConccred[conc.GetKey()] = conc
	}
	return mapConccred, nil
}
