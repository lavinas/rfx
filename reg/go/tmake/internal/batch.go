package internal  

// Batch struct to hold a batch of transactions
type Batch struct {
	db *DB
	Transactions []*Transaction
}

// NewBatch creates a new Batch instance
func NewBatch() *Batch {
	return &Batch{
		Transactions: make([]*Transaction, 0),
	}
}

// AddTransaction adds a transaction to the batch
func (b *Batch) AddTransaction(transaction *Transaction) {
	b.Transactions = append(b.Transactions, transaction)
}

// Clear removes all transactions from the batch
func (b *Batch) Clear() {
	b.Transactions = make([]*Transaction, 0)
}

// Insert inserts all transactions in the batch into the database and returns the number of inserted transactions
func (b *Batch) Insert(db *DB) error {
	if len(b.Transactions) == 0 {
		return nil
	}
	sql := `insert into transaction.transaction(id, created_at, updated_at, key1, establishment_code, establishment_nature, establishment_mcc, establishment_terminal_code,
	        bin, authorization_code, transaction_nsu, transaction_date, transaction_amount, transaction_installments, transaction_brand,
			transaction_product, transaction_capture, revenue_mdr_value, cost_interchange_value, high_source_priority, status_id, status_name, status_count,
			period_date,period_closing_id,transac_id) values `
	for _, transaction := range b.Transactions {
		sql += transaction.GetInsertRow() + ","		
	}
	sql = sql[:len(sql)-1] // Remove the trailing comma
	_, err := db.Exec(sql)
	b.Clear()
	return err
}

// GetLastID retrieves the last transaction ID from the database
func (b *Batch) GetLastID(db *DB) int64 {
	var lastID int64
	err := db.QueryRow("SELECT COALESCE(MAX(id), 0) FROM transaction.transaction").Scan(&lastID)
	if err != nil {
		panic(err)
	}
	return lastID
}
