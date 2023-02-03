package auth

import (
	"maply/models"
	"maply/repository"
)

func CreateUser(u *models.User) (string, error) {
	result := repository.DB.Create(&u)
	return u.ID, result.Error
}

func CheckUser(email, password string) (models.User, error) {
	var user models.User
	result := repository.DB.Where(
		"email = ? AND password = ?",
		email,
		password,
	).First(&user)
	return user, result.Error
}
