package ports

import "time"

type Service interface {
	Run(start_date, end_date time.Time) error
}
