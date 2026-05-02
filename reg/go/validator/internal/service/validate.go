package service

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/text/encoding/charmap"
	"validator/internal/domain"
	"validator/internal/port"
)

// ValidatorService represents the use case for checking or validating data
type ValidatorService struct {
	repo    port.Repository
	year    int
	quarter int
	// Add any dependencies or configurations needed for the use case
}

// NewValidatorService creates a new instance of ValidatorService
func NewValidatorService(repo port.Repository) *ValidatorService {
	return &ValidatorService{
		repo: repo,
	}
}

// ExecuteAll executes the check use case for all reports
func (uc *ValidatorService) ExecuteAll(year int, quarter int, path string) error {
	files := []string{
		"RANKING.TXT",
		"CONCCRED.TXT",
		"INFRESTA.TXT",
		"INFRTERM.TXT",
		"DESCONTO.TXT",
		"INTERCAM.TXT",
		"SEGMENTO.TXT",
		"LUCRCRED.TXT",
		"CONTATOS.TXT",
	}
	reports := []port.Report{
		domain.NewRanking(),
		domain.NewConccred(),
		domain.NewInfresta(),
		domain.NewInfrterm(),
		domain.NewDiscount(),
		domain.NewIntercam(),
		domain.NewSegment(),
		domain.NewLucrCred(),
		domain.NewContact(),
	}
	// validate path
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("path does not exist: %s", path)
	}
	// execute reports
	for i, file := range files {
		filename := fmt.Sprintf("%s/%s", path, file)
		if err := uc.ExecuteReport(reports[i], filename, year, quarter); err != nil {
			return err
		}
	}
	return nil
}

// ExecuteReport executes the check use case for a specific report
func (uc *ValidatorService) ExecuteReport(report port.Report, filename string, year int, quarter int) error {
	fmt.Printf("Reconciliating %s\n", filename)
	defer fmt.Println("---------------------------------------------------------------------------------------------------------")
	// Get db data
	loaded, err := report.GetDB(uc.repo, year, quarter)
	if err != nil {
		return fmt.Errorf("Error loading report data: %v", err)
	}
	// Get file data
	filed, err := uc.getFile(report, filename)
	if err != nil {
		return fmt.Errorf("Error parsing report file: %v", err)
	}
	// validate DB
	if errs := uc.validateReport(loaded); errs != nil {
		return fmt.Errorf("DB validation errors found for %s: %w", filename, errors.Join(errs...))
	}
	// validate File
	if errs := uc.validateReport(filed); errs != nil {
		return fmt.Errorf("File validation errors found for %s: %w", filename, errors.Join(errs...))
	}
	// Match and report discrepancies
	if errs := uc.match(loaded, filed); errs != nil {
		return fmt.Errorf("Discrepancies found in %s: %w", filename, errors.Join(errs...))
	}
	fmt.Printf("No discrepancies found for %s\n", filename)
	return nil
}

// getfile retrieves and parses the file data for a given report and filename, returning a map of records keyed by their unique identifiers.
func (uc *ValidatorService) getFile(report port.Report, filename string) (map[string]port.Report, error) {
	file, err := uc.OpenFileCaseInsensitive(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := charmap.ISO8859_1.NewDecoder()
	decodedReader := decoder.Reader(file)
	scanner := bufio.NewScanner(decodedReader)
	return report.GetParsedFile(scanner)
}

// OpenFileCaseInsensitive attempts to open a file with case-insensitive matching of the filename.
func (uc *ValidatorService) OpenFileCaseInsensitive(path string) (*os.File, error) {
	// 1. Try a direct open first (efficient for Windows/Mac)
	f, err := os.Open(path)
	if err == nil {
		return f, nil
	}

	// 2. If it fails, search the parent directory
	dir := filepath.Dir(path)
	targetBase := filepath.Base(path)

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		// EqualFold is the Go way to compare strings case-insensitively
		if strings.EqualFold(entry.Name(), targetBase) {
			realPath := filepath.Join(dir, entry.Name())
			return os.Open(realPath)
		}
	}

	return nil, fmt.Errorf("file not found: %s", path)
}

// validate validates records from both sources.
func (uc *ValidatorService) validateReport(report map[string]port.Report) []error {
	var errs []error
	for key, dbRecord := range report {
		if err := dbRecord.Validate(); err != nil {
			errs = append(errs, fmt.Errorf("validation error for DB record with key %s: %v", key, err))
		}
	}
	return errs
}

// match compares two maps of records and returns a slice of errors for any discrepancies found.
func (uc *ValidatorService) match(db map[string]port.Report, file map[string]port.Report) []error {
	var errs []error
	// compare lengths
	if len(db) != len(file) {
		return []error{fmt.Errorf("length mismatch: DB has %d records, File has %d records", len(db), len(file))}
	}
	errs = uc.matchDB(db, file)
	errs = append(errs, uc.matchFile(db, file)...)
	return errs
}

// matchDB compares the records from the database and the file, returning a slice of errors for any discrepancies found.
func (uc *ValidatorService) matchDB(db map[string]port.Report, file map[string]port.Report) []error {
	var errs []error
	encoder := charmap.ISO8859_1.NewEncoder()

	for key, dbRecord := range db {
		fileRecord, exists := file[key]
		if !exists {
			errs = append(errs, fmt.Errorf("db record with key %s exists in DB but not in file", key))
			continue
		}
		encodedDBBytes, err := encoder.Bytes([]byte(dbRecord.String()))
		if err != nil {
			errs = append(errs, fmt.Errorf("error encoding DB record with key %s: %v", key, err))
			continue
		}
		encoderFileBytes, err := encoder.Bytes([]byte(fileRecord.String()))
		if err != nil {
			errs = append(errs, fmt.Errorf("error encoding File record with key %s: %v", key, err))
			continue
		}
		encodedDBString := string(encodedDBBytes)
		encodedFileString := string(encoderFileBytes)
		if encodedDBString != encodedFileString {
			errs = append(errs, fmt.Errorf("mismatch for key %s:\nDB: %s\nFile: %s", key, encodedDBString, encodedFileString))
		}
	}
	return errs

}

// matchFileDB compares the records from the file and the database, returning a slice of errors for any discrepancies found.
func (uc *ValidatorService) matchFile(db map[string]port.Report, file map[string]port.Report) []error {
	var errs []error
	encoder := charmap.ISO8859_1.NewEncoder()

	for key, fileRecord := range file {
		dbRecord, exists := db[key]
		if !exists {
			errs = append(errs, fmt.Errorf("filerecord with key %s exists in File but not in DB", key))
			continue
		}
		encodedDBBytes, err := encoder.Bytes([]byte(dbRecord.String()))
		if err != nil {
			errs = append(errs, fmt.Errorf("error encoding DB record with key %s: %v", key, err))
			continue
		}
		encoderFileBytes, err := encoder.Bytes([]byte(fileRecord.String()))
		if err != nil {
			errs = append(errs, fmt.Errorf("error encoding File record with key %s: %v", key, err))
			continue
		}
		encodedDBString := string(encodedDBBytes)
		encodedFileString := string(encoderFileBytes)
		if encodedDBString != encodedFileString {
			errs = append(errs, fmt.Errorf("mismatch for key %s:\nDB: %s\nFile: %s", key, encodedDBString, encodedFileString))
		}
	}
	return errs
}

