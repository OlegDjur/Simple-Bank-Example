package postgres

import (
	"database/sql"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	SereverAddres string `mapstructure:"SERVER_ADDRESS"`
}

func InitDB() *sql.DB {
	config, err := loadConfig(".")
	if err != nil {
		log.Fatal("cannot load")
	}

	db, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	return db
}

func loadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
