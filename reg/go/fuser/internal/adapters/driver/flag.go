package driver

import (
	"flag"
	"fmt"
	"time"

	"fuser/internal/ports"
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
	start_date := flag.String("start_date", "", "Start date for processing transactions (format: YYYY-MM-DD)")
	end_date := flag.String("end_date", "", "End date for processing transactions (format: YYYY-MM-DD)")
	flag.Parse()

	sd, err := time.Parse("2006-01-02", *start_date)
	if err != nil {
		return fmt.Errorf("invalid command: use fuser -start_date=yyyy-mm-dd -end_date=yyyy-mm-dd")
	}
	ed, err := time.Parse("2006-01-02", *end_date)
	if err != nil {
		return fmt.Errorf("invalid command: use fuser -start_date=yyyy-mm-dd -end_date=yyyy-mm-dd")
	}
	return d.service.Run(sd, ed)
}
