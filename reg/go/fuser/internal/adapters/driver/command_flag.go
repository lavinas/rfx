package driver

import (
	"flag"
	"fmt"
	"time"

	"fuser/internal/core/ports"
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
	start_date := flag.String("start", "", "Start date for processing transactions (format: YYYY-MM-DD)")
	end_date := flag.String("end", "", "End date for processing transactions (format: YYYY-MM-DD)")
	focus := flag.String("focus", "all", "Focus of the processing (options: all, transactions, accounts)")
	
	flag.Parse()

	sd, err := time.Parse("2006-01-02", *start_date)
	if err != nil {
		return fmt.Errorf("invalid command: use fuser -start=yyyy-mm-dd -end=yyyy-mm-dd")
	}
	ed, err := time.Parse("2006-01-02", *end_date)
	if err != nil {
		return fmt.Errorf("invalid command: use fuser -start=yyyy-mm-dd -end=yyyy-mm-dd")
	}
	if *focus != "all" && *focus != "intercam" && *focus != "management" {
		return fmt.Errorf("invalid focus: use fuser -focus=all|intercam|management")
	}
	return d.service.Run(sd, ed, *focus)
}
