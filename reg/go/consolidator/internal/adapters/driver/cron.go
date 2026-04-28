package driver

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/robfig/cron/v3"

	"consolidator/internal/core/ports"
	"github.com/postfinance/single"
)

// CronDriver is a driver that implements the necessary interfaces to run the FuseService on a schedule
type CronDriver struct {
	service       ports.Service
	logger        ports.Logger
	schedule      []string
	timezone      string
	backtrackDays int
}

// NewCronDriver creates a new instance of CronDriver
func NewCronDriver(service ports.Service, logger ports.Logger, config ports.Config) *CronDriver {
	var schedule []string
	var timezone string
	var backtrackDays int
	config.GetCronData(&schedule, &timezone, &backtrackDays)
	return &CronDriver{
		service:       service,
		logger:        logger,
		schedule:      schedule,
		timezone:      timezone,
		backtrackDays: backtrackDays,
	}
}

// Run executes the main logic of the CronDriver by calling the Run method of the service
func (d *CronDriver) Run() error {
	// Use single instance lock to ensure only one instance of CronDriver is running at a time
	s, err := single.New("fuser-cron-driver")
	if err != nil {
		return fmt.Errorf("error creating single instance lock: %w", err)
	}
	if err := s.Lock(); err != nil {
		return fmt.Errorf("another instance of CronDriver is already running: %w", err)
	}
	defer s.Unlock()
	return d.runCron()
}

// runCron sets up the cron scheduler and adds the function to be executed on the specified schedule, then starts the scheduler and waits for an interrupt signal to gracefully shut down
func (d *CronDriver) runCron() error {
	// Create a new cron scheduler and add the function to be executed on the specified schedule, then start the scheduler and wait for an interrupt signal to gracefully shut down
	c := cron.New()
	if loc, err := time.LoadLocation(d.timezone); err == nil {
		c = cron.New(cron.WithLocation(loc))
	}
	if err := d.addFunc(c); err != nil {
		return fmt.Errorf("failed to add cron function: %w", err)
	}
	// Start the cron scheduler and wait for an interrupt signal to gracefully shut down
	d.logger.Println("Starting CronDriver with schedule:", d.schedule)
	c.Start()
	// Configurate listener
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	sig := <-stop
	d.logger.Printf("Stop Signal [%v] received", sig)
	ctx := c.Stop()
	<-ctx.Done()
	d.logger.Println("CronDriver stopped")
	return nil
}

// addFunc adds a function to the cron scheduler with the specified schedule and logger
func (d *CronDriver) addFunc(cron *cron.Cron) error {
	if len(d.schedule) == 0 {
		return fmt.Errorf("no cron schedules found in configuration")
	}
	for _, schedule := range d.schedule {
		_, err := cron.AddFunc(schedule, func() {
			d.callService()
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// callService calls the Run method of the service with the provided parameters and logs the result
func (d *CronDriver) callService() {
	d.logger.Println("Starting scheduled task")
	if err := d.service.Run(2026, 1, 1); err != nil {
		d.logger.Println("Error running service:", err)
		return
	}
	d.logger.Println("Scheduled fuser completed successfully")
}
