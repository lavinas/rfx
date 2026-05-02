package adapter

import (
	"encoding/json"
	"os"
)

// JsonConfig represents the configuration structure for the application
type JsonConfig struct {
	DB   JsonDBConfig   `json:"db"`
	Log  JsonLogConfig  `json:"log"`
	Cron JsonCronConfig `json:"cron"`
}

// DBConfig represents the database configuration structure
type JsonDBConfig struct {
	Host           string `json:"host"`
	Port           int    `json:"port"`
	User           string `json:"user"`
	Password       string `json:"password"`
	DBName         string `json:"dbname"`
	SSLMode        string `json:"sslmode"`
	TimeZone       string `json:"timezone"`
	ConnectTimeout int    `json:"connect_timeout"`
	SourceSchema   string `json:"source_schema"`
}

// CronConfig represents the cron configuration structure
type JsonCronConfig struct {
	Schedules     []string `json:"schedules"`
	TimeZone      string   `json:"timezone"`
	BacktrackDays int      `json:"backtrackdays"`
}

// LogConfig represents the logging configuration structure
type JsonLogConfig struct {
	Output string `json:"output"`
	Level  int    `json:"level"`
}

// LoadJsonConfig reads the configuration from a JSON file and unmarshals it into a JsonConfig struct
func NewConfig(path string) (*JsonConfig, error) {
	// Attempt to read the configuration file, if it fails, use the default configuration string
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	// Unmarshal the JSON data into the Config struct
	var cfg JsonConfig
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

// GetDBData returns the database configuration data as a JsonDBConfig struct
func (v *JsonConfig) GetDBData(host *string, port *int, user *string, password *string, dbname *string, sslmode *string,
	timezone *string, connect_timeout *int, sourceSchema *string) {
	*host = v.DB.Host
	*port = v.DB.Port
	*user = v.DB.User
	*password = v.DB.Password
	*dbname = v.DB.DBName
	*sslmode = v.DB.SSLMode
	*timezone = v.DB.TimeZone
	*connect_timeout = v.DB.ConnectTimeout
	*sourceSchema = v.DB.SourceSchema
}

// GetDBTimeZone returns the database time zone from the configuration
func (v *JsonConfig) GetDBTimeZone() string {
	return v.DB.TimeZone
}
