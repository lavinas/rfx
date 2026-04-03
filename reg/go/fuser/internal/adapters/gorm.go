package adapters

import (
	"context"
	"gorm.io/gorm"
	"fuser/internal/domain"
)

// GormAdapter is an adapter for GORM database operations
type GormAdapter struct {
	DB *gorm.DB
}

// NewGormAdapter creates a new instance of GormAdapter
func NewGormAdapter(db *gorm.DB) *GormAdapter {
	return &GormAdapter{DB: db}
}

// GetIntercamTransactions retrieves Intercam transactions from the database
func (a *GormAdapter) GetIntercamTransactions(ctx context.Context) ([]domain.Intercam, error) {
	var transactions []domain.Intercam
	if err := a.DB.WithContext(ctx).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

// GetManagementTransactions retrieves Management transactions from the database
func (a *GormAdapter) GetManagementTransactions(ctx context.Context) ([]domain.Management, error) {
	var transactions []domain.Management
	if err := a.DB.WithContext(ctx).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

// GetWebserviceTransactions retrieves Webservice transactions from the database
func (a *GormAdapter) GetWebserviceTransactions(ctx context.Context) ([]domain.Webservice, error) {
	var transactions []domain.Webservice
	if err := a.DB.WithContext(ctx).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

// Insert Transactions inserts a list of transactions into the database
func (a *GormAdapter) InsertTransactions(ctx context.Context, transactions []domain.Transaction) error {
	return a.DB.WithContext(ctx).Create(&transactions).Error
}
