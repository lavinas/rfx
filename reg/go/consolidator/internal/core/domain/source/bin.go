package source_domain

const (
	DefaultProductCode = 31
	DefaultCardType    = "P"
)

// Bin represents the Bank Identification Number (BIN) associated with a transaction.
type Bin struct {
	Bin      int64   `gorm:"column:bin"`
	Product  *int    `gorm:"column:produto_final"`
	CardType *string `gorm:"column:modalidade_final"`
}

// TableName specifies the table name for Bin struct
func (Bin) TableName() string {
	return "bins.bins"
}

// GetProductCode returns the product code based on the BIN information.
func (b *Bin) GetProductCode() int {
	if b.Product != nil {
		return *b.Product
	}
	return DefaultProductCode
}

// GetCardType returns the card type based on the BIN information.
func (b *Bin) GetCardType() string {
	if b.CardType != nil {
		return *b.CardType
	}
	return DefaultCardType
}
