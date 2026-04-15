package ports

import (
	"time"

	source_domain "consolidator/internal/core/domain/source"
	target_domain "consolidator/internal/core/domain/target"
)

// Repository defines the interface for data access operations related to transactions and associated entities.
type Repository interface {
	GetTransactionsByDate(date time.Time) ([]*source_domain.Transaction, error)
	GetBins() ([]*source_domain.Bin, error)
	SaveDesconto(desconto []*target_domain.Desconto) error
	DeleteDesconto(year int, quarter int) error
	SaveRanking(ranking []*target_domain.Ranking) error
	DeleteRanking(year int, quarter int) error
	SaveIntercam(intercam []*target_domain.Intercam) error
	DeleteIntercam(year int, quarter int) error
	SaveConcCred(conccred []*target_domain.ConcCred) error
	DeleteConcCred(year int, quarter int) error
}
