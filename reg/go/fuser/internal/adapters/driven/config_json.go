package driven

import (
	"fmt"
	"encoding/json"
	"os"
)

// JsonConfig represents the configuration structure for the application
type JsonConfig struct {
	DB JsonDBConfig `json:"db"`
}

// DBConfig represents the database configuration structure
type JsonDBConfig struct {
	Host	 string `json:"host"`
	Port	 int    `json:"port"`
	User	 string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
	SSLMode  string `json:"sslmode"`
	TimeZone string `json:"timezone"`
}

// GetDNS returns the DNS string from the DBConfig struct
func (v *JsonConfig) GetDNS() string {
	dns := "postgresql://%s:%s@%s:%d/%s?sslmode=%s&TimeZone=%s"
	return fmt.Sprintf(dns, v.DB.User, v.DB.Password, v.DB.Host, v.DB.Port, v.DB.DBName, v.DB.SSLMode, v.DB.TimeZone)
}

// GetDBTimeZone returns the TimeZone from the DBConfig struct
func (v *JsonConfig) GetDBTimeZone() string {
	return v.DB.TimeZone
}

// LoadJsonConfig reads the configuration from a JSON file and unmarshals it into a JsonConfig struct
func LoadJsonConfig(path string) (*JsonConfig, error) {
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
