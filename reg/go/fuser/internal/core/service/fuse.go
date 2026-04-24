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
	// Placeholder for actual implementation of the main logic of the FuseService
	s.Logger.IPrintf(0, "Running FuseService...\n")
	// Process main logic based on the provided focus and date range
	if err := s.mainLogic(start_date, end_date, focus); err != nil {
		return err
	}
	// Process leftover transactions if the flag is set
	if err := s.LetfOver(start_date, end_date, leftover); err != nil {
		return err
	}
	// Placeholder for actual implementation of the main logic of the FuseService
	s.Logger.IPrintf(0, "Finished FuseService.\n")
	return nil
}

// main logic of the FuseService would be implemented in the Run method.
func (s *FuseService) mainLogic(start_date time.Time, end_date time.Time, focus string) error {
	// If focus is set to "none", we skip processing transactions and return early
	if focus == "none" {
		s.Logger.IPrintf(1, "Focus is set to 'none', skipping transaction processing.\n")
		return nil
	}
	// Process by date range
	for date := start_date; !date.After(end_date); date = date.AddDate(0, 0, 1) {
		s.Logger.IPrintf(1, "Processing date: %s\n", date.Format("2006-01-02"))
		// Process Exchange transactions if the focus is set to "all" or "exchange"
		if err := s.processExchange(focus, date); err != nil {
			return err
		}
		// Process Management transactions if the focus is set to "all" or "management"
		if err := s.processManagement(focus, date); err != nil {
			return err
		}
		// Insert merged transactions back into the repository if the focus is not set to "none"
		if err := s.insertTransactions(focus, date); err != nil {
			return err
		}
		// reset transactions map for the next date to avoid memory issues and ensure we only keep transactions relevant to the current date in memory
		s.transactions = make(map[string]*domain.Transaction)
		// Log the completion of processing for the current date
		s.Logger.IPrintf(1, "Processed date: %s\n", date.Format("2006-01-02"))
	}
	return nil
}

// getExchange is a helper method to fetch Exchange transactions for a specific date
func (s *FuseService) processExchange(focus string, date time.Time) error {
	if focus != "all" && focus != "exchange" {
		s.Logger.IPrintf(1, "Focus is set to '%s', skipping Exchange transaction processing for date %s.\n", focus, date.Format("2006-01-02"))
		return nil
	}
	s.Logger.IPrintf(2, "Processing Exchange transactions for date %s\n", date.Format("2006-01-02"))
	transactions, err := s.getExchangeTransactions(date)
	if err != nil {
		return err
	}
	if err := s.getTransactionsByKey("management", date, transactions); err != nil {
		return err
	}
	merged := s.mergeTransactions("exchange", date, transactions)
	merged = s.filterDuplicates(merged)

	s.Logger.IPrintf(2, "Finished processing Exchange transactions for date %s\n", date.Format("2006-01-02"))
	return nil
}

// getManagement is a helper method to fetch Management transactions for a specific date
func (s *FuseService) processManagement(focus string, date time.Time) error {
	if focus != "all" && focus != "management" {
		s.Logger.IPrintf(1, "Focus is set to '%s', skipping Management transaction processing for date %s.\n", focus, date.Format("2006-01-02"))
		return nil
	}
	s.Logger.IPrintf(2, "Processing Management transactions for date %s\n", date.Format("2006-01-02"))
	transactions, err := s.getManagementTransactions(date)
	if err != nil {
		return err
	}
	if err := s.getTransactionsByKey("management", date, transactions); err != nil {
		return err
	}
	merged := s.mergeTransactions("management", date, transactions)
	merged = s.filterDuplicates(merged)

	s.Logger.IPrintf(2, "Finished processing Management transactions for date %s\n", date.Format("2006-01-02"))
	return nil
}

