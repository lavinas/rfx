package driven

import (
	// "fmt"
	"context"
	"time"

	"fuser/internal/core/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

// GormRepository is an adapter for GORM database operations
type GormRepository struct {
	DB  *gorm.DB
	ctx *context.Context
}

// NewGormRepository creates a new instance of GormRepository
func NewGormRepository(dns string, ctx *context.Context) (*GormRepository, error) {
	rep := &GormRepository{DB: nil, ctx: ctx}
	if err := rep.Connect(dns); err != nil {
		return nil, err
	}
	return rep, nil
}

// Connect establishes a connection to the database (placeholder for actual connection logic)
func (a *GormRepository) Connect(dns string) error {
	// Placeholder for actual connection logic, using GORM to connect to the database
	gConfig := gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // Disables all SQL logging
	}
	sqlDB, err := gorm.Open(postgres.Open(dns), &gConfig)
	if err != nil {
		return err
	}
	a.DB = sqlDB
	// Verify the connection by pinging the database
	return a.Ping()
}

// Close closes the database connection
func (a *GormRepository) Close() error {
	db, err := a.DB.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

// Ping checks the database connection
func (a *GormRepository) Ping() error {
	db, err := a.DB.DB()
	if err != nil {
		return err
	}
	return db.PingContext(*a.ctx)
}

// GetManagementTransactions retrieves Management transactions from the database
func (a *GormRepository) GetManagementTransactions(dt_transaction time.Time) ([]*domain.Management, error) {
	var transactions []*domain.Management
	start_date := dt_transaction.Format("2006-01-02") + " 00:00:00"
	end_date := dt_transaction.AddDate(0, 0, 1).Format("2006-01-02") + " 00:00:00"
	if err := a.DB.WithContext(*a.ctx).Where("dt_processamento >= ? AND dt_processamento < ?", start_date, end_date).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

// GetIntercamTransactions retrieves Intercam transactions from the database
func (a *GormRepository) GetIntercamTransactions(dt_transaction time.Time) ([]*domain.Intercam, error) {
	var transactions []*domain.Intercam
	start_date := dt_transaction.Format("2006-01-02") + " 00:00:00"
	end_date := dt_transaction.AddDate(0, 0, 1).Format("2006-01-02") + " 00:00:00"
	if err := a.DB.WithContext(*a.ctx).Where("dt_processamento >= ? AND dt_processamento < ?", start_date, end_date).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

// GetTransactionsByKey retrieves transactions by their keys from the database
func (a *GormRepository) GetTransactionsByKey(keys []string) ([]*domain.Transaction, error) {
	var transactions []*domain.Transaction
	if err := a.DB.WithContext(*a.ctx).Where("key1 IN ?", keys).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

// GetTransactionsByDateRangeAndStatus retrieves transactions by date range and status from the database
func (a *GormRepository) GetTransactionsByDateRangeAndStatus(start, end time.Time, status int) ([]*domain.Transaction, error) {
	var transactions []*domain.Transaction
	start_date := start.Format("2006-01-02") + " 00:00:00"
	end_date := end.AddDate(0, 0, 1).Format("2006-01-02") + " 00:00:00"
	if err := a.DB.WithContext(*a.ctx).Where("transaction_date >= ? AND transaction_date < ? AND status_id = ?", start_date, end_date, status).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

// Insert Transactions inserts a list of transactions into the database
func (a *GormRepository) InsertTransactions(transactions []*domain.Transaction) error {
	return a.DB.WithContext(*a.ctx).Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&transactions).Error
}
