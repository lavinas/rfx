package service

import (
	"time"

	"fuser/internal/core/domain"
)

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
			s.mergeManagement(t1, t0)
			// Cancel the transaction with status 1 (t1) and set reference IDs for both transactions to link them together
			t1.Cancel()
			// Set reference IDs for both transactions to link them together
			t1.ReferenceID = &t0.ID
			t0.ReferenceID = &t1.ID
			// Add the merged transaction to the result slice
			s.transactions[t0.GetKey1()] = t0
			s.transactions[t1.GetKey1()] = t1
			mergedCount++
		}
	}
	return mergedCount
}
