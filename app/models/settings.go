package models

import "mime/multipart"

type Settings struct {
	Name     string                `json:"name" validate:"omitempty,min=2,max=24"`
	Username string                `json:"username" validate:"omitempty,min=4,max=24"`
	Avatar   *multipart.FileHeader `json:"-" validate:"omitempty"`
}
