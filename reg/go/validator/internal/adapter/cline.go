package adapter

import (
	"os"
	"os/signal"

	ports "validator/internal/port"
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

// Run executes the main logic of the FlagDriver by calling the Run method of the service
func (d *FlagDriver) callService() error {
	// parse parameters
	return d.service.ExecuteAll()
}
