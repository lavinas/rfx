package ports

import "time"

type Service interface {
	Run(start_date time.Time, end_date time.Time, focus string, leftover bool) error	
}
