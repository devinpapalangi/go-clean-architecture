package pkg

import (
	"log"
	"os"
	"sync"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type JWTConfig struct {
	Secret string `mapstructure:"secret"`
	Expire int    `mapstructure:"expire"`
}

type AppConfig struct {
	Name        string `mapstructure:"name"`
	Environment string `mapstructure:"environment"`
	Debug       bool   `mapstructure:"debug"`
	LogLevel    string `mapstructure:"log_level"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
}

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	App      AppConfig      `mapstructure:"app"`
	Database DatabaseConfig `mapstructure:"database"`
	JWT      JWTConfig      `mapstructure:"jwt"`
}

var (
	cfg  *Config
	once sync.Once
)

func LoadConfig() *Config {
	once.Do(
		func() {
			v := viper.New()

			currentEnv := os.Getenv("GO_ENV")

			if currentEnv == "" {
				currentEnv = "local"
			}

			v.SetConfigFile("./config-" + currentEnv + ".yaml")
			v.SetConfigType("yaml")

			if err := v.ReadInConfig(); err != nil {
				log.Fatalf("Error reading config file: %v", err)
			}

			cfg = &Config{}
			if err := v.Unmarshal(cfg); err != nil {
				log.Fatalf("Unable to decode config: %v", err)
			}

		},
	)
	return cfg
}

func SetDefaults(v *viper.Viper) {
	v.SetDefault("server.host", "localhost")
	v.SetDefault("server.port", 8080)

	v.SetDefault("database.host", "localhost")
	v.SetDefault("database.port", 5432)
	v.SetDefault("database.sslmode", "disable")

	v.SetDefault("app.environment", "development")
	v.SetDefault("app.debug", false)
	v.SetDefault("app.log_level", "info")

	v.AutomaticEnv()
}

func GetConfig() *Config {
	if cfg == nil {
		log.Fatal("Configuration not loaded. Call Load() first.")
	}
	return cfg
}
