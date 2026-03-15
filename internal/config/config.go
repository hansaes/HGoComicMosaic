package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App      AppConfig      `yaml:"app"`
	HTTP     HTTPConfig     `yaml:"http"`
	Database DatabaseConfig `yaml:"database"`
	Auth     AuthConfig     `yaml:"auth"`
	Storage  StorageConfig  `yaml:"storage"`
	Log      LogConfig      `yaml:"log"`
}

type AppConfig struct {
	Name string `yaml:"name" env:"APP_NAME" env-default:"GoComicMosaic"`
	Env  string `yaml:"env" env:"APP_ENV" env-default:"local"`
}

type HTTPConfig struct {
	Port         int           `yaml:"port" env:"HTTP_PORT" env-default:"8080"`
	ReadTimeout  time.Duration `yaml:"read_timeout" env:"HTTP_READ_TIMEOUT" env-default:"10s"`
	WriteTimeout time.Duration `yaml:"write_timeout" env:"HTTP_WRITE_TIMEOUT" env-default:"10s"`
}

type DatabaseConfig struct {
	DSN string `yaml:"dsn" env:"DB_DSN"`
}

type AuthConfig struct {
	JWTSecret string        `yaml:"jwt_secret" env:"JWT_SECRET"`
	JWTIssuer string        `yaml:"jwt_issuer" env:"JWT_ISSUER" env-default:"gocomicmosaic"`
	JWTExpire time.Duration `yaml:"jwt_expire" env:"JWT_EXPIRE" env-default:"24h"`
}

type StorageConfig struct {
	RootDir string `yaml:"root_dir" env:"STORAGE_ROOT_DIR" env-default:"./storage"`
}

type LogConfig struct {
	Level     string `yaml:"level" env:"LOG_LEVEL" env-default:"info"`
	Format    string `yaml:"format" env:"LOG_FORMAT" env-default:"text"`
	Output    string `yaml:"output" env:"LOG_OUTPUT" env-default:"stdout"`
	AddSource bool   `yaml:"add_source" env:"LOG_ADD_SOURCE" env-default:"false"`
}

func Load() (*Config, error) {
	cfg := &Config{}
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "configs/config.local.yaml"
	}

	if err := cleanenv.ReadConfig(configPath, cfg); err != nil {
		return nil, fmt.Errorf("read config file failed : %w", err)
	}

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, fmt.Errorf("read env failed: %w", err)
	}
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	return cfg, nil

}

func (c *Config) Validate() error {
	if c.Database.DSN == "" {
		return fmt.Errorf("DB_DSN is required")
	}

	if c.Auth.JWTSecret == "" {
		return fmt.Errorf("JWT_SECRET is required")
	}

	level := strings.ToLower(c.Log.Level)
	if level != "debug" && level != "info" && level != "warn" && level != "error" {
		return fmt.Errorf("LOG_LEVEL must be one of: debug, info, warn, error")
	}

	format := strings.ToLower(c.Log.Format)
	if format != "json" && format != "text" {
		return fmt.Errorf("LOG_FORMAT must be one of: text , json")
	}

	output := strings.ToLower(c.Log.Output)
	if output != "stdout" && output != "stderr" {
		return fmt.Errorf("LOG_OUTPUT must be one of: stdout , stderr")
	}

	return nil
}
