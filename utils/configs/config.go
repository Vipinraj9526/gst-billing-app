package configs

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// PostgresConfig represents Postgres database configuration
type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	SSlMode  string `yaml:"sslmode"`
	TimeZone string `yaml:"timezone"`
}

type ApplicationConfig struct {
	JwtSecretKey string `yaml:"jwtsecretkey"`
}

// Config represents the overall application configuration
type Config struct {
	Postgres    PostgresConfig    `yaml:"postgres"`
	Application ApplicationConfig `yaml:"application"`
}

// LoadConfig loads configuration from the given file path
func LoadConfig(filePath string) (Config, error) {
	var config Config

	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("Failed to read YAML file %s: %v", filePath, err)
		return config, err
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Printf("Failed to unmarshal YAML file %s: %v", filePath, err)
		return config, err
	}

	return config, nil
}
