package managers

import (
	"maply/models"
	"maply/repository"
)

// UpdateUser ...
func UpdateUser(userID string, s *models.Settings) error {
	query := "UPDATE users SET name = ?, username = ?, avatar = ? WHERE id = ?;"
	err := repository.DB.Exec(query, s.Name, s.Username, s.Avatar.Filename, userID).Error
	return err
}
