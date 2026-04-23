package ports

import (
	"time"

	"fuser/internal/core/domain"
)

type Repository interface {
	Connect(dns string) error
	Ping() error
	Close() error
	GetManagementTransactions(dt_transaction time.Time) ([]*domain.Management, error)
	GetExchangeTransactions(dt_transaction time.Time) ([]*domain.Exchange, error)
	GetTransactionsByKey(keys []string) ([]*domain.Transaction, error)
	GetTransactionsByDateRangeAndStatus(start, end time.Time, status int) ([]*domain.Transaction, error)
	InsertTransactions(transactions []*domain.Transaction) error
}
