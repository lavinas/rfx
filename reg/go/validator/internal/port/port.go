package port

// report domain interface
type Report interface {
	Validate() error
	GetParsedFile(filename string) (map[string]Report, error)
	GetDB(repo Repository) (map[string]Report, error)
	String() string
	GetName() string
}

// repository domain interface
type Repository interface {
	FindAll(dest interface{}, limit int, offset int, orderBy string, conditions ...interface{}) error
	FindByPrimaryKey(dest interface{}, keyName string, keyValue interface{}) error
}

// Config interface defines methods to get configuration data
type Config interface {
	GetDBData(host *string, port *int, user *string, password *string, dbname *string, sslmode *string,
		timezone *string, connect_timeout *int, sourceSchema *string, targetSchema *string)
	GetLogData(output *string, level *int)
}

// Service interface defines methods to interact with the service
type Service interface {
	ExecuteAll() error
}