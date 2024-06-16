package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	ServerConfig Config
	Logger       *zap.Logger
	EasterDB     *gorm.DB
)

type MysqlConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Charset  string `yaml:"charset"`
}

type Config struct {
	Db  MysqlConfig `yaml:"db"`
	Log LogConfig   `yaml:"log"`
}

type LogConfig struct {
	Path  string `yaml:"path"`
	File  string `yaml:"file"`
	Level string `yaml:"level"`
}
