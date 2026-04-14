package ports

import (
	"time"

	source_domain "consolidator/internal/core/domain/source"
	target_domain "consolidator/internal/core/domain/target"

)

// Repository defines the interface for data access operations related to transactions and associated entities.
type Repository interface {
	GetTransactionsByDate(date time.Time) ([]*source_domain.Transaction, error)
	SaveDesconto(desconto []*target_domain.Desconto) error
}
