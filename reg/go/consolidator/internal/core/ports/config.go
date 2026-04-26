package ports

// Config defines the interface for configuration management in the consolidation service.
type Config interface {
	// GetDBData retrieves the database configuration data and populates the provided pointers with the respective values.
	GetDBData(host *string, port *int, user *string, password *string, dbname *string, sslmode *string,
		timezone *string, connect_timeout *int, source_schema *string, target_schema *string, bin_schema *string)
}
