package service

import (
	"time"

	"fuser/internal/core/domain"
	"fuser/internal/ports"
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
func (s *FuseService) Run(start_date time.Time, end_date time.Time) error {
	s.Logger.Println("Running FuseService...")
	// Process intercam
	for date := start_date; !date.After(end_date); date = date.AddDate(0, 0, 1) {
		s.Logger.Printf("Processing intercam for date: %s\n", date.Format("2006-01-02"))
		// getting intercam transactions for the date
		err := s.processIntercam(date)
		if err != nil {
			return err
		}
	}
	s.Logger.Println("Finished processing transactions.")
	return nil
}

// getIntercam is a helper method to fetch Intercam transactions for a specific date
func (s *FuseService) processIntercam(date time.Time) error {
	s.Logger.Printf("Processing Intercam transactions for date %s\n", date.Format("2006-01-02"))
	transactions, err := s.getIntercamTransactions(date)
	if err != nil {
		return err
	}
	byKey, err := s.getTransactionsByKey(transactions)
	if err != nil {
		return err
	}
	merged := s.mergeTransactions("intercam", transactions, byKey)
	err = s.insertTransactions(merged)
	if err != nil {
		return err
	}
	s.Logger.Printf("Finished processing Intercam transactions for date %s\n", date.Format("2006-01-02"))
	return nil
}


// getIntercamTransactions is a helper method to fetch Intercam transactions for a specific date
func (s *FuseService) getIntercamTransactions(date time.Time) ([]*domain.Transaction, error) {
	intercams, err := s.Repository.GetIntercamTransactions(date)
	if err != nil {
		s.Logger.Printf("Error fetching transactions for date %s: %v\n", date.Format("2006-01-02"), err)
		return nil,  err
	}
	transactions := []*domain.Transaction{}
	for _, intercam := range intercams {
		transactions = append(transactions, intercam.GetTransaction())
	}
	return transactions, nil
}

// getTransactionsByKey is a helper method to fetch transactions by their keys
func (s *FuseService) getTransactionsByKey(transactions []*domain.Transaction) ([]*domain.Transaction, error) {
	repTransactions := []*domain.Transaction{}
	keys := []string{}
	count := 0
	total := len(transactions)
	for _, transaction := range transactions {
		keys = append(keys, transaction.Key1)
		count++
		if count%2000 == 0 {
			s.Logger.Printf("Fetching transactions by keys: %d/%d\n", count, total)
			repTrans, err := s.Repository.GetTransactionsByKey(keys)
			if err != nil {
				return nil, err
			}
			repTransactions = append(repTransactions, repTrans...)
			keys = []string{}
		}
	}
	if len(keys) > 0 {
		s.Logger.Printf("Fetching transactions by keys: %d/%d\n", count, total)
		repTrans, err := s.Repository.GetTransactionsByKey(keys)
		if err != nil {
			return nil, err
		}
		repTransactions = append(repTransactions, repTrans...)
	}
	return repTransactions, nil
}

// mergeTransactions is a helper method to merge Intercam transactions with existing transactions in the repository
func (s *FuseService) mergeTransactions(transType string, localTransactions []*domain.Transaction, repositoryTransactions []*domain.Transaction) []*domain.Transaction {
	merged := []*domain.Transaction{}
	repoMap := make(map[string]*domain.Transaction)
	for _, repoTrans := range repositoryTransactions {
		repoMap[repoTrans.Key1] = repoTrans
	}
	for _, localTrans := range localTransactions {
		if repoTrans, exists := repoMap[localTrans.Key1]; exists {
			if transType == "intercam" {
				domain.MergeIntercam(localTrans, repoTrans)
			}
			merged = append(merged, repoTrans)
		} else {
			merged = append(merged, localTrans)
		}
	}
	s.Logger.Printf("Merged %d transactions (local: %d, repository: %d)\n", len(merged), len(localTransactions), len(repositoryTransactions))
	return merged
}

// CommitTransactions is a helper method to insert a batch of transactions into the repository
func (s *FuseService) insertTransactions(transactions []*domain.Transaction) error {
	count := 0
	total := len(transactions)
	lot := []*domain.Transaction{}
	for _, transaction := range transactions {
		lot = append(lot, transaction)
		count++
		if count%2000 == 0 {
			if err := s.Repository.InsertTransactions(lot);err != nil {
				s.Logger.Printf("Error inserting transactions: %v\n", err)
				return err
			}
	        s.Logger.Printf("Inserted %d/%d transactions\n", count, total)
			lot = []*domain.Transaction{}
		}
	}
	if len(lot) > 0 {
		 if err := s.Repository.InsertTransactions(lot); err != nil {
			s.Logger.Printf("Error inserting transactions: %v\n", err)
			return err
		}
	}
	return nil
}


