package users

import (
	"github.com/ulule/deepcopier"
	"maply/models"
	"maply/repository/managers/users"
)

func GetUser(id string) (*models.PrivateUser, error) {
	u, err := users.GetUser(id)
	if err != nil {
		return &models.PrivateUser{}, err
	}

	resp := &models.PrivateUser{}
	for i := range u.Friends {
		resp.Friends = append(resp.Friends, &models.PublicUserWithoutFriends{})
		deepcopier.Copy(resp.Friends[i]).From(u.Friends[i])
	}
	deepcopier.Copy(resp).From(u)
	return resp, nil
}

func GetUserByID(id string) (*models.PublicUser, error) {
	u, err := users.GetUser(id)
	if err != nil {
		return &models.PublicUser{}, err
	}

	resp := &models.PublicUser{}
	for i := range u.Friends {
		resp.Friends = append(resp.Friends, &models.PublicUserWithoutFriends{})
		deepcopier.Copy(resp.Friends[i]).From(u.Friends[i])
	}
	deepcopier.Copy(resp).From(u)
	return resp, nil
}

func FindUser(username string) ([]*models.PublicUserWithoutFriends, error) {
	u, err := users.FindUser(username)
	if err != nil {
		return []*models.PublicUserWithoutFriends{}, err
	}

	var resp []*models.PublicUserWithoutFriends
	for i := range u {
		resp = append(resp, &models.PublicUserWithoutFriends{})
		deepcopier.Copy(resp[i]).From(u[i])
	}
	return resp, nil
}
