package driven

import (
	// "fmt"
	"context"
	"fmt"
	"time"

	source_domain "consolidator/internal/core/domain/source"
	"consolidator/internal/core/ports"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

const (
	batchSizeInsertTransaction = 2000
)

// GormRepository is an adapter for GORM database operations
type GormRepository struct {
	DB            *gorm.DB
	ctx           *context.Context
	rawdata_schema string
	transaction_schema string
	consolidator_schema string
	bin_schema    string
}

// NewGormRepository creates a new instance of GormRepository
func NewGormRepository(config ports.Config, ctx *context.Context) (*GormRepository, error) {
	rep := &GormRepository{DB: nil, ctx: ctx}
	var host, user, password, dbname, sslmode, timezone string
	var port, connect_timeout int
	config.GetDBData(&host, &port, &user, &password, &dbname, &sslmode, &timezone, &connect_timeout, &rep.rawdata_schema, &rep.transaction_schema, &rep.consolidator_schema, &rep.bin_schema)
	dns := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=%s connect_timeout=%d", host, port, user, password, dbname, sslmode, timezone, connect_timeout)
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

// GetTransactionsByDate retrieves transactions from the database for a specific date
func (a *GormRepository) GetTransactionsByDate(date time.Time) ([]*source_domain.Transaction, error) {
	var transactions []*source_domain.Transaction
	start_date := date.Format("2006-01-02") + " 00:00:00"
	end_date := date.AddDate(0, 0, 1).Format("2006-01-02") + " 00:00:00"

	a.DB.Exec(fmt.Sprintf("SET search_path TO %s", a.transaction_schema))
	if err := a.DB.WithContext(*a.ctx).Where("period_date >= ? AND period_date < ? and status_id = 2", start_date, end_date).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

// GetBins retrieves BIN information from the database
func (a *GormRepository) GetBins() ([]*source_domain.Bin, error) {
	var bins []*source_domain.Bin

	a.DB.Exec(fmt.Sprintf("SET search_path TO %s", a.bin_schema))
	if err := a.DB.WithContext(*a.ctx).Find(&bins).Error; err != nil {
		return nil, err
	}
	return bins, nil
}

// GetEstablishments retrieves establishment information from the database
func (a *GormRepository) GetEstablishments() ([]*source_domain.Establishment, error) {
	var establishments []*source_domain.Establishment
	a.DB.Exec(fmt.Sprintf("SET search_path TO %s", a.rawdata_schema))
	if err := a.DB.WithContext(*a.ctx).Find(&establishments).Error; err != nil {
		return nil, err
	}
	return establishments, nil
}

// GetTerminals retrieves terminal information from the database
func (a *GormRepository) GetTerminals() ([]*source_domain.Terminal, error) {
	var terminals []*source_domain.Terminal
	a.DB.Exec(fmt.Sprintf("SET search_path TO %s", a.rawdata_schema))
	if err := a.DB.WithContext(*a.ctx).Find(&terminals).Error; err != nil {
		return nil, err
	}
	return terminals, nil
}

// GeneralDelete is a helper function to delete records from the database for a specific year and quarter
func (a *GormRepository) Delete(model interface{}, year int, quarter int) error {
	a.DB.Exec(fmt.Sprintf("SET search_path TO %s", a.consolidator_schema))
	if err := a.DB.Where("year = ? AND quarter = ?", year, quarter).Delete(model).Error; err != nil {
		return err
	}
	return nil
}

// GeneralSave is a helper function to save records to the database with conflict handling
func (a *GormRepository) Save(model interface{}) error {
	a.DB.Exec(fmt.Sprintf("SET search_path TO %s", a.consolidator_schema))
	return a.DB.WithContext(*a.ctx).Clauses(clause.OnConflict{
		UpdateAll: true,
	}).CreateInBatches(model, batchSizeInsertTransaction).Error
}
