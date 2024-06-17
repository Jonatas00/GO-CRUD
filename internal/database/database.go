package database

import (
	"fmt"
	"log"

	"github.com/Jonatas00/GO-anime-list/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var connectionString = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.DBCfg.DB_HOST, config.DBCfg.DB_USER, config.DBCfg.DB_PASSWORD,
		config.DBCfg.DB_NAME, config.DBCfg.DB_PORT, config.DBCfg.DB_SSLMODE,
	)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB = db
}
