package ports

type Service interface {
	Run(year int, quarter int, days int) error
}