// getExchangeTransactions is a helper method to fetch Exchange transactions for a specific date
func (s *FuseService) getExchangeTransactions(date time.Time) ([]*domain.Transaction, error) {
	s.Logger.IPrintf(3, "Reading Exchange transactions for date %s\n", date.Format("2006-01-02"))
	exchanges, err := s.Repository.GetExchangeTransactions(date)
	if err != nil {
		s.Logger.IPrintf(3, "Error reading exchange transactions for date %s: %v\n", date.Format("2006-01-02"), err)
		return nil, err
	}
	transactions := []*domain.Transaction{}
	for _, exchange := range exchanges {
		transactions = append(transactions, exchange.Translate())
	}
	s.Logger.IPrintf(3, "Read %d Exchange transactions for date %s\n", len(transactions), date.Format("2006-01-02"))
	return transactions, nil
}

// getManagementTransactions is a helper method to fetch Management transactions for a specific date
func (s *FuseService) getManagementTransactions(date time.Time) ([]*domain.Transaction, error) {
	s.Logger.IPrintf(3, "Reading Management transactions for date %s\n", date.Format("2006-01-02"))
	managements, err := s.Repository.GetManagementTransactions(date)
	if err != nil {
		s.Logger.IPrintf(3, "Error reading management transactions for date %s: %v\n", date.Format("2006-01-02"), err)
		return nil, err
	}
	transactions := []*domain.Transaction{}
	for _, management := range managements {
		transactions = append(transactions, management.Translate())
	}
	s.Logger.IPrintf(3, "Read %d Management transactions for date %s\n", len(transactions), date.Format("2006-01-02"))
	return transactions, nil
}

// getTransactionsByKey is a helper method to fetch transactions by their keys
func (s *FuseService) getTransactionsByKey(transType string, transDate time.Time, transactions []*domain.Transaction) error {
	s.Logger.IPrintf(3, "Fetching %s transactions by keys for date %s\n", transType, transDate.Format("2006-01-02"))
	keys := []string{}
	// Fetch transactions in batches of const loadRate to optimize database performance and avoid memory issues
	for _, transaction := range transactions {
		if _, exists := s.transactions[transaction.Key1]; !exists {
			keys = append(keys, transaction.Key1)
		}
	}
	// Fetch any remaining transactions from the repository by their keys that were not fetched in the previous loop
	tr, err := s.Repository.GetTransactionsByKey(keys)
	if err != nil {
		return err
	}
	// Store fetched transactions in the service's transactions map for later use in merging and inserting transactions
	for _, transaction := range tr {
		s.transactions[transaction.Key1] = transaction
	}
	s.Logger.IPrintf(3, "Fetched %d %s transactions and %d found in cache for date %s\n", len(tr), transType, len(transactions)-len(tr), transDate.Format("2006-01-02"))
	return nil
}

// mergeTransactions is a helper method to merge Exchange transactions with existing transactions in the repository
func (s *FuseService) mergeTransactions(transType string, transDate time.Time, localTransactions []*domain.Transaction) []*domain.Transaction {
	merged := []*domain.Transaction{}
	repoMap := make(map[string]*domain.Transaction)
	for _, repoTrans := range s.transactions {
		repoMap[repoTrans.Key1] = repoTrans
	}
	for _, localTrans := range localTransactions {
		if repoTrans, exists := repoMap[localTrans.Key1]; exists {
			if transType == "exchange" {
				s.MergeExchange(localTrans, repoTrans)
			} else {
				s.MergeManagement(localTrans, repoTrans)
			}
			merged = append(merged, repoTrans)
		} else {
			merged = append(merged, localTrans)
		}
	}
	s.Logger.IPrintf(3, "Merged %s transactions for date %s (local: %d, repository: %d, merged: %d)\n", transType, transDate.Format("2006-01-02"), len(localTransactions), len(s.transactions), len(merged))
	return merged
}

