package driven

import (
	"context"
	"fmt"
	"time"

	"fuser/internal/core/domain"
	"fuser/internal/core/ports"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

const (
	batchSizeInsertTransaction    = 2000
	batchSizeFindByKeyTransaction = 50000
)

// PostgresRepository is an adapter for GORM database operations
type PostgresRepository struct {
	DB           *gorm.DB
	ctx          *context.Context
	sourceSchema string
	targetSchema string
}

// NewPostgresRepository creates a new instance of PostgresRepository
func NewPostgresRepository(config ports.Config, ctx *context.Context) (*PostgresRepository, error) {
	rep := &PostgresRepository{DB: nil, ctx: ctx}
	var host, user, password, dbname, sslmode, timezone string
	var port, connect_timeout int
	config.GetDBData(&host, &port, &user, &password, &dbname, &sslmode, &timezone, &connect_timeout, &rep.sourceSchema, &rep.targetSchema)
	dns := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=%s connect_timeout=%d",
		host, port, user, password, dbname, sslmode, timezone, connect_timeout)

	if err := rep.Connect(dns); err != nil {
		return nil, err
	}
	return rep, nil
}

// Connect establishes a connection to the database (placeholder for actual connection logic)
func (a *PostgresRepository) Connect(dns string) error {
	// Placeholder for actual connection logic, using GORM to connect to the database
	gConfig := gorm.Config{
		Logger:      logger.Default.LogMode(logger.Silent), // Disables all SQL logging
		PrepareStmt: true,
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
func (a *PostgresRepository) Close() error {
	db, err := a.DB.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

// Ping checks the database connection
func (a *PostgresRepository) Ping() error {
	db, err := a.DB.DB()
	if err != nil {
		return err
	}
	return db.PingContext(*a.ctx)
}

// GetManagementTransactions retrieves Management transactions from the database
func (a *PostgresRepository) GetManagementTransactions(dt_transaction time.Time) ([]*domain.Management, error) {
	a.DB.Exec(fmt.Sprintf("SET search_path TO %s", a.sourceSchema))
	var transactions []*domain.Management
	start_date := dt_transaction.Format("2006-01-02") + " 00:00:00"
	end_date := dt_transaction.AddDate(0, 0, 1).Format("2006-01-02") + " 00:00:00"
	if err := a.DB.WithContext(*a.ctx).Where("dt_processamento >= ? AND dt_processamento < ?", start_date, end_date).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

// GetExchangeTransactions retrieves Exchange transactions from the database
func (a *PostgresRepository) GetExchangeTransactions(dt_transaction time.Time) ([]*domain.Exchange, error) {
	a.DB.Exec(fmt.Sprintf("SET search_path TO %s", a.sourceSchema))
	var transactions []*domain.Exchange
	start_date := dt_transaction.Format("2006-01-02") + " 00:00:00"
	end_date := dt_transaction.AddDate(0, 0, 1).Format("2006-01-02") + " 00:00:00"
	if err := a.DB.WithContext(*a.ctx).Where("dt_processamento >= ? AND dt_processamento < ?", start_date, end_date).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

// GetTransactionsByKeyBatch retrieves transactions by their keys from the database in batches
func (a *PostgresRepository) GetTransactionsByKey(keys []string) ([]*domain.Transaction, error) {
	a.DB.Exec(fmt.Sprintf("SET search_path TO %s", a.targetSchema))
	if len(keys) == 0 {
		return []*domain.Transaction{}, nil
	}
	var transactions []*domain.Transaction
	for i := 0; i < len(keys); i += batchSizeFindByKeyTransaction {
		end := i + batchSizeFindByKeyTransaction
		if end > len(keys) {
			end = len(keys)
		}
		batchKeys := keys[i:end]
		var batchTransactions []*domain.Transaction
		if err := a.DB.WithContext(*a.ctx).Where("key1 IN ?", batchKeys).Find(&batchTransactions).Error; err != nil {
			return nil, err
		}
		transactions = append(transactions, batchTransactions...)
	}

	return transactions, nil
}

// GetTransactionsByDateRangeAndStatus retrieves transactions by date range and status from the database
func (a *PostgresRepository) GetTransactionsByDateRangeAndStatus(start, end time.Time, status int) ([]*domain.Transaction, error) {
	a.DB.Exec(fmt.Sprintf("SET search_path TO %s", a.targetSchema))
	var transactions []*domain.Transaction
	start_date := start.Format("2006-01-02") + " 00:00:00"
	end_date := end.AddDate(0, 0, 1).Format("2006-01-02") + " 00:00:00"
	if err := a.DB.WithContext(*a.ctx).Where("transaction_date >= ? AND transaction_date < ? AND status_id = ?", start_date, end_date, status).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

// InsertTransactions inserts a batch of transactions into the database
func (a *PostgresRepository) InsertTransactions(transactions []*domain.Transaction) error {
	return a.DB.Transaction(func(tx *gorm.DB) error {
		a.DB.Exec(fmt.Sprintf("SET search_path TO %s", a.targetSchema))
		if err := a.DB.WithContext(*a.ctx).Clauses(clause.OnConflict{
			UpdateAll: true,
		}).CreateInBatches(&transactions, batchSizeInsertTransaction).Error; err != nil {
			return err
		}
		return nil
	})
}
