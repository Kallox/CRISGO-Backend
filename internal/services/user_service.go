package services

import (
	db "github.com/Kallox/CRIS-Backend/internal/database"
	"github.com/Kallox/CRIS-Backend/internal/models"
	"github.com/Kallox/CRIS-Backend/internal/services/auth"
)

func GetAllUsers() ([]models.User, error) {
	// Implement the logic to get all users from the database
	var users []models.User

	if err := db.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func GetUser(id string) (*models.User, error) {
	// Implement the logic to get a user by ID
	var user models.User

	if err := db.DB.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func CreateUser(user *models.User) error {
	// Implement the logic to create a new user

	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	if err := db.DB.Create(user).Error; err != nil {
		return err
	}

	return nil
}
