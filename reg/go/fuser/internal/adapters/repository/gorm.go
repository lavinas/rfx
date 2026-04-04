package repository

import (
	"context"
	"time"

	"fuser/internal/core/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	sqlDB, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
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
func (a *GormRepository) GetManagementTransactions(dt_transaction time.Time) ([]domain.Management, error) {
	var transactions []domain.Management
	if err := a.DB.WithContext(*a.ctx).Where("dt_processamento = ?", dt_transaction).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

// GetWebserviceTransactions retrieves Webservice transactions from the database
func (a *GormRepository) GetWebserviceTransactions(dt_transaction time.Time, page int) ([]domain.Webservice, error) {
	var transactions []domain.Webservice
	if err := a.DB.WithContext(*a.ctx).Where("dt_processamento = ?", dt_transaction).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

// GetIntercamTransactions retrieves Intercam transactions from the database
func (a *GormRepository) GetIntercamTransactions(dt_transaction time.Time) ([]domain.Intercam, error) {
	var transactions []domain.Intercam
	if err := a.DB.WithContext(*a.ctx).Where("dt_processamento = ?", dt_transaction).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

// Insert Transactions inserts a list of transactions into the database
func (a *GormRepository) InsertTransactions(transactions []domain.Transaction) error {
	return a.DB.WithContext(*a.ctx).Create(&transactions).Error
}
