package models

type Settings struct {
	Name     string `json:"name" validate:"omitempty,min=2,max=24"`
	Username string `json:"username" validate:"omitempty,min=4,max=24"`
	Avatar   string `json:"avatar" validate:",min=4,max=24omitempty"`
}
