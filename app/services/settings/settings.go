package settings

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"maply/config"
	"maply/models"
	"maply/repository/managers"
	"maply/services/utils"
)

func Settings(userId string, s *models.Settings) error {
	u, err := managers.GetUser(userId)
	if err != nil {
		return err
	}

	if s.Name == "" {
		s.Name = u.Name
	}

	if s.Username == "" {
		s.Username = u.Username
	}

	if s.Avatar != nil && s.Avatar.Size != 0 {
		var avatar = fmt.Sprintf("%s.jpg", utils.HashFileName(s.Avatar.Filename))
		fasthttp.SaveMultipartFile(s.Avatar, config.C.App.BaseDir+config.C.App.MediaDir+avatar)
		s.Avatar.Filename = avatar
	}

	err = managers.UpdateUser(u.ID, s)
	if err != nil {
		return err
	}
	return err
}
