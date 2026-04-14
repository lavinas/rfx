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
	// Define valid focus options
	focusMap := map[string]bool{
		"all":      true,
		"conccred": true,
		"desconto": true,
		"infresta": true,
		"infrterm": true,
		"luccred":  true,
		"ranking":  true,
		"segmento": true,
	}
	// parse parameters
	year := flag.Int("year", 0, "year for processing transactions (format: YYYY)")
	quarter := flag.Int("quarter", 0, "quarter for processing transactions (format: 1, 2, 3, 4)")
	focus := flag.String("focus", "all", "Focus of the processing (options: all, conccred, desconto, infresta, infrterm, luccred, ranking, segmento)")
	flag.Parse()

	// validate parameters
	if *year < 2000 || *year > time.Now().Year() {
		return fmt.Errorf("invalid year: use consolidate -year=YYYY -quarter=(1, 2, 3, or 4) -focus=all|conccred|desconto|infresta|infrterm|luccred|ranking|segmento")
	}
	if *quarter < 1 || *quarter > 4 {
		return fmt.Errorf("invalid quarter: use consolidate -year=YYYY -quarter=Q (1, 2, 3, or 4) -focus=all|conccred|desconto|infresta|infrterm|luccred|ranking|segmento")
	}
	if !focusMap[*focus] {
		return fmt.Errorf("invalid focus: use consolidate -year=YYYY -quarter=Q (1, 2, 3, or 4) -focus=all|conccred|desconto|infresta|infrterm|luccred|ranking|segmento")
	}

	// run service
	return d.service.Run(*year, *quarter, *focus)
}
