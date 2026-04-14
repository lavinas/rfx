package ports

type Service interface {
	Run(year int, quarter int, focus string) error
}
