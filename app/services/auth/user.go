package auth

import (
	"github.com/google/uuid"
	"maply/core/utils"
	"maply/models"
	"maply/repository/managers/auth"
	"time"
)

func CreateUser(u *models.User) (string, error) {
	var now = time.Now()
	u.ID = uuid.New().String()
	u.Password = utils.HashPassword(u.Password)
	u.CreatedAt = now
	u.UpdatedAt = now
	return auth.CreateUser(u)
}
