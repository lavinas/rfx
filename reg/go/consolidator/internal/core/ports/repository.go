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
	GetEstablishments() ([]*source_domain.Establishment, error)
	GetTerminals() ([]*source_domain.Terminal, error)
	GetConcCred(year int, quarter int) ([]*target_domain.ConcCred, error)
	Delete(model interface{}, year int, quarter int) error
	Save(model interface{}) error

	SaveDesconto(desconto []*target_domain.Desconto) error
	DeleteDesconto(year int, quarter int) error
	SaveRanking(ranking []*target_domain.Ranking) error
	DeleteRanking(year int, quarter int) error
	SaveRankingFiltered(rankingFiltered []*target_domain.RankingFiltered) error
	DeleteRankingFiltered(year int, quarter int) error
	SaveIntercam(intercam []*target_domain.Intercam) error
	DeleteIntercam(year int, quarter int) error
	SaveConcCred(conccred []*target_domain.ConcCred) error
	DeleteConcCred(year int, quarter int) error
	SaveSegmento(segmento []*target_domain.Segmento) error
	DeleteSegmento(year int, quarter int) error
	SaveLuccred(luccred []*target_domain.Luccred) error
	DeleteLuccred(year int, quarter int) error
	SaveInfresta(infresta []*target_domain.Infresta) error
	DeleteInfresta(year int, quarter int) error
	SaveInfrterm(infrterm []*target_domain.Infrterm) error
	DeleteInfrterm(year int, quarter int) error
}
