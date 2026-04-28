package driver

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
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
	// Run the service in a separate goroutine and wait for an interrupt signal to gracefully shut down
	var err error
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	// Run the service in a separate goroutine to allow for graceful shutdown on interrupt signal
	go func() {
		err= d.callService()
		sigs <- os.Interrupt
	}()
	// Wait for an interrupt signal to gracefully shut down the service
	<-sigs
	return err
}

// Run executes the main logic of the FlagDriver by calling the Run method of the service
func (d *FlagDriver) callService() error {
	// parse parameters
	start_date := flag.String("start", "", "Start date for processing transactions (format: YYYY-MM-DD)")
	end_date := flag.String("end", "", "End date for processing transactions (format: YYYY-MM-DD)")
	focus := flag.String("focus", "all", "Focus of the processing (options: all, exchange, management, none)")
	leftover := flag.Bool("leftover", true, "Whether to process leftover transactions (default: true)")
	flag.Parse()
	// validate start and end dates
	sd, err := time.Parse("2006-01-02", *start_date)
	if err != nil {
		return fmt.Errorf("invalid command: use fuser -start=yyyy-mm-dd -end=yyyy-mm-dd")
	}
	ed, err := time.Parse("2006-01-02", *end_date)
	if err != nil {
		return fmt.Errorf("invalid command: use fuser -start=yyyy-mm-dd -end=yyyy-mm-dd")
	}
	if *focus != "all" && *focus != "exchange" && *focus != "management" && *focus != "none" {
		return fmt.Errorf("invalid focus: use fuser -focus=all|exchange|management|none")
	}
	//
	return d.service.Run(sd, ed, *focus, *leftover)
}
