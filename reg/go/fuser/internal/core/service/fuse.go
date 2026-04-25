package service

import (
	"time"

	"fuser/internal/core/domain"
	"fuser/internal/core/ports"
)

const (
	managementGapDays = 1
)

// FuseService is the service layer that interacts with the repository to perform business logic
type FuseService struct {
	Repository   ports.Repository
	Logger       ports.Logger
	transactions map[string]*domain.Transaction
}

// NewFuseService creates a new instance of FuseService with the provided repository and logger
func NewFuseService(repository ports.Repository, logger ports.Logger) *FuseService {
	return &FuseService{Repository: repository, Logger: logger, transactions: make(map[string]*domain.Transaction)}
}

// Run executes the main logic of the FuseService (placeholder for actual implementation)
func (s *FuseService) Run(start_date time.Time, end_date time.Time, focus string, leftover bool) error {
	// Log start of the FuseService execution
	s.Logger.IPrintf(0, "Running FuseService...\n")

	// Process main logic based on the provided focus and date range
	if err := s.MainLogic(start_date, end_date, focus); err != nil {
		return err
	}
	// Process leftover transactions if the flag is set
	if err := s.LetfOver(start_date, end_date, leftover); err != nil {
		return err
	}
	// Log end of the FuseService execution
	s.Logger.IPrintf(0, "Finished FuseService.\n")
	return nil
}

// insertTransactionsBatch is a helper method to insert a batch of transactions into the repository using batch processing to optimize performance and reduce memory usage
func (s *FuseService) insertTransactions(focus string, transDate time.Time) error {
	s.Logger.IPrintf(2, "Inserting %s transactions for date %s using batch processing\n", focus, transDate.Format("2006-01-02"))
	// If focus is set to "none", we skip inserting transactions and return early
	if focus == "none" {
		s.Logger.IPrintf(1, "Focus is set to 'none', skipping transaction insertion for date %s.\n", transDate.Format("2006-01-02"))
		return nil
	}
	// prepare transactions for insert by setting the Key2 field based on available data and calculating the status
	transactions := []*domain.Transaction{}
	for _, transaction := range s.transactions {
		transaction.PrepareForInsert()
		transactions = append(transactions, transaction)
	}
	// Insert transactions in batches to optimize database performance and reduce memory usage
	if err := s.Repository.InsertTransactions(transactions); err != nil {
		return err
	}
	// Log the completion of transaction insertion for the current date
	s.Logger.IPrintf(2, "Inserted %d %s transactions for date %s using batch processing\n", len(transactions), focus, transDate.Format("2006-01-02"))
	return nil
}

// mergeExchange is a helper method to merge an Exchange transaction with an existing transaction in the repository, giving priority to non-nil values from the Exchange transaction
func (s *FuseService) mergeExchange(excTransaction *domain.Transaction, repoTransaction *domain.Transaction) {
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

// mergeManagement is a helper method to merge a Management transaction with an existing transaction in the repository, giving priority to non-nil values from the Management transaction
func (s *FuseService) mergeManagement(mgtTransaction *domain.Transaction, repoTransaction *domain.Transaction) {
	// Update fields from management transaction if they are not nil
	repoTransaction.EstablishmentNature = mgtTransaction.EstablishmentNature
	repoTransaction.EstablishmentMCC = mgtTransaction.EstablishmentMCC
	repoTransaction.EstablishmentTerminalCode = mgtTransaction.EstablishmentTerminalCode
	repoTransaction.RevenueMDRValue = mgtTransaction.RevenueMDRValue
	repoTransaction.TransactionSecondaryDate = mgtTransaction.TransactionDate
	repoTransaction.TransactionSecondaryAmount = mgtTransaction.TransactionAmount

	// Calculate status
	if *repoTransaction.StatusID == 0 {
		repoTransaction.StatusCount = 0
		*repoTransaction.StatusID = 2
		*repoTransaction.StatusName = "Pronto"
	}

}

// restartTransactionsMap is a helper method to reset the transactions map to free up memory after processing each date
func (s *FuseService) restartTransactionsMap() {
	s.transactions = make(map[string]*domain.Transaction)
}
