package driver

import (
	"flag"
	"fmt"
	"time"

	"consolidator/internal/core/ports"
)

// FlagDriver is a driver that implements the necessary interfaces to run the FuseService
type FlagDriver struct {
	service ports.Service
}

// NewFlagDriver creates a new instance of FlagDriver
func NewFlagDriver(service ports.Service) *FlagDriver {
	return &FlagDriver{
		service: service,
	}
}

// Run executes the main logic of the FlagDriver by calling the Run method of the service
func (d *FlagDriver) Run() error {
	// parse parameters
	year := flag.Int("year", 0, "year for processing transactions (format: YYYY)")
	quarter := flag.Int("quarter", 0, "quarter for processing transactions (format: 1, 2, 3, 4)")
	delete := flag.Bool("delete", true, "Delete existing consolidated data before processing")
	start := flag.String("start", "", "Start date for processing transactions (format: YYYY-MM-DD)")
	end := flag.String("end", "", "End date for processing transactions (format: YYYY-MM-DD)")
	flag.Parse()

	// validate parameters
	if *year < 2000 || *year > time.Now().Year() {
		return fmt.Errorf("invalid year: use consolidate -year=YYYY -quarter=(1, 2, 3, or 4)")
	}
	if *quarter < 1 || *quarter > 4 {
		return fmt.Errorf("invalid quarter: use consolidate -year=YYYY -quarter=Q (1, 2, 3, or 4)")
	}
	var st, ed time.Time
	var err error
	if *start != "" {
		if st, err = time.Parse("2006-01-02", *start); err != nil {
			return fmt.Errorf("invalid start date: use consolidate -year=YYYY -quarter=Q (1, 2, 3, or 4) -start=YYYY-MM-DD")
		}

	}
	if *end != "" {
		if ed, err = time.Parse("2006-01-02", *end); err != nil {
			return fmt.Errorf("invalid end date: use consolidate -year=YYYY -quarter=Q (1, 2, 3, or 4) -end=YYYY-MM-DD")
		}
	}

	// run service
	return d.service.Run(*year, *quarter, *delete, &st, &ed)
}
