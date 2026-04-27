package ports

import (
	"time"

	source_domain "consolidator/internal/core/domain/source"
)

// Repository defines the interface for data access operations related to transactions and associated entities.
type Repository interface {
	GetTransactionsByDate(date time.Time) ([]*source_domain.Transaction, error)
	GetBins() ([]*source_domain.Bin, error)
	GetEstablishments() ([]*source_domain.Establishment, error)
	GetTerminals() ([]*source_domain.Terminal, error)
	Delete(model interface{}, year int, quarter int) error
	Save(model interface{}) error
}
