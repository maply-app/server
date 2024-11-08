package users

import (
	"maply/models"
	"maply/repository"
)

func GetUser(id string) (models.User, error) {
	var user models.User
	err := repository.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}

	repository.DB.Model(&user).Association("Friends").Find(&user.Friends)
	return user, nil
}

func FindUser(username string) ([]models.User, error) {
	var users []models.User
	result := repository.DB.Limit(15).Where("username like ?", username+"%").Find(&users)
	return users, result.Error
}
