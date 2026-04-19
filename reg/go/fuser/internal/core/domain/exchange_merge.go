package domain

// MergeTransactions merges two transactions into one, prioritizing non-nil values from the second transaction
func MergeExchange(excTransaction *Transaction, repoTransaction *Transaction) {
	// Update fields from exchange transaction if they are not nil
	repoTransaction.BIN = excTransaction.BIN
	repoTransaction.AuthorizationCode = excTransaction.AuthorizationCode
	repoTransaction.TransactionNSU = excTransaction.TransactionNSU
	repoTransaction.TransactionDate = excTransaction.TransactionDate
	repoTransaction.TransactionAmount = excTransaction.TransactionAmount
	repoTransaction.TransactionInstallments = excTransaction.TransactionInstallments
	repoTransaction.TransactionBrand = excTransaction.TransactionBrand
	repoTransaction.TransactionProduct = excTransaction.TransactionProduct
	repoTransaction.TransactionCapture = excTransaction.TransactionCapture
	repoTransaction.CostInterchangeValue = excTransaction.CostInterchangeValue
	repoTransaction.HighSourcePriority = excTransaction.HighSourcePriority
	repoTransaction.PeriodDate = excTransaction.PeriodDate
	repoTransaction.PeriodClosingID = excTransaction.PeriodClosingID
	repoTransaction.TransacID = excTransaction.TransacID

	// Calculate status
	if *repoTransaction.StatusID == 1 {
		repoTransaction.StatusCount = 0
		*repoTransaction.StatusID = 2
		*repoTransaction.StatusName = "Pronto"
	}

}
