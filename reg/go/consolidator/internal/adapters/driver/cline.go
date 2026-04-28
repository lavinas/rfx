package driver

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
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
	// Run the service in a separate goroutine and wait for an interrupt signal to gracefully shut down
	var err error
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	// Run the service in a separate goroutine to allow for graceful shutdown on interrupt signal
	go func() {
		err = d.callService()
		sigs <- os.Interrupt
	}()
	// Wait for an interrupt signal to gracefully shut down the service
	<-sigs
	return err
}

// callService executes the main logic of the FlagDriver by calling the Run method of the service
func (d *FlagDriver) callService() error {
	// parse parameters
	year := flag.Int("year", 0, "year for processing transactions (format: YYYY)")
	quarter := flag.Int("quarter", 0, "quarter for processing transactions (format: 1, 2, 3, 4)")
	days := flag.Int("days", -1, "number of days to processing into the quarter (optional)")
	flag.Parse()

	// validate year
	if *year < 2000 || *year > time.Now().Year() {
		return fmt.Errorf("invalid year: use consolidate -year=YYYY -quarter=(1, 2, 3, or 4)")
	}

	// validate quarter
	if *quarter < 1 || *quarter > 4 {
		return fmt.Errorf("invalid quarter: use consolidate -year=YYYY -quarter=(1, 2, 3, or 4)")
	}

	// validate days
	if *days != -1 && *days <= 0 {
		return fmt.Errorf("invalid days: use consolidate -year=YYYY -quarter=(1, 2, 3, or 4) -days=N (where N is a positive integer)")
	}

	// run service
	return d.service.Run(*year, *quarter, *days)
}
