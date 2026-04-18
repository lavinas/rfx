package domain

// MergeTransactions merges two Intercam transactions into one, prioritizing non-nil values from the second transaction
func MergeIntercam(interTransaction *Transaction, repoTransaction *Transaction) {
	// Update fields from Intercam transaction if they are not nil
	repoTransaction.BIN = interTransaction.BIN
	repoTransaction.AuthorizationCode = interTransaction.AuthorizationCode
	repoTransaction.TransactionNSU = interTransaction.TransactionNSU
	repoTransaction.TransactionDate = interTransaction.TransactionDate
	repoTransaction.TransactionAmount = interTransaction.TransactionAmount
	repoTransaction.TransactionInstallments = interTransaction.TransactionInstallments
	repoTransaction.TransactionBrand = interTransaction.TransactionBrand
	repoTransaction.TransactionProduct = interTransaction.TransactionProduct
	repoTransaction.TransactionCapture = interTransaction.TransactionCapture
	repoTransaction.CostInterchangeValue = interTransaction.CostInterchangeValue
	repoTransaction.HighSourcePriority = interTransaction.HighSourcePriority
	repoTransaction.PeriodDate = interTransaction.PeriodDate
	repoTransaction.PeriodClosingID = interTransaction.PeriodClosingID
	repoTransaction.TransacID = interTransaction.TransacID

	// Calculate status
	if *repoTransaction.StatusID == 1 {
		repoTransaction.StatusCount = 0
		*repoTransaction.StatusID = 2
		*repoTransaction.StatusName = "Pronto"
	}

}
