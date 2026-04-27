package target_domain

import (
	"fmt"
	"maps"
	"slices"
	"time"

	source_domain "consolidator/internal/core/domain/source"
	"consolidator/internal/core/ports"
	"math"
)

// DescontoItem represents the data structure for discounts which will be used for fusing data between intercam, management and webservice
type DescontoItem struct {
	ID                  int64     `gorm:"column:id"`
	CreatedAt           time.Time `gorm:"column:created_at"`
	UpdatedAt           time.Time `gorm:"column:updated_at"`
	Year                int       `gorm:"column:year"`
	Quarter             int       `gorm:"column:quarter"`
	Function            string    `gorm:"column:function"`
	Brand               int       `gorm:"column:brand"`
	CaptureMode         int       `gorm:"column:capture_mode"`
	Installments        int       `gorm:"column:installments"`
	SegmentCode         int       `gorm:"column:segment_code"`
	AvgMDRFee           float64   `gorm:"column:avg_mdr_fee"`
	MinMDRFee           float64   `gorm:"column:min_mdr_fee"`
	MaxMDRFee           float64   `gorm:"column:max_mdr_fee"`
	StdevMDRFee         float64   `gorm:"column:stdev_mdr_fee"`
	SqrdiffMDRFee       float64   `gorm:"column:sqrdiff_mdr_fee"`
	TransactionAmount   float64   `gorm:"column:transaction_amount"`
	TransactionQuantity int64     `gorm:"column:transaction_quantity"`
}

// Desconto represents the consolidated discount transactions for a specific year and quarter.
type Desconto struct {
	DomainBase
	desconto      *DescontoItem
	consolidation map[string]*DescontoItem
}

// NewDesconto creates a new instance of Desconto.
func NewDesconto() *Desconto {
	return &Desconto{
		desconto:      &DescontoItem{},
		consolidation: make(map[string]*DescontoItem),
	}
}

// TableName specifies the table name for DescontoItem struct
func (i *DescontoItem) TableName() string {
	return "cadoc_6334_v2.desconto"
}

// GetKey generates a unique key for the DescontoItem struct based on its fields.
func (i *DescontoItem) GetKey() string {
	return fmt.Sprintf("%d-%d-%s-%d-%d-%d-%d", i.Year, i.Quarter, i.Function, i.Brand, i.CaptureMode, i.Installments, i.SegmentCode)
}

// GetFromTransaction returns the discount data for a given transaction.
func (i *DescontoItem) GetFromTransaction(transaction *source_domain.Transaction) *DescontoItem {
	return &DescontoItem{
		Year:                transaction.GetYear(),
		Quarter:             transaction.GetQuarter(),
		Function:            transaction.GetFunctionCode(),
		Brand:               transaction.GetBrandCode(),
		CaptureMode:         transaction.GetCaptureModeCode(),
		Installments:        transaction.GetInstallments(),
		SegmentCode:         transaction.GetSegmentCode(),
		AvgMDRFee:           transaction.GetRevenueMDRValueRate(),
		MinMDRFee:           transaction.GetRevenueMDRValueRate(),
		MaxMDRFee:           transaction.GetRevenueMDRValueRate(),
		StdevMDRFee:         0,
		SqrdiffMDRFee:       0,
		TransactionAmount:   transaction.GetTransactionAmount(),
		TransactionQuantity: 1,
	}
}

// Delete removes the consolidated data for a specific year and quarter from the consolidation map.
func (i *Desconto) Delete(year int, quarter int, repository ports.Repository) error {
	// delete the consolidated data for the specified year and quarter from the repository
	if err := repository.Delete(&DescontoItem{}, year, quarter); err != nil {
		return err
	}
	return nil
}

// Save persists the consolidated data for a specific year and quarter to the repository.
func (i *Desconto) Save(repository ports.Repository) error {
	if err := repository.Save(slices.Collect(maps.Values(i.consolidation))); err != nil {
		return err
	}
	return nil
}

// GetFromTransactions returns a map of Desconto structs for a given list of transactions.
func (i *Desconto) AddTransactions(transactions []*source_domain.Transaction) {
	// for each transaction, get the corresponding Desconto instance and update the transaction amount, quantity and mdr fee statistics
	for _, t := range transactions {
		desconto := i.desconto.GetFromTransaction(t)
		key := desconto.GetKey()

		// if the key already exists in the items map, update the existing Desconto instance with the new transaction data
		if existing, exists := i.consolidation[key]; exists {
			// sum amount and quantity
			existing.TransactionAmount += desconto.TransactionAmount
			existing.TransactionQuantity += desconto.TransactionQuantity
			// update min and max MDR fees
			if desconto.MinMDRFee < existing.MinMDRFee {
				existing.MinMDRFee = desconto.MinMDRFee
			}
			if desconto.MaxMDRFee > existing.MaxMDRFee {
				existing.MaxMDRFee = desconto.MaxMDRFee
			}
			// update average and standard deviation using Welford's online algorithm
			delta := desconto.AvgMDRFee - existing.AvgMDRFee
			existing.AvgMDRFee += delta / float64(existing.TransactionQuantity)
			delta2 := desconto.AvgMDRFee - existing.AvgMDRFee
			existing.SqrdiffMDRFee += delta2 * delta2
			variance := existing.SqrdiffMDRFee / (float64(existing.TransactionQuantity) - 1)
			existing.StdevMDRFee = math.Sqrt(variance)
			i.consolidation[key] = existing
		} else {
			i.consolidation[key] = desconto
		}
	}
}
