package users

import (
	"github.com/ulule/deepcopier"
	"maply/models"
	"maply/repository/managers"
)

// GetUser ...
func GetUser(id string) (*models.PrivateUser, error) {
	u, err := managers.GetUser(id)
	if err != nil {
		return &models.PrivateUser{}, err
	}

	// Pretty response
	resp := &models.PrivateUser{}
	for i := range u.Friends {
		resp.Friends = append(resp.Friends, &models.PublicUserWithoutFriends{})
		deepcopier.Copy(resp.Friends[i]).From(u.Friends[i])
	}
	deepcopier.Copy(resp).From(u)
	return resp, nil
}

// GetUserByID ...
func GetUserByID(id string) (*models.PublicUser, error) {
	u, err := managers.GetUser(id)
	if err != nil {
		return &models.PublicUser{}, err
	}

	// Pretty response
	resp := &models.PublicUser{}
	for i := range u.Friends {
		resp.Friends = append(resp.Friends, &models.PublicUserWithoutFriends{})
		deepcopier.Copy(resp.Friends[i]).From(u.Friends[i])
	}
	deepcopier.Copy(resp).From(u)
	return resp, nil
}

// FindUser ...
func FindUser(username string) ([]*models.PublicUserWithoutFriends, error) {
	u, err := managers.FindUser(username)
	if err != nil {
		return []*models.PublicUserWithoutFriends{}, err
	}

	// Pretty response
	var resp []*models.PublicUserWithoutFriends
	for i := range u {
		resp = append(resp, &models.PublicUserWithoutFriends{})
		deepcopier.Copy(resp[i]).From(u[i])
	}
	return resp, nil
}
