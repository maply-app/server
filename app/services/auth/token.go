package auth

import (
	"fmt"
	"maply/config"
	utils "maply/core/utils"
	"maply/models"
	"maply/repository/managers/auth"
)

func GenerateToken(email, password string) (string, error) {
	var (
		token string
		user  models.User
		err   error
	)

	user, err = auth.CheckUser(email, utils.HashPassword(password))
	if err != nil {
		return token, err
	}

	token, err = utils.GenerateJWT(user.ID, config.C.Auth.TTL)
	fmt.Println(err)
	if err != nil {
		return token, err
	}
	return token, err
}
