package config

import (
	"github.com/m1keru/pushbot/internal/databases"
	"github.com/m1keru/pushbot/internal/endpoints"
	"github.com/m1keru/pushbot/internal/logging"
	log "github.com/sirupsen/logrus"
	"os"

	"gopkg.in/yaml.v2"
)

// Config -- Global Config
type Config struct {
	Daemon       Daemon                 `yaml:"daemon"`
	Logger       logging.Logger         `yaml:"logger"`
	Alertmanager endpoints.Alertmanager `yaml:"alertmanager"`
	Rabbitmq     databases.RabbitMQ     `yaml:"rabbitmq"`
}

//Daemon - DaemonConfig
type Daemon struct {
	Port int `yaml:"port"`
}

//Setup - Setup
func (cfg *Config) Setup(filename *string) error {
	configFile, err := os.Open(*filename)
	if err != nil {
		log.Fatalf("Unable to read config file. Error:\n %v\n", err)
	}
	defer configFile.Close()
	decoder := yaml.NewDecoder(configFile)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatalf("Unable to Unmarshal Config, Error:\n %v\n", err)
	}
	log.Printf("Config: %+v", cfg)
	return nil
}
