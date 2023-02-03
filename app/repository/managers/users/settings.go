package users

import (
	"maply/models"
	"maply/repository"
)

func UpdateUser(userId string, s *models.Settings) error {
	if s.Avatar != nil {
		query := "UPDATE users SET name = ?, username = ?, avatar = ? WHERE id = ?;"
		err := repository.DB.Exec(query, s.Name, s.Username, s.Avatar.Filename, userId).Error
		return err
	} else {
		query := "UPDATE users SET name = ?, username = ? WHERE id = ?;"
		err := repository.DB.Exec(query, s.Name, s.Username, userId).Error
		return err
	}
}
