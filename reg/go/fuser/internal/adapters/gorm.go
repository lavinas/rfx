package adapters

import (
	"context"
	"time"

	"fuser/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GormAdapter is an adapter for GORM database operations
type GormAdapter struct {
	DB *gorm.DB
	ctx *context.Context
}

// NewGormAdapter creates a new instance of GormAdapter
func NewGormAdapter(db *gorm.DB, ctx *context.Context) *GormAdapter {
	return &GormAdapter{DB: db, ctx: ctx}
}

// Connect establishes a connection to the database (placeholder for actual connection logic)
func (a *GormAdapter) Connect(dns string) error {
	// Placeholder for actual connection logic, using GORM to connect to the database
	sqlDB, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		return err
	}
	a.DB = sqlDB
	// Verify the connection by pinging the database
	return a.Ping()
}

// Ping checks the database connection
func (a *GormAdapter) Ping() error {
	db, err := a.DB.DB()
	if err != nil {
		return err
	}
	return db.PingContext(*a.ctx)
}

// GetManagementTransactions retrieves Management transactions from the database
func (a *GormAdapter) GetManagementTransactions(dt_transaction time.Time) ([]domain.Management, error) {
	var transactions []domain.Management
	if err := a.DB.WithContext(*a.ctx).Where("dt_processamento = ?", dt_transaction).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

// GetWebserviceTransactions retrieves Webservice transactions from the database
func (a *GormAdapter) GetWebserviceTransactions(dt_transaction time.Time, page int) ([]domain.Webservice, error) {
	var transactions []domain.Webservice
	if err := a.DB.WithContext(*a.ctx).Where("dt_processamento = ?", dt_transaction).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

// Insert Transactions inserts a list of transactions into the database
func (a *GormAdapter) InsertTransactions(transactions []domain.Transaction) error {
	return a.DB.WithContext(*a.ctx).Create(&transactions).Error
}
