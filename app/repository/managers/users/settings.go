package users

import (
	"maply/models"
	"maply/repository"
)

func UpdateUser(userId string, s *models.Settings) error {
	query := "UPDATE users SET name = ?, username = ?, avatar = ? WHERE id = ?;"
	err := repository.DB.Exec(query, s.Name, s.Username, s.Avatar, userId).Error
	return err
}
