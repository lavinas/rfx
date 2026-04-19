package domain

// MergeManagement merges two  transactions into one, prioritizing Intercam values over Management values when both are available
func MergeManagement(manegTransaction *Transaction, repoTransaction *Transaction) {
	// Update fields from Intercam transaction if they are not nil
	if repoTransaction.EstablishmentCode == nil {
		repoTransaction.EstablishmentCode = manegTransaction.EstablishmentCode
	}
	if repoTransaction.EstablishmentNature == nil {
		repoTransaction.EstablishmentNature = manegTransaction.EstablishmentNature
	}
	if repoTransaction.EstablishmentMCC == nil {
		repoTransaction.EstablishmentMCC = manegTransaction.EstablishmentMCC
	}
	if repoTransaction.EstablishmentTerminalCode == nil {
		repoTransaction.EstablishmentTerminalCode = manegTransaction.EstablishmentTerminalCode
	}
	if repoTransaction.BIN == nil {
		repoTransaction.BIN = manegTransaction.BIN
	}
	if repoTransaction.AuthorizationCode == nil {
		repoTransaction.AuthorizationCode = manegTransaction.AuthorizationCode
	}
	if repoTransaction.TransactionNSU == nil {
		repoTransaction.TransactionNSU = manegTransaction.TransactionNSU
	}
	if repoTransaction.TransactionDate == nil {
		repoTransaction.TransactionDate = manegTransaction.TransactionDate
	}
	if repoTransaction.TransactionAmount == nil {
		repoTransaction.TransactionAmount = manegTransaction.TransactionAmount
	}
	if repoTransaction.TransactionInstallments == nil {
		repoTransaction.TransactionInstallments = manegTransaction.TransactionInstallments
	}
	if repoTransaction.TransactionBrand == nil {
		repoTransaction.TransactionBrand = manegTransaction.TransactionBrand
	}
	if repoTransaction.TransactionProduct == nil {
		repoTransaction.TransactionProduct = manegTransaction.TransactionProduct
	}
	if repoTransaction.TransactionCapture == nil {
		repoTransaction.TransactionCapture = manegTransaction.TransactionCapture
	}
	if repoTransaction.CostInterchangeValue == nil {
		repoTransaction.CostInterchangeValue = manegTransaction.CostInterchangeValue
	}
	if repoTransaction.HighSourcePriority == nil {
		repoTransaction.HighSourcePriority = manegTransaction.HighSourcePriority
	}
	if repoTransaction.PeriodDate == nil {
		repoTransaction.PeriodDate = manegTransaction.PeriodDate
	}
	if repoTransaction.PeriodClosingID == nil {
		repoTransaction.PeriodClosingID = manegTransaction.PeriodClosingID
	}
	if repoTransaction.TransacID == nil {
		repoTransaction.TransacID = manegTransaction.TransacID
	}
	if repoTransaction.RevenueMDRValue == nil {
		repoTransaction.RevenueMDRValue = manegTransaction.RevenueMDRValue
	}

	// Update secondary date
	repoTransaction.TransactionSecondaryDate = manegTransaction.TransactionDate
	repoTransaction.TransactionSecondaryAmount = manegTransaction.TransactionAmount

	// Calculate status
	if *repoTransaction.StatusID == 0 {
		repoTransaction.StatusCount = 0
		*repoTransaction.StatusID = 2
		*repoTransaction.StatusName = "Pronto"
	}

}
