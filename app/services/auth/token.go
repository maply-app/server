package auth

import (
	"maply/config"
	utils2 "maply/core/utils"
	"maply/models"
	"maply/repository/managers/auth"
)

func GenerateToken(email, password string) (string, error) {
	var (
		token string
		user  models.User
		err   error
	)

	user, err = auth.CheckUser(email, utils2.HashPassword(password))
	if err != nil {
		return token, err
	}

	token, err = utils2.GenerateJWT(user.ID, config.C.Auth.TTL)
	if err != nil {
		return token, err
	}
	return token, err
}
