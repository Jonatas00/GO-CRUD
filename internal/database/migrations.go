package database

import "github.com/Jonatas00/GO-CRUD/internal/model"

func GenerateMigrations() {
	DB.AutoMigrate(&model.User{})
}