// insertTransactionsBatch is a helper method to insert a batch of transactions into the repository using batch processing to optimize performance and reduce memory usage
func (s *FuseService) insertTransactions(focus string, transDate time.Time) error {
	// If focus is set to "none", we skip inserting transactions and return early
	if focus == "none" {
		s.Logger.IPrintf(1, "Focus is set to 'none', skipping transaction insertion for date %s.\n", transDate.Format("2006-01-02"))
		return nil
	}
	// prepare transactions for insert by setting the Key2 field based on available data and calculating the status
	s.Logger.IPrintf(2, "Preparing %d %s transactions for date %s for batch insertion\n", len(s.transactions), focus, transDate.Format("2006-01-02"))
	transactions := []*domain.Transaction{}
	for _, transaction := range s.transactions {
		transaction.PrepareForInsert()
		transactions = append(transactions, transaction)
	}
	s.Logger.IPrintf(2, "Prepared %d %s transactions for date %s for batch insertion\n", len(s.transactions), focus, transDate.Format("2006-01-02"))
	// Insert transactions in batches to optimize database performance and reduce memory usage
	s.Logger.IPrintf(2, "Inserting %d %s transactions for date %s using batch processing\n", len(transactions), focus, transDate.Format("2006-01-02"))
	if err := s.Repository.InsertTransactions(transactions); err != nil {
		return err
	}
	// Log the completion of transaction insertion for the current date
	s.Logger.IPrintf(2, "Inserted %d %s transactions for date %s using batch processing\n", len(transactions), focus, transDate.Format("2006-01-02"))
	return nil
}

// filterDuplicates is a helper method to filter out duplicate transactions based on their keys
func (s *FuseService) filterDuplicates(transactions []*domain.Transaction) []*domain.Transaction {
	// Log the number of transactions before filtering duplicates
	s.Logger.IPrintf(3, "Filtering duplicates from %d transactions\n", len(transactions))
	// Use a map to track unique transactions by their keys
	unique := make(map[string]*domain.Transaction)
	for _, transaction := range transactions {
		if _, exists := unique[transaction.Key1]; exists {
			s.Logger.IPrintf(4, "Duplicate transaction found with key: %s\n", transaction.Key1)
		}
		unique[transaction.Key1] = transaction
	}
	// Convert the map of unique transactions back to a slice
	result := []*domain.Transaction{}
	for _, transaction := range unique {
		result = append(result, transaction)
	}
	// Log the number of transactions after filtering duplicates
	s.Logger.IPrintf(3, "Filtered duplicates, resulting in %d unique transactions\n", len(result))
	return result
}

// LetfOver is a helper that treats transactions that were not merged (i.e., they exist in the repository but not in the local data) - placeholder for actual implementation
func (s *FuseService) LetfOver(start_date time.Time, end_date time.Time, leftover bool) error {
	// if the leftover flag is not set, we skip processing leftover transactions and return early
	if !leftover {
		s.Logger.IPrintf(1, "Leftover processing is disabled, skipping.\n")
		return nil
	}
	// Log the start of leftover processing
	s.Logger.IPrintf(1, "Processing leftover transactions...\n")
	// Get leftover transactions from the repository for the given date range
	transactions_0, transactions_1, err := s.getLeftover(start_date, end_date)
	if err != nil {
		return err
	}
	// Merge transactions_0 and transactions_1, giving priority to transactions_0 (status 0) over transactions_1 (status 1)
	s.mergeLeftover(transactions_0, transactions_1)
	// Insert merged transactions back into the repository
	err = s.insertTransactions("leftover", time.Now())
	if err != nil {
		return err
	}
	// Log the completion of leftover processing
	s.Logger.IPrintf(1, "Finished processing leftover transactions.\n")
	return nil
}

// getLeftover is a helper method to fetch transactions that exist in the repository but were not merged (i.e., they do not exist in the local data) - placeholder for actual implementation
func (s *FuseService) getLeftover(start, end time.Time) ([]*domain.Transaction, []*domain.Transaction, error) {
	// Placeholder for actual implementation of fetching leftover transactions
	s.Logger.IPrintf(2, "Fetching leftover transactions for date %s\n", start.Format("2006-01-02"))
	// Get transaction with status 0 (not processed) from the repository for the given date range
	transactions_0, err := s.Repository.GetTransactionsByDateRangeAndStatus(start, end, 0)
	if err != nil {
		return nil, nil, err
	}
	// Get transaction with status 1 (processed) from the repository for the given extended date range (to account for transactions that might have been processed but not merged)
	extended_start := start.AddDate(0, 0, -3)
	extended_end := end.AddDate(0, 0, 3)
	transactions_1, err := s.Repository.GetTransactionsByDateRangeAndStatus(extended_start, extended_end, 1)
	if err != nil {
		return nil, nil, err
	}
	// Log the number of leftover transactions fetched for both status 0 and status 1
	s.Logger.IPrintf(2, "Fetched %d leftover transactions with status 0 and %d with status 1 for date range %s to %s\n", len(transactions_0), len(transactions_1), start.Format("2006-01-02"), end.Format("2006-01-02"))
	return transactions_0, transactions_1, nil
}

