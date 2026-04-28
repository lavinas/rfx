package target_domain

import (
	"fmt"
	"maps"
	"slices"
	"time"

	source_domain "consolidator/internal/core/domain/source"
	"consolidator/internal/core/ports"
)

// RankingItem represents the ranking of establishments based on transaction data.
type RankingItem struct {
	ID                  int64     `gorm:"column:id"`
	CreatedAt           time.Time `gorm:"column:created_at"`
	UpdatedAt           time.Time `gorm:"column:updated_at"`
	Year                int       `gorm:"column:year"`
	Quarter             int       `gorm:"column:quarter"`
	EstablishmentCode   int64     `gorm:"column:establishment_code"`
	Function            string    `gorm:"column:function"`
	Brand               int       `gorm:"column:brand"`
	CaptureMode         int       `gorm:"column:capture_mode"`
	Installments        int       `gorm:"column:installments"`
	SegmentCode         int       `gorm:"column:segment_code"`
	TransactionAmount   float64   `gorm:"column:transaction_amount"`
	TransactionQuantity int       `gorm:"column:transaction_quantity"`
	AvgMccFee           float64   `gorm:"column:avg_mcc_fee"`
}

// Ranking represents the consolidated ranking data for a specific year and quarter.
type Ranking struct {
	DomainBase
	Ranking       *RankingItem
	consolidation map[string]*RankingItem
}

// NewRanking creates a new instance of Ranking.
func NewRanking() *Ranking {
	return &Ranking{
		Ranking:       &RankingItem{},
		consolidation: make(map[string]*RankingItem),
	}
}

// TableName returns the name of the database table for the RankingItem model.
func (i *RankingItem) TableName() string {
	return "ranking"
}

// GetFromTransaction returns a RankingItem instance populated with data from a given transaction.
func (i *RankingItem) GetFromTransaction(transaction *source_domain.Transaction) *RankingItem {
	return &RankingItem{
		Year:                transaction.GetYear(),
		Quarter:             transaction.GetQuarter(),
		EstablishmentCode:   transaction.GetEstablishmentCode(),
		Function:            transaction.GetFunctionCode(),
		Brand:               transaction.GetBrandCode(),
		CaptureMode:         transaction.GetCaptureModeCode(),
		Installments:        transaction.GetInstallments(),
		SegmentCode:         transaction.GetSegmentCode(),
		TransactionAmount:   transaction.GetTransactionAmount(),
		TransactionQuantity: 1,
		AvgMccFee:           transaction.GetRevenueMDRValueRate(),
	}
}

// GetKey generates a unique key for the RankingItem struct based on its fields.
func (i *RankingItem) GetKey() string {
	return fmt.Sprintf("%d-%d-%d-%s-%d-%d-%d-%d", i.Year, i.Quarter, i.EstablishmentCode, i.Function, i.Brand, i.CaptureMode, i.Installments, i.SegmentCode)
}

// Delete removes the consolidated data for a specific year and quarter from the consolidation map.
func (i *Ranking) Delete(year int, quarter int, repository ports.Repository) error {
	// delete the consolidated data for the specified year and quarter from the repository
	if err := repository.Delete(&RankingItem{}, year, quarter); err != nil {
		return err
	}
	return nil
}

// Save persists the consolidated data for a specific year and quarter to the repository.
func (i *Ranking) Save(repository ports.Repository) error {
	if err := repository.Save(slices.Collect(maps.Values(i.consolidation))); err != nil {
		return err
	}
	return nil
}

// GetFromTransactions processes a slice of transactions and returns a map of Ranking instances keyed by their unique keys.
func (i *Ranking) AddTransactions(transactions []*source_domain.Transaction) {
	// for each transaction, get the corresponding Ranking instance and update the transaction amount, quantity and average mcc fee
	for _, t := range transactions {
		ranking := i.Ranking.GetFromTransaction(t)
		key := ranking.GetKey()
		if existing, exists := i.consolidation[key]; exists {
			existing.TransactionAmount += ranking.TransactionAmount
			existing.TransactionQuantity += ranking.TransactionQuantity
			delta := ranking.AvgMccFee - existing.AvgMccFee
			existing.AvgMccFee += delta / float64(existing.TransactionQuantity)
		} else {
			i.consolidation[key] = ranking
		}
	}
}
