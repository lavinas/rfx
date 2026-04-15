package ports

import "time"

type Service interface {
	 Run(year int, quarter int, delete bool, filter_ranking bool, start *time.Time, end *time.Time) error
}
