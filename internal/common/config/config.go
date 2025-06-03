package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string     `yaml:"env"`
	HTTPServer HTTPServer `yaml:"http_server"`
	Datebase   Datebase   `yaml:"database"`
	RabbitMQ   RabbitMQ   `yaml:"kafka"`
	Logger     Logger     `yaml:"logger"`
}

type HTTPServer struct {
	Address     string        `yaml:"address"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
	User        string        `yaml:"user"`
	Password    string        `yaml:"password"`
}

type Datebase struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}

type RabbitMQ struct {
	Brokers []string `yaml:"brokers"`
	Topic   string   `yaml:"topic"`
	GroupID string   `yaml:"group_id"`
}

type Logger struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
}

func MustLoad() *Config {
	const configPath = "/home/abdu1bari/go/projects/transaction-worker/internal/common/config/config.yaml"

	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("config file %s does not exist: %v", configPath, err)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	log.Println("config loaded successfully")
	return &cfg
}
