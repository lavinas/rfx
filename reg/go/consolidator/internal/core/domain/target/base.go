package target_domain

import (
	source_domain "consolidator/internal/core/domain/source"
	"consolidator/internal/core/ports"
)

// DomainBase is a base struct for domain entities in the consolidator service. It can be extended by specific domain entities to inherit common functionality or properties.
type DomainBase struct {
}

// Save is a method that can be overridden by specific domain entities to implement the logic for saving consolidated data to the repository. By default, it does nothing and returns nil.
func (d *DomainBase) Save(repository ports.Repository) error {
	return nil
}

// Delete is a method that can be overridden by specific domain entities to implement the logic for deleting consolidated data from the repository. By default, it does nothing and returns nil.
func (d *DomainBase) Delete(year int, quarter int, repository ports.Repository) error {
	return nil
}

// AddTransactions is a method that can be overridden by specific domain entities to implement the logic for adding transactions to the consolidated data. By default, it does nothing.
func (d *DomainBase) AddTransactions(transactions []*source_domain.Transaction) {
}

// AddEstablishments is a method that can be overridden by specific domain entities to implement the logic for adding establishments to the consolidated data. By default, it does nothing.
func (d *DomainBase) AddEstablishments(year int, quarter int, establishments []*source_domain.Establishment) {
}

// AddTerminals is a method that can be overridden by specific domain entities to implement the logic for adding terminals to the consolidated data. By default, it does nothing.
func (d *DomainBase) AddTerminals(year int, quarter int, terminals []*source_domain.Terminal, estabMap map[int64]string) {
}

// Build is a method that can be overridden by specific domain entities to implement the logic for building or processing the consolidated data after all transactions, establishments, and terminals have been added. By default, it does nothing.
func (d *DomainBase) Build() {
}
