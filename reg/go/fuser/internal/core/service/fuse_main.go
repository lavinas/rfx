package service

import (
	"time"

	"fuser/internal/core/domain"
)

// main logic of the FuseService would be implemented in the Run method.
func (s *FuseService) MainLogic(start_date time.Time, end_date time.Time, focus string) error {
	// If focus is set to "none", we skip processing transactions and return early
	if focus == "none" {
		s.Logger.IPrintf(1, "Focus is set to 'none', skipping transaction processing.\n")
		return nil
	}
	// Process by date range
	for date := start_date; !date.After(end_date); date = date.AddDate(0, 0, 1) {
		if err := s.mainLogicDay(date, focus); err != nil {
			return err
		}
	}
	return nil
}

// mainLogicDay is a helper method to process transactions for a specific date, handling both Exchange and Management transactions based on the provided focus
func (s *FuseService) mainLogicDay(date time.Time, focus string) error {
	// Log the start of processing for the current date
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
	// restart transactions map to free up memory after processing each date
	s.restartTransactionsMap()

	// Log the completion of processing for the current date
	s.Logger.IPrintf(1, "Processed date: %s\n", date.Format("2006-01-02"))

	return nil
}

// getExchange is a helper method to fetch Exchange transactions for a specific date
func (s *FuseService) processExchange(focus string, date time.Time) error {
	// If focus is set to "all" or "exchange", we proceed with processing Exchange transactions; otherwise, we skip and return early
	if focus != "all" && focus != "exchange" {
		s.Logger.IPrintf(1, "Focus is set to '%s', skipping Exchange transaction processing for date %s.\n", focus, date.Format("2006-01-02"))
		return nil
	}
	// Get Exchange transactions for the current date from the repository and log the number of transactions read for debugging and monitoring purposes
	s.Logger.IPrintf(2, "Processing Exchange transactions for date %s\n", date.Format("2006-01-02"))
	transactions, err := s.getExchangeTransactions(date)
	if err != nil {
		return err
	}
	// Fetch any additional transactions from the repository by their keys that were not fetched in the previous step
	if err := s.getTransactionsByKey("exchange", date, transactions); err != nil {
		return err
	}
	// Merge Exchange transactions with existing transactions in the repository, giving priority to non-nil values from the Exchange transactions, and log the number of transactions merged for debugging and monitoring purposes
	s.mergeTransactions("exchange", date, transactions)
	// Filter out duplicate transactions based on their keys and log the number of transactions after filtering duplicates for debugging and monitoring purposes
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
	s.mergeTransactions("management", date, transactions)

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
	s.Logger.IPrintf(3, "Fetched %d %s transactions: %d found in cache and %d fetched from repository for date %s\n", len(transactions), transType, len(s.transactions)-len(tr), len(tr), transDate.Format("2006-01-02"))
	return nil
}

// mergeTransactions is a helper method to merge Exchange transactions with existing transactions in the repository
func (s *FuseService) mergeTransactions(transType string, transDate time.Time, localTransactions []*domain.Transaction) {
	merged := []*domain.Transaction{}
	repoMap := make(map[string]*domain.Transaction)
	for _, repoTrans := range s.transactions {
		repoMap[repoTrans.Key1] = repoTrans
	}
	for _, localTrans := range localTransactions {
		if repoTrans, exists := repoMap[localTrans.Key1]; exists {
			if transType == "exchange" {
				s.mergeExchange(localTrans, repoTrans)
			} else {
				s.mergeManagement(localTrans, repoTrans)
			}
			merged = append(merged, repoTrans)
		} else {
			merged = append(merged, localTrans)
		}
	}
	merged = s.filterDuplicates(merged)
	for _, transaction := range merged {
		s.transactions[transaction.Key1] = transaction
	}
	s.Logger.IPrintf(3, "Merged %s transactions for date %s (local: %d, repository: %d, merged: %d)\n", transType, transDate.Format("2006-01-02"), len(localTransactions), len(s.transactions), len(merged))
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
	s.Logger.IPrintf(3, "Filtered duplicates, resulting in %d duplicated and %d unique transactions\n", len(transactions)-len(result), len(result))
	return result
}
