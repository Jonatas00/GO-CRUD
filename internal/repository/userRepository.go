package repository

import (
	"github.com/Jonatas00/GO-CRUD/internal/database"
	"github.com/Jonatas00/GO-CRUD/internal/model"
	"gorm.io/gorm"
)

func CreateUser(user *model.User) *gorm.DB {
	result := database.DB.Create(user)

	return result
}

func GetUsers() *gorm.DB {
	var users []model.User

	result := database.DB.Select("id", "name", "email", "password", "created_at", "updated_at").First(&users)

	return result
}
