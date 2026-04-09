package ports

import (
	"time"

	"fuser/internal/core/domain"
)

type Repository interface {
	Connect(dns string) error
	Ping() error
	GetManagementTransactions(dt_transaction time.Time) ([]*domain.Management, error)
	GetWebserviceTransactions(dt_transaction time.Time, page int) ([]*domain.Webservice, error)
	GetIntercamTransactions(dt_transaction time.Time) ([]*domain.Intercam, error)
	GetTransactionsByKey(keys []string) ([]*domain.Transaction, error)
	InsertTransactions(transactions []*domain.Transaction) error
}
