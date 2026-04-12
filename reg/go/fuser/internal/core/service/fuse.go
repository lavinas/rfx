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
func (s *FuseService) Run(start_date time.Time, end_date time.Time, focus string) error {
	s.Logger.Println("Running FuseService...")
	// Process intercam
	for date := start_date; !date.After(end_date); date = date.AddDate(0, 0, 1) {
		s.Logger.Printf("Processing intercam for date: %s\n", date.Format("2006-01-02"))
		if focus == "all" || focus == "intercam" {
			err := s.processIntercam(date)
			if err != nil {
				return err
			}
		}
		if focus == "all" || focus == "management" {
			err := s.processManagement(date)
			if err != nil {
				return err
			}
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
	byKey, err := s.getTransactionsByKey("intercam", date, transactions)
	if err != nil {
		return err
	}
	merged := s.mergeTransactions("intercam", date, transactions, byKey)
	err = s.insertTransactions("intercam", date, merged)
	if err != nil {
		return err
	}
	s.Logger.Printf("Finished processing Intercam transactions for date %s\n", date.Format("2006-01-02"))
	return nil
}

// getManagement is a helper method to fetch Management transactions for a specific date
func (s *FuseService) processManagement(date time.Time) error {
	s.Logger.Printf("Processing Management transactions for date %s\n", date.Format("2006-01-02"))
	transactions, err := s.getManagementTransactions(date)
	if err != nil {
		return err
	}
	byKey, err := s.getTransactionsByKey("management", date, transactions)
	if err != nil {
		return err
	}
	merged := s.mergeTransactions("management", date, transactions, byKey)
	err = s.insertTransactions("management", date, merged)
	if err != nil {
		return err
	}
	s.Logger.Printf("Finished processing Management transactions for date %s\n", date.Format("2006-01-02"))
	return nil
}


// getIntercamTransactions is a helper method to fetch Intercam transactions for a specific date
func (s *FuseService) getIntercamTransactions(date time.Time) ([]*domain.Transaction, error) {
	intercams, err := s.Repository.GetIntercamTransactions(date)
	if err != nil {
		s.Logger.Printf("Error fetching intercamtransactions for date %s: %v\n", date.Format("2006-01-02"), err)
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
		s.Logger.Printf("Error fetching management transactions for date %s: %v\n", date.Format("2006-01-02"), err)
		return nil, err
	}
	transactions := []*domain.Transaction{}
	for _, management := range managements {
		transactions = append(transactions, management.Translate())
	}
	return transactions, nil
}

// getTransactionsByKey is a helper method to fetch transactions by their keys
func (s *FuseService) getTransactionsByKey(transType string, transDate time.Time,transactions []*domain.Transaction) ([]*domain.Transaction, error) {
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
			s.Logger.Printf("Fetched %d %s transactions by keys for %s date %s (%d/%d)\n", len(repTrans), transType, transType, transDate.Format("2006-01-02"), count, total)
			repTransactions = append(repTransactions, repTrans...)
			keys = []string{}
		}
	}
	if len(keys) > 0 {
		repTrans, err := s.Repository.GetTransactionsByKey(keys)
		if err != nil {
			return nil, err
		}
		s.Logger.Printf("Fetched %d %s transactions by keys for %s date %s (%d/%d)\n", len(repTrans), transType, transType, transDate.Format("2006-01-02"), count, total)
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
				s.Logger.Printf("Unknown transaction type: %s\n", transType)
			}
			merged = append(merged, repoTrans)
		} else {
			merged = append(merged, localTrans)
		}
	}
	s.Logger.Printf("Merged %s transactions for date %s (local: %d, repository: %d, merged: %d)\n", transType, transDate.Format("2006-01-02"), len(localTransactions), len(repositoryTransactions), len(merged))
	return merged
}

// insertTransactions is a helper method to insert a batch of transactions into the repository
func (s *FuseService) insertTransactions(transType string, transDate time.Time, transactions []*domain.Transaction) error {
	count := 0
	total := len(transactions)
	lot := []*domain.Transaction{}
	for _, transaction := range transactions {
		transaction.PrepareForInsert()
		lot = append(lot, transaction)
		count++
		if count%2000 == 0 {
			if err := s.Repository.InsertTransactions(lot); err != nil {
				s.Logger.Printf("Error inserting %s transactions for date %s: %v\n", transType, transDate.Format("2006-01-02"), err)
				return err
			}
			s.Logger.Printf("Inserted %s transactions for date %s (%d/%d)\n", transType, transDate.Format("2006-01-02"), count, total)
			lot = []*domain.Transaction{}
		}
	}
	if len(lot) > 0 {
		if err := s.Repository.InsertTransactions(lot); err != nil {
			s.Logger.Printf("Error inserting %s transactions for date %s: %v\n", transType, transDate.Format("2006-01-02"), err)
			return err
		}
		s.Logger.Printf("Inserted %s transactions for date %s (%d/%d)\n", transType, transDate.Format("2006-01-02"), count, total)
	}
	return nil
}


