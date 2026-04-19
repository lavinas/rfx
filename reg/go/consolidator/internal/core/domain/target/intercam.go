package target_domain

import (
	"fmt"
	"time"

	source_domain "consolidator/internal/core/domain/source"
)

// Intercam represents the interchange fee for a given transaction.
type Intercam struct {
	ID                  int64                        `gorm:"column:id"`
	CreatedAt           time.Time                    `gorm:"column:created_at"`
	UpdatedAt           time.Time                    `gorm:"column:updated_at"`
	Year                int                          `gorm:"column:year"`
	Quarter             int                          `gorm:"column:quarter"`
	ProductCode         int                          `gorm:"column:product_code"`
	CardType            string                       `gorm:"column:card_type"`
	Function            string                       `gorm:"column:function"`
	Brand               int                          `gorm:"column:brand"`
	CaptureMode         int                          `gorm:"column:capture_mode"`
	Installments        int                          `gorm:"column:installments"`
	SegmentCode         int                          `gorm:"column:segment_code"`
	InterchangeFee      float64                      `gorm:"column:interchange_fee"`
	TransactionAmount   float64                      `gorm:"column:transaction_amount"`
	TransactionQuantity int64                        `gorm:"column:transaction_quantity"`
	Bins                map[int64]*source_domain.Bin `gorm:"-"`
}

// NewIntercam creates a new instance of Intercam with the provided BIN information.
func NewIntercam(bins map[int64]*source_domain.Bin) *Intercam {
	return &Intercam{
		Bins: bins,
	}
}

// TableName returns the name of the table in the database.
func (i *Intercam) TableName() string {
	return "cadoc_6334_v2.intercam"
}

// GetFromTransaction returns the interchange fee for a given transaction.
func (i *Intercam) GetFromTransaction(transaction *source_domain.Transaction) *Intercam {
	// Get the product code and card type from the BIN information, or use default values if not available
	product_code := source_domain.DefaultProductCode
	card_type := source_domain.DefaultCardType
	product, ok := i.Bins[transaction.GetBin()]
	if ok {
		product_code = product.GetProductCode()
		card_type = product.GetCardType()
	}
	// Create and return the Intercam struct based on the transaction data and BIN information
	return &Intercam{
		Year:                transaction.GetYear(),
		Quarter:             transaction.GetQuarter(),
		ProductCode:         product_code,
		CardType:            card_type,
		Function:            transaction.GetFunctionCode(),
		Brand:               transaction.GetBrandCode(),
		CaptureMode:         transaction.GetCaptureModeCode(),
		Installments:        transaction.GetInstallments(),
		SegmentCode:         transaction.GetSegmentCode(),
		InterchangeFee:      transaction.GetInterchangeFeeRate(),
		TransactionAmount:   transaction.GetTransactionAmount(),
		TransactionQuantity: 1,
	}
}

// GetKey generates a unique key for the Intercam struct based on its fields.
func (i *Intercam) GetKey() string {
	return fmt.Sprintf("%d-%d-%d-%s-%s-%d-%d-%d-%d", i.Year, i.Quarter, i.ProductCode, i.CardType, i.Function, i.Brand, i.CaptureMode, i.Installments, i.SegmentCode)
}

// Translate transforms the Intercam struct into a format suitable for database storage, if necessary.
func (i *Intercam) AddTransactions(transactions []*source_domain.Transaction, items map[string]*Intercam, bins map[int64]*source_domain.Bin) {
	for _, t := range transactions {
		interchange := i.GetFromTransaction(t)
		key := interchange.GetKey()
		if existing, exists := items[key]; exists {
			existing.TransactionAmount += interchange.TransactionAmount
			existing.TransactionQuantity += interchange.TransactionQuantity
			delta := interchange.InterchangeFee - existing.InterchangeFee
			existing.InterchangeFee += delta / float64(existing.TransactionQuantity)
			items[key] = existing
		} else {
			items[key] = interchange
		}
	}
}
