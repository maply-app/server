package settings

import (
	"maply/models"
	"maply/repository/managers/users"
)

func Settings(userId string, s *models.Settings) error {
	u, err := users.GetUser(userId)
	if err != nil {
		return err
	}

	if s.Name == "" {
		s.Name = u.Name
	}

	if s.Username == "" {
		s.Username = u.Username
	}

	if s.Username == "" {
		s.Avatar = u.Avatar
	}

	err = users.UpdateUser(u.ID, s)
	if err != nil {
		return err
	}
	return err
}
