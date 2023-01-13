package managers

import (
	"maply/models"
	"maply/repository"
)

// CreateUser create new user
func CreateUser(u *models.User) (string, error) {
	result := repository.DB.Create(&u)
	return u.ID, result.Error
}

// CheckUser get user by login and password
func CheckUser(email, password string) (models.User, error) {
	var user models.User
	result := repository.DB.Where(
		"email = ? AND password = ?",
		email,
		password,
	).First(&user)
	return user, result.Error
}
