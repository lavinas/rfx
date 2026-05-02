package domain

import (
	"bufio"
	"fmt"

	"github.com/ianlopshire/go-fixedwidth"
	"validator/internal/port"
)

// Segment represents a segment of a path.
type Segment struct {
	Year        int    `gorm:"column:year"`
	Quarter     int    `gorm:"column:quarter"`
	Name        string `fixed:"1,50" gorm:"column:segment_name"`
	Description string `fixed:"51,300" gorm:"column:segment_description"`
	CodeStr     string `fixed:"301,303"`
	Code        int64  `gorm:"column:segment_code"`
}

// TableName specifies the table name for Segment struct
func (s *Segment) TableName() string {
	return "segmento"
}

// NewSegment creates a new Segment instance
func NewSegment() *Segment {
	return &Segment{}
}

// GetName gets name of the report
func (s *Segment) GetName() string {
	return "SEGMENTO"
}

// Format marshals the Segment struct into a fixed-width format.
func (s *Segment) Format() string {
	ret := ""
	ret += fmt.Sprintf("%-50s", s.Name)
	ret += fmt.Sprintf("%-250s", s.Description)
	ret += fmt.Sprintf("%03d", s.Code)
	return ret
}

// Validate validates the Segment information.
func (s *Segment) Validate() error {
	if s.Name == "" {
		return fmt.Errorf("invalid name in segment")
	}
	if s.Description == "" {
		return fmt.Errorf("invalid description in segment")
	}
	if s.Code <= 0 {
		fmt.Println(1, s)
		return fmt.Errorf("invalid code in segment")
	}
	return nil
}

// GetKey generates a unique key for the Segment record.
func (s *Segment) GetKey() string {
	return fmt.Sprintf("%03d", s.Code)
}

// FindAll retrieves all Segment records.
func (s *Segment) GetDB(repo port.Repository, year int, quarter int) (map[string]port.Report, error) {
	var records []*Segment
	err := repo.FindAll(&records, 0, 0, "", "year = ? AND quarter = ?", year, quarter)
	if err != nil {
		return nil, err
	}
	ret := make(map[string]port.Report)
	for _, r := range records {
		name, err := RemoveAccents(r.Name)
		if err != nil {
			return nil, fmt.Errorf("error removing accents from name: %w", err)
		}
		r.Name = name
		ret[r.GetKey()] = r
	}
	return ret, nil
}

// Parse parses a fixed-width string into a Segment struct
func (s *Segment) Parse(line string) error {
	line, err := RemoveAccents(line)
	if err != nil {
		return fmt.Errorf("error removing accents: %w", err)
	}
	err = fixedwidth.Unmarshal([]byte(line), s)
	if err != nil {
		return err
	}
	if s.CodeStr == "" {
		s.Code = 0
		return nil
	}
	_, err = fmt.Sscanf(s.CodeStr, "%d", &s.Code)
	if err != nil {
		return fmt.Errorf("error parsing Code: %w", err)
	}
	return nil
}

// String returns a string representation of the Segment struct
func (s *Segment) String() string {
	return fmt.Sprintf("Name: %s, Description: %s, Code: %d", s.Name, s.Description, s.Code)
}

// ParseSegmentFile parses a file of segments into a slice of Segment structs
func (s *Segment) ParseSegmentFile(file *bufio.Scanner) ([]*Segment, error) {
	// read header
	if !file.Scan() {
		return nil, fmt.Errorf("file is empty")
	}
	headerLine := file.Text()
	header := &RankingHeader{}
	if _, err := header.Parse(headerLine); err != nil {
		return nil, fmt.Errorf("error parsing header: %w", err)
	}
	// read records
	segments := []*Segment{}
	count := 0
	for file.Scan() {
		line := file.Text()
		segment := &Segment{}
		err := segment.Parse(line)
		if err != nil {
			return nil, fmt.Errorf("error parsing line: %w", err)
		}
		segments = append(segments, segment)
		count++
	}
	if err := file.Err(); err != nil {
		return nil, err
	}
	if err := header.Validate("SEGMENTO", int64(count)); err != nil {
		return nil, err
	}
	return segments, nil
}

// GetParsedFile retrieves and maps Segment records from a file.
func (s *Segment) GetParsedFile(file *bufio.Scanner) (map[string]port.Report, error) {
	fileSegments, err := s.ParseSegmentFile(file)
	if err != nil {
		return nil, err
	}
	segmentMap := make(map[string]port.Report)
	for _, seg := range fileSegments {
		segmentMap[seg.GetKey()] = seg
	}
	return segmentMap, nil
}
