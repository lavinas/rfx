package ports

import (
	source_domain "consolidator/internal/core/domain/source"
)

type Domain interface {
	Delete(year int, quarter int, repository Repository) error
	Save(repository Repository) error
	AddTransactions(transactions []*source_domain.Transaction)
	AddEstablishments(year int, quarter int, establishments []*source_domain.Establishment)
	AddTerminals(year int, quarter int, terminals []*source_domain.Terminal, esblishmentMap map[int64]string)
	Build()
}
