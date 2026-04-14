package driven

import (
	// "fmt"
	"context"
	"fmt"
	"time"

	source_domain "consolidator/internal/core/domain/source"
	target_domain "consolidator/internal/core/domain/target"
	"consolidator/internal/core/ports"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/clause"
)

// GormRepository is an adapter for GORM database operations
type GormRepository struct {
	DB  *gorm.DB
	ctx *context.Context
}

// NewGormRepository creates a new instance of GormRepository
func NewGormRepository(config ports.Config, ctx *context.Context) (*GormRepository, error) {
	rep := &GormRepository{DB: nil, ctx: ctx}
	var host, user, password, dbname, sslmode, timezone string
	var port, connect_timeout int
	config.GetDBData(&host, &port, &user, &password, &dbname, &sslmode, &timezone, &connect_timeout)
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

	if err := a.DB.WithContext(*a.ctx).Where("period_date >= ? AND period_date < ? and status_id = 2", start_date, end_date).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

// SaveDesconto saves the consolidated Desconto data to the database
func (a *GormRepository) SaveDesconto(desconto []*target_domain.Desconto) error {
	// Placeholder for actual save logic, using GORM to save the consolidated Desconto data to the database
	return a.DB.WithContext(*a.ctx).Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&desconto).Error
}

// DeleteDesconto deletes existing consolidated Desconto data from the database for a specific date
func (a *GormRepository) DeleteDesconto(year int, quarter int) error {
	// delete all records of desconto for the specified year and quarter
	if err := a.DB.Where("year = ? AND quarter = ?", year, quarter).Delete(&target_domain.Desconto{}).Error; err != nil {
		return err
	}
	return nil
}