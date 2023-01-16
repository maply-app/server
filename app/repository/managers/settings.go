package managers

import (
	"maply/models"
	"maply/repository"
)

// UpdateUser ...
func UpdateUser(userID string, s *models.Settings) error {
	if s.Avatar != nil {
		query := "UPDATE users SET name = ?, username = ?, avatar = ? WHERE id = ?;"
		err := repository.DB.Exec(query, s.Name, s.Username, s.Avatar.Filename, userID).Error
		return err
	} else {
		query := "UPDATE users SET name = ?, username = ? WHERE id = ?;"
		err := repository.DB.Exec(query, s.Name, s.Username, userID).Error
		return err
	}
}
