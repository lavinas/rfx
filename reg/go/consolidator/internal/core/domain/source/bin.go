package source_domain

// Bin represents the Bank Identification Number (BIN) associated with a transaction.
type Bin struct {
	Bin      int64  `gorm:"column:bin"`
	Product  string `gorm:"column:produto_final"`
	CardType string `gorm:"column:modalidade_final"`
}

// TableName specifies the table name for Bin struct
func (Bin) TableName() string {
	return "bins.bins"
}

// BinMap represents a collection of Bin entries, which can be used for mapping BIN numbers to their corresponding product and card type information.
type BinMap struct {
	BinMap []Bin
}
