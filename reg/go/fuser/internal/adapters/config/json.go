package config

import (
	"encoding/json"
	"os"
)

// JsonConfig represents the configuration structure for the application
type JsonConfig struct {
	DB JsonDBConfig `json:"db"`
}

// DBConfig represents the database configuration structure
type JsonDBConfig struct {
	DNS string `json:"dns"`
}

// GetDNS returns the DNS string from the DBConfig struct
func (v *JsonConfig) GetDNS() string {
	return v.DB.DNS
}

// LoadConfig reads the configuration from a YAML file and unmarshals it into a Config struct
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
	// Return a pointer to the Config struct
	return &cfg, nil
}
