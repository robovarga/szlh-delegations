package repository

import (
	"database/sql"
	"fmt"

	"github.com/robovarga/szlh-delegations/internal/config"

	_ "github.com/lib/pq"
)

func NewDBConnection(conf *config.DatabaseConfig) (*sql.DB, error) {
	db, err := sql.Open(conf.DriverName, conf.DatabaseURI)

	if err != nil {
		return nil, fmt.Errorf("can't connect to db: %s", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("can't connect to db: %s", err)
	}

	return db, nil
}
