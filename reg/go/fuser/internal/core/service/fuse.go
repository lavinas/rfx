package service

import (
	"time"

	"fuser/internal/core/domain"
	"fuser/internal/core/ports"
)

// FuseService is the service layer that interacts with the repository to perform business logic
type FuseService struct {
	Repository ports.Repository
	Logger     ports.Logger
}

// NewFuseService creates a new instance of FuseService with the provided repository and logger
func NewFuseService(repository ports.Repository, logger ports.Logger) *FuseService {
	return &FuseService{Repository: repository, Logger: logger}
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

// main logic of the FuseService would be implemented in the Run method, which would call helper methods to process Intercam and Management transactions, as well as handle leftover transactions based on the provided flags and focus.
func (s *FuseService) mainLogic(start_date time.Time, end_date time.Time, focus string) error {
	// If focus is set to "none", we skip processing transactions and return early
	if focus == "none" {
		s.Logger.IPrintf(1, "Focus is set to 'none', skipping transaction processing.\n")
		return nil
	}
	// Process by date range
	for date := start_date; !date.After(end_date); date = date.AddDate(0, 0, 1) {
		s.Logger.IPrintf(1, "Processing date: %s\n", date.Format("2006-01-02"))
		// Process Intercam transactions if the focus is set to "all" or "intercam"
		if focus == "all" || focus == "intercam" {
			err := s.processIntercam(date)
			if err != nil {
				return err
			}
		}
		// Process Management transactions if the focus is set to "all" or "management"
		if focus == "all" || focus == "management" {
			err := s.processManagement(date)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// getIntercam is a helper method to fetch Intercam transactions for a specific date
func (s *FuseService) processIntercam(date time.Time) error {
	s.Logger.IPrintf(2, "Processing Intercam transactions for date %s\n", date.Format("2006-01-02"))
	transactions, err := s.getIntercamTransactions(date)
	if err != nil {
		return err
	}
	byKey, err := s.getTransactionsByKey("intercam", date, transactions)
	if err != nil {
		return err
	}
	merged := s.mergeTransactions("intercam", date, transactions, byKey)
	merged = s.filterDuplicates(merged)

	err = s.insertTransactions("intercam", date, merged)
	if err != nil {
		return err
	}
	s.Logger.IPrintf(2, "Finished processing Intercam transactions for date %s\n", date.Format("2006-01-02"))
	return nil
}

// getManagement is a helper method to fetch Management transactions for a specific date
func (s *FuseService) processManagement(date time.Time) error {
	s.Logger.IPrintf(2, "Processing Management transactions for date %s\n", date.Format("2006-01-02"))
	transactions, err := s.getManagementTransactions(date)
	if err != nil {
		return err
	}
	byKey, err := s.getTransactionsByKey("management", date, transactions)
	if err != nil {
		return err
	}
	merged := s.mergeTransactions("management", date, transactions, byKey)
	merged = s.filterDuplicates(merged)

	err = s.insertTransactions("management", date, merged)
	if err != nil {
		return err
	}
	s.Logger.IPrintf(2, "Finished processing Management transactions for date %s\n", date.Format("2006-01-02"))
	return nil
}

// getIntercamTransactions is a helper method to fetch Intercam transactions for a specific date
func (s *FuseService) getIntercamTransactions(date time.Time) ([]*domain.Transaction, error) {
	intercams, err := s.Repository.GetIntercamTransactions(date)
	if err != nil {
		s.Logger.IPrintf(3, "Error fetching intercamtransactions for date %s: %v\n", date.Format("2006-01-02"), err)
		return nil, err
	}
	transactions := []*domain.Transaction{}
	for _, intercam := range intercams {
		transactions = append(transactions, intercam.Translate())
	}
	return transactions, nil
}

// getManagementTransactions is a helper method to fetch Management transactions for a specific date
func (s *FuseService) getManagementTransactions(date time.Time) ([]*domain.Transaction, error) {
	managements, err := s.Repository.GetManagementTransactions(date)
	if err != nil {
		s.Logger.IPrintf(3, "Error fetching management transactions for date %s: %v\n", date.Format("2006-01-02"), err)
		return nil, err
	}
	transactions := []*domain.Transaction{}
	for _, management := range managements {
		transactions = append(transactions, management.Translate())
	}
	return transactions, nil
}

// getTransactionsByKey is a helper method to fetch transactions by their keys
func (s *FuseService) getTransactionsByKey(transType string, transDate time.Time, transactions []*domain.Transaction) ([]*domain.Transaction, error) {
	repTransactions := []*domain.Transaction{}
	keys := []string{}
	count := 0
	total := len(transactions)
	for _, transaction := range transactions {
		keys = append(keys, transaction.Key1)
		count++
		if count%2000 == 0 {
			repTrans, err := s.Repository.GetTransactionsByKey(keys)
			if err != nil {
				return nil, err
			}
			s.Logger.IPrintf(2, "Fetched %d %s transactions by keys for %s date %s (%d/%d)\n", len(repTrans), transType, transType, transDate.Format("2006-01-02"), count, total)
			repTransactions = append(repTransactions, repTrans...)
			keys = []string{}
		}
	}
	if len(keys) > 0 {
		repTrans, err := s.Repository.GetTransactionsByKey(keys)
		if err != nil {
			return nil, err
		}
		s.Logger.IPrintf(2, "Fetched %d %s transactions by keys for %s date %s (%d/%d)\n", len(repTrans), transType, transType, transDate.Format("2006-01-02"), count, total)
		repTransactions = append(repTransactions, repTrans...)
	}
	return repTransactions, nil
}

// mergeTransactions is a helper method to merge Intercam transactions with existing transactions in the repository
func (s *FuseService) mergeTransactions(transType string, transDate time.Time, localTransactions []*domain.Transaction, repositoryTransactions []*domain.Transaction) []*domain.Transaction {
	merged := []*domain.Transaction{}
	repoMap := make(map[string]*domain.Transaction)
	for _, repoTrans := range repositoryTransactions {
		repoMap[repoTrans.Key1] = repoTrans
	}
	for _, localTrans := range localTransactions {
		if repoTrans, exists := repoMap[localTrans.Key1]; exists {
			switch transType {
			case "intercam":
				domain.MergeIntercam(localTrans, repoTrans)
			case "management":
				domain.MergeManagement(localTrans, repoTrans)
			default:
				s.Logger.IPrintf(2, "Unknown transaction type: %s\n", transType)
			}
			merged = append(merged, repoTrans)
		} else {
			merged = append(merged, localTrans)
		}
	}
	s.Logger.IPrintf(2, "Merged %s transactions for date %s (local: %d, repository: %d, merged: %d)\n", transType, transDate.Format("2006-01-02"), len(localTransactions), len(repositoryTransactions), len(merged))
	return merged
}

// insertTransactions is a helper method to insert a batch of transactions into the repository
func (s *FuseService) insertTransactions(transType string, transDate time.Time, transactions []*domain.Transaction) error {
	count := 0
	total := len(transactions)
	lot := []*domain.Transaction{}
	// Insert transactions in batches of 2000 to optimize database performance and avoid memory issues
	for _, transaction := range transactions {
		transaction.PrepareForInsert()
		lot = append(lot, transaction)
		count++
		// When the batch size reaches 2000, we insert the batch into the repository and reset the batch
		if count%2000 == 0 {
			if err := s.Repository.InsertTransactions(lot); err != nil {
				s.Logger.IPrintf(2, "Error inserting %s transactions for date %s: %v\n", transType, transDate.Format("2006-01-02"), err)
				return err
			}
			s.Logger.IPrintf(2, "Inserted %s transactions for date %s (%d/%d)\n", transType, transDate.Format("2006-01-02"), count, total)
			lot = []*domain.Transaction{}
		}
	}
	// Insert any remaining transactions in the batch that were not inserted in the previous loop
	if len(lot) > 0 {
		if err := s.Repository.InsertTransactions(lot); err != nil {
			s.Logger.IPrintf(2, "Error inserting %s transactions for date %s: %v\n", transType, transDate.Format("2006-01-02"), err)
			return err
		}
		// Log the number of transactions inserted for the last batch
		s.Logger.IPrintf(2, "Inserted %s transactions for date %s (%d/%d)\n", transType, transDate.Format("2006-01-02"), count, total)
	}
	return nil
}

// filterDuplicates is a helper method to filter out duplicate transactions based on their keys
func (s *FuseService) filterDuplicates(transactions []*domain.Transaction) []*domain.Transaction {
	// Log the number of transactions before filtering duplicates
	s.Logger.IPrintf(2, "Filtering duplicates from %d transactions\n", len(transactions))
	// Use a map to track unique transactions by their keys
	unique := make(map[string]*domain.Transaction)
	for _, transaction := range transactions {
		if _, exists := unique[transaction.Key1]; exists {
			s.Logger.IPrintf(2, "Duplicate transaction found with key: %s\n", transaction.Key1)
		}
		unique[transaction.Key1] = transaction
	}
	// Convert the map of unique transactions back to a slice
	result := []*domain.Transaction{}
	for _, transaction := range unique {
		result = append(result, transaction)
	}
	// Log the number of transactions after filtering duplicates
	s.Logger.IPrintf(2, "Filtered duplicates, resulting in %d unique transactions\n", len(result))
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
	merged := s.mergeLeftover(transactions_0, transactions_1)
	// Insert merged transactions back into the repository
	err = s.insertTransactions("leftover", time.Now(), merged)
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
func (s *FuseService) mergeLeftover(transactions_0, transactions_1 []*domain.Transaction) []*domain.Transaction {
	// Placeholder for actual implementation of merging leftover transactions
	s.Logger.IPrintf(2, "Merging leftover transactions (status 0: %d, status 1: %d)\n", len(transactions_0), len(transactions_1))
	t0_map := s.getTransactionsMap(transactions_0)
	t1_map := s.getTransactionsMap(transactions_1)
	result := s.mergeMaps(t0_map, t1_map)
	// Log the number of merged leftover transactions
	s.Logger.IPrintf(2, "Merged %d leftover transactions (status 0: %d, status 1: %d)\n", len(result), len(transactions_0), len(transactions_1))
	return result
}

// getTransactionsMap creates a map from a slice of transactions based on their keys
func (s *FuseService) getTransactionsMap(transactions []*domain.Transaction) map[string]*domain.Transaction {
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

// mergeMaps merges two maps of transactions based on their keys, giving priority to transactions in the first map over those in the second map
func (s *FuseService) mergeMaps(t0_map, t1_map map[string]*domain.Transaction) []*domain.Transaction {
	result := []*domain.Transaction{}
	for key, t0 := range t0_map {
		if t1, exists := t1_map[key]; exists {
			if t0.TransactionDate == nil || t1.TransactionDate == nil {
				continue
			}
			if t0.TransactionDate.After(t1.TransactionDate.AddDate(0, 0, 3)) ||
				t0.TransactionDate.Before(t1.TransactionDate.AddDate(0, 0, -3)) {
				continue
			}
			domain.MergeManagement(t1, t0)
			t1.Cancel()
			t1.ReferenceID = &t0.ID
			t0.ReferenceID = &t1.ID
			result = append(result, t0)
			result = append(result, t1)
		}
	}
	return result
}