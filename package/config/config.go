package config

import (
	"github.com/spf13/viper"
)

// Config Struct that maps the YAML structure
type Config struct {
	Server struct {
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
	} `mapstructure:"server"`
	Database struct {
		Host     string `mapstructure:"host"`
		Name     string `mapstructure:"name"`
		Username string `mapstructure:"username"`
		Port     string `mapstructure:"port"`
		Password string `mapstructure:"password"`
	} `mapstructure:"database"`
}

func ReadConfig(path string) (Config, error) {
	var cfg Config
	// Thiết lập để tự động đọc biến môi trường
	viper.AutomaticEnv()

	// Đọc file cấu hình
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return cfg, err
	}

	// Thực hiện unmarshal từ cấu hình
	if err := viper.Unmarshal(&cfg); err != nil {
		return cfg, err
	}

	// Thay thế các giá trị trong cấu hình bằng biến môi trường
	if cfg.Database.Host == "${DATABASE_HOST}" {
		cfg.Database.Host = viper.GetString("DATABASE_HOST")
	}
	if cfg.Database.Name == "${DATABASE_NAME}" {
		cfg.Database.Name = viper.GetString("DATABASE_NAME")
	}
	if cfg.Database.Username == "${DATABASE_USERNAME}" {
		cfg.Database.Username = viper.GetString("DATABASE_USERNAME")
	}
	if cfg.Database.Password == "${DATABASE_PASSWORD}" {
		cfg.Database.Password = viper.GetString("DATABASE_PASSWORD")
	}

	return cfg, nil
}
