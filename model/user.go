package model

import (
	"blog/database"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	Password string `gorm:"not null;unique"`
}

func GetAllUsers() ([]User, error) {
	var users []User
	result := database.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
