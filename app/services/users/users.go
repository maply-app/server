package users

import (
	"encoding/json"
	"github.com/ulule/deepcopier"
	"maply/cache/managers/stats"
	"maply/models"
	"maply/repository/managers/users"
	"maply/ws"
)

func setupUserInfo(userId string) (*models.UserInfo, error) {
	i := &models.UserInfo{}
	s, _ := setupUserStats(userId)
	i.Coords = s
	return i, nil
}

func setupUserStats(userId string) (*models.Stats, error) {
	s, err := stats.GetStats(userId)
	if err != nil {
		return nil, err
	}

	m := &models.Stats{}
	if err := json.Unmarshal([]byte(s.(string)), m); err != nil {
		return nil, err
	}

	// Update online status
	if ws.GetClientConnection(userId) != nil {
		m.IsOnline = true
	}

	return m, nil
}

func GetUser(userId string) (*models.PrivateUser, error) {
	u, err := users.GetUser(userId)
	if err != nil {
		return &models.PrivateUser{}, err
	}

	resp := &models.PrivateUser{}
	for i := range u.Friends {
		resp.Friends = append(resp.Friends, &models.Friend{})
		deepcopier.Copy(resp.Friends[i]).From(u.Friends[i])

		// Get user info
		info, _ := setupUserInfo(u.Friends[i].ID)
		resp.Friends[i].Info = info
	}
	deepcopier.Copy(resp).From(u)

	// Get user info
	i, _ := setupUserInfo(userId)
	resp.Info = i

	return resp, nil
}

func GetUserByID(selfId, userId string) (*models.PublicUser, error) {
	u, err := users.GetUser(userId)
	if err != nil {
		return &models.PublicUser{}, err
	}

	// Is it friends?
	isFriend := false
	for i := range u.Friends {
		if u.Friends[i].ID == selfId {
			isFriend = true
			break
		}
	}

	resp := &models.PublicUser{}
	for i := range u.Friends {
		resp.Friends = append(resp.Friends, &models.PublicUserWithoutFriends{})
		deepcopier.Copy(resp.Friends[i]).From(u.Friends[i])
	}
	deepcopier.Copy(resp).From(u)

	// Get user info
	if isFriend {
		i, _ := setupUserInfo(userId)
		resp.Info = i
	}

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
