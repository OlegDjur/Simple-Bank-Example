package postgres

import (
	"database/sql"
	"log"
	"sbank/config"

	_ "github.com/lib/pq"
)

func InitDB(config *config.Config) *sql.DB {
	db, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	return db
}
