package services

import (
	db "github.com/Kallox/CRIS-Backend/internal/database"
	"github.com/Kallox/CRIS-Backend/internal/models"
)

func GetAllUsers() ([]models.User, error) {
	// Implement the logic to get all users from the database
	var users []models.User

	if err := db.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func CreateUser(user *models.User) error {
	// Implement the logic to create a new user
	if err := db.DB.Create(user).Error; err != nil {
		return err
	}

	return nil
}
