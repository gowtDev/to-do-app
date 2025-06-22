package data

import (
	"errors"

	"github.com/gowtDev/to-do-app/utils"
	"gorm.io/gorm"
)

type User struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Email    string `gorm:"unique"`
	Password string
}

// CreateUser hashes password and stores user, returns JWT token
func (user *User) CreateUser() (string, error) {
	// Hash password before saving
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return "", err
	}
	user.Password = hashedPassword

	// Save user
	if err := DB.Create(user).Error; err != nil {
		return "", err
	}

	return utils.GenerateToken(user.Email)
}

// Login verifies credentials and returns JWT token if correct
func (user *User) Login() (string, error) {

	var dbUser User
	err := DB.Where("email = ?", user.Email).First(&dbUser).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("user not found")
		}
		return "", err
	}

	if !utils.CheckPasswordHash(user.Password, dbUser.Password) {
		return "", errors.New("invalid credentials")
	}

	return utils.GenerateToken(dbUser.Email)
}
