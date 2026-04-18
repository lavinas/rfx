package domain

// MergeManagement merges two  transactions into one, prioritizing Intercam values over Management values when both are available
func MergeManagement(interTransaction *Transaction, repoTransaction *Transaction) {
	// Update fields from Intercam transaction if they are not nil
	if repoTransaction.EstablishmentCode == nil {
		repoTransaction.EstablishmentCode = interTransaction.EstablishmentCode
	}
	if repoTransaction.EstablishmentNature == nil {
		repoTransaction.EstablishmentNature = interTransaction.EstablishmentNature
	}
	if repoTransaction.EstablishmentMCC == nil {
		repoTransaction.EstablishmentMCC = interTransaction.EstablishmentMCC
	}
	if repoTransaction.EstablishmentTerminalCode == nil {
		repoTransaction.EstablishmentTerminalCode = interTransaction.EstablishmentTerminalCode
	}
	if repoTransaction.BIN == nil {
		repoTransaction.BIN = interTransaction.BIN
	}
	if repoTransaction.AuthorizationCode == nil {
		repoTransaction.AuthorizationCode = interTransaction.AuthorizationCode
	}
	if repoTransaction.TransactionNSU == nil {
		repoTransaction.TransactionNSU = interTransaction.TransactionNSU
	}
	if repoTransaction.TransactionDate == nil {
		repoTransaction.TransactionDate = interTransaction.TransactionDate
	}
	if repoTransaction.TransactionAmount == nil {
		repoTransaction.TransactionAmount = interTransaction.TransactionAmount
	}
	if repoTransaction.TransactionInstallments == nil {
		repoTransaction.TransactionInstallments = interTransaction.TransactionInstallments
	}
	if repoTransaction.TransactionBrand == nil {
		repoTransaction.TransactionBrand = interTransaction.TransactionBrand
	}
	if repoTransaction.TransactionProduct == nil {
		repoTransaction.TransactionProduct = interTransaction.TransactionProduct
	}
	if repoTransaction.TransactionCapture == nil {
		repoTransaction.TransactionCapture = interTransaction.TransactionCapture
	}
	if repoTransaction.CostInterchangeValue == nil {
		repoTransaction.CostInterchangeValue = interTransaction.CostInterchangeValue
	}
	if repoTransaction.HighSourcePriority == nil {
		repoTransaction.HighSourcePriority = interTransaction.HighSourcePriority
	}
	if repoTransaction.PeriodDate == nil {
		repoTransaction.PeriodDate = interTransaction.PeriodDate
	}
	if repoTransaction.PeriodClosingID == nil {
		repoTransaction.PeriodClosingID = interTransaction.PeriodClosingID
	}
	if repoTransaction.TransacID == nil {
		repoTransaction.TransacID = interTransaction.TransacID
	}
	if repoTransaction.RevenueMDRValue == nil {
		repoTransaction.RevenueMDRValue = interTransaction.RevenueMDRValue
	}

	// Update secondary date
	repoTransaction.TransactionSecondaryDate = interTransaction.TransactionDate

	// Calculate status
	if *repoTransaction.StatusID == 0 {
		repoTransaction.StatusCount = 0
		*repoTransaction.StatusID = 2
		*repoTransaction.StatusName = "Pronto"
	}

}
