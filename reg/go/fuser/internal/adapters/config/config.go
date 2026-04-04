package config

import (
	"os"
	"encoding/json"
)

// Config represents the configuration structure for the application
type Config struct {
	DB DBConfig `json:"db"`
}

// DBConfig represents the database configuration structure
type DBConfig struct {
	DNS string `json:"dns"`
}


// GetDNS returns the DNS string from the DBConfig struct
func (v *Config) GetDNS() string {
	return v.DB.DNS
}

// LoadConfig reads the configuration from a YAML file and unmarshals it into a Config struct
func LoadConfig(path string) (*Config, error) {
	// Attempt to read the configuration file, if it fails, use the default configuration string
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	// Unmarshal the JSON data into the Config struct
	var cfg Config
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	// Return a pointer to the Config struct
	return &cfg, nil
}