package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/Jonatas00/GO-CRUD/internal/config"
)

func Connect() (*sql.DB, error) {
	var connectionString = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.DBCfg.Host, config.DBCfg.User, config.DBCfg.Password,
		config.DBCfg.Name, config.DBCfg.Port, config.DBCfg.SslMode,
	)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