// mergeLeftover is a helper method to merge leftover transactions, giving priority to transactions with status 0 over those with status 1 - placeholder for actual implementation
func (s *FuseService) mergeLeftover(transactions_0, transactions_1 []*domain.Transaction) {
	// Placeholder for actual implementation of merging leftover transactions
	s.Logger.IPrintf(2, "Merging leftover transactions (status 0: %d, status 1: %d)\n", len(transactions_0), len(transactions_1))
	t0_map := s.getLeftoverMap(transactions_0)
	t1_map := s.getLeftoverMap(transactions_1)
	mergedCount := s.mergeLeftoverMaps(t0_map, t1_map)
	// Log the number of merged leftover transactions
	s.Logger.IPrintf(2, "Merged %d leftover transactions (status 0: %d, status 1: %d)\n", mergedCount, len(transactions_0), len(transactions_1))
}

// getLeftoverMap creates a map from a slice of transactions based on their keys
func (s *FuseService) getLeftoverMap(transactions []*domain.Transaction) map[string]*domain.Transaction {
	tmap := make(map[string]*domain.Transaction)
	for _, transaction := range transactions {
		if transaction.Key2 == nil {
			continue
		}
		// Eliminate duplicates by checking if the transaction key already exists in the map, and if it does, we skip adding it again to the map
		if _, exists := tmap[*transaction.Key2]; exists {
			delete(tmap, *transaction.Key2)
			s.Logger.IPrintf(2, "Duplicate transaction found with key: %s, removing from merged map\n", *transaction.Key2)
			continue
		}
		tmap[*transaction.Key2] = transaction
	}
	return tmap
}

// mergeLeftoverMaps merges two maps of transactions based on their keys, giving priority to transactions in the first map over those in the second map
func (s *FuseService) mergeLeftoverMaps(t0_map, t1_map map[string]*domain.Transaction) int {
	s.transactions = make(map[string]*domain.Transaction)
	// Iterate over the first map (status 0) and check if there is a corresponding transaction in the second map (status 1) with the same key
	mergedCount := 0
	for key, t0 := range t0_map {
		// if there is a corresponding transaction in the second map with the same key
		if t1, exists := t1_map[key]; exists {
			if t0.TransactionDate == nil || t1.TransactionDate == nil {
				continue
			}
			// Allow a difference of up to 3 days between the transaction dates to account for potential delays in processing and merging transactions
			if t0.TransactionDate.After(t1.TransactionDate.AddDate(0, 0, 3)) || t0.TransactionDate.Before(t1.TransactionDate.AddDate(0, 0, -3)) {
				continue
			}
			// Merge transactions by giving priority to non-nil values from the transaction with status 0 (t0) over the transaction with status 1 (t1)
			s.MergeManagement(t1, t0)
			// Cancel the transaction with status 1 (t1) and set reference IDs for both transactions to link them together
			t1.Cancel()
			// Set reference IDs for both transactions to link them together
			t1.ReferenceID = &t0.ID
			t0.ReferenceID = &t1.ID
			// Add the merged transaction to the result slice
			s.transactions[t0.GetKey1()] = t0
			s.transactions[t1.GetKey1()+"_cancelled"] = t1
			mergedCount++
		}
	}
	return mergedCount
}

// MergeExchange is a helper method to merge an Exchange transaction with an existing transaction in the repository, giving priority to non-nil values from the Exchange transaction
func (s *FuseService) MergeExchange(excTransaction *domain.Transaction, repoTransaction *domain.Transaction) {
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

// MergeManagement is a helper method to merge a Management transaction with an existing transaction in the repository, giving priority to non-nil values from the Management transaction
func (s *FuseService) MergeManagement(mgtTransaction *domain.Transaction, repoTransaction *domain.Transaction) {
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
