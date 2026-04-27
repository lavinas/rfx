package target_domain

import (
	"fmt"
	"maps"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"

	source_domain "consolidator/internal/core/domain/source"
	"consolidator/internal/core/ports"
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

// Segmento represents the consolidated segment data for a specific year and quarter.
type Segmento struct {
	DomainBase
	Segmento      *SegmentoItem
	consolidation map[string]*SegmentoItem
}

// NewSegmento creates a new instance of Segmento.
func NewSegmento() *Segmento {
	return &Segmento{
		Segmento:      &SegmentoItem{},
		consolidation: make(map[string]*SegmentoItem),
	}
}

// TableName specifies the table name for SegmentoItem struct
func (i *SegmentoItem) TableName() string {
	return "segmento"
}

// Delete removes the consolidated data for a specific year and quarter from the consolidation map.
func (s *Segmento) Delete(year int, quarter int, repository ports.Repository) error {
	// delete the consolidated data for the specified year and quarter from the repository
	if err := repository.Delete(&SegmentoItem{}, year, quarter); err != nil {
		return err
	}
	return nil
}

// Save persists the consolidated data for a specific year and quarter to the repository.
func (s *Segmento) Save(repository ports.Repository) error {
	if err := repository.Save(slices.Collect(maps.Values(s.consolidation))); err != nil {
		return err
	}
	return nil
}

// AddTransactions adds the transaction amount and quantity from another Segmento to the current one.
func (s *Segmento) AddTransactions(transactions []*source_domain.Transaction) {

	// Iterate over the transactions and update the Segmento instance with the segment code, segment name and description based on the mcc code.
	for _, transaction := range transactions {
		// Get the segment code and mcc code from the transaction. If either of them is zero, skip the transaction.
		segmentCode := transaction.GetSegmentCode()
		mccCode := transaction.GetMccCode()

		// If the segment code or mcc code is zero, skip the transaction.
		if segmentCode == 0 || mccCode == 0 {
			continue
		}

		// Generate a key for the segment code to be used in the items map.
		key := strconv.Itoa(segmentCode)

		// Create a new Segmento instance from the transaction data if the key does not exist in the items map and continue to the next transaction.
		segment, exists := s.consolidation[key]
		if !exists {
			s.consolidation[key] = &SegmentoItem{
				Year:        transaction.GetYear(),
				Quarter:     transaction.GetQuarter(),
				SegmentCode: segmentCode,
				SegmentName: transaction.GetSegmentName(),
				Description: s.mountDescription(mccCode, ""),
			}
			continue
		}

		// If the key already exists in the items map, update the description of the existing Segmento instance with the new mcc code.
		segment.Description = s.mountDescription(mccCode, segment.Description)
	}

}

// mountDescription updates the description of a Segmento instance with a new mcc code. If the description already contains an mcc code, it appends the new mcc code to the existing description. If the description does not contain an mcc code, it sets the description to the new mcc code.
func (s *Segmento) mountDescription(mccCode int, description string) string {
	// If the description is empty, set the description to the new mcc code and return it.
	if description == "" {
		return fmt.Sprintf("MCC: %d", mccCode)
	}

	// If the description already contains the new mcc code, return the existing description.
	if strings.Contains(description, strconv.Itoa(mccCode)) {
		return description
	}

	// If the description already contains an mcc code, append the new mcc code to the existing description and return it.
	prefs := strings.Split(description, ": ")
	mccs := strings.Split(prefs[1], ", ")
	mccs = append(mccs, strconv.Itoa(mccCode))

	// sort the mcc codes in the description and return the updated description.
	sort.Slice(mccs, func(i, j int) bool {
		// mccCodeI, _ := strconv.Atoi(mccs[i])
		// mccCodeJ, _ := strconv.Atoi(mccs[j])
		return mccs[i] < mccs[j]
	})

	return fmt.Sprintf("%s: %s", prefs[0], strings.Join(mccs, ", "))

}
