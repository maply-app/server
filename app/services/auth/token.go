package auth

import (
	"maply/config"
	"maply/models"
	"maply/repository/managers"
	"maply/services/utils"
)

func GenerateToken(email, password string) (string, error) {
	var (
		token string
		user  models.User
		err   error
	)

	// Find a user in database
	user, err = managers.CheckUser(email, utils.HashPassword(password))
	if err != nil {
		return token, err
	}

	// Generate a token (JWT format)
	token, err = utils.GenerateJWT(user.ID, config.C.Auth.TTL)
	if err != nil {
		return token, err
	}
	return token, err
}
