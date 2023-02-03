package models

import (
	"time"
)

// Database structures

type User struct {
	ID        string    `gorm:"type:uuid;primary_key" json:"id"`
	Name      string    `gorm:"not null;size:24;" json:"name" validate:"required,min=2,max=24"`
	Username  string    `gorm:"not null;size:24;unique" json:"username" validate:"required,min=4,max=24"`
	Email     string    `gorm:"not null; type:varchar(100);unique" json:"email" validate:"required,email,min=6,max=32"`
	Avatar    string    `gorm:"default:null; type:varchar(256);" json:"avatar" validate:"omitempty"`
	Password  string    `gorm:"size:255;" json:"password" validate:"required,min=8,max=24"`
	Friends   []*User   `gorm:"many2many:user_friends;" json:"friends"`
	IsActive  bool      `gorm:"not null;default:true" json:"isActive"`
	IsAdmin   bool      `gorm:"not null;default:false" json:"isAdmin"`
	CreatedAt time.Time `gorm:"autoCreateTime:false" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoCreateTime:false" json:"updatedAt"`
}

// Output structures

type UserInfo struct {
	Coords *Stats `json:"coords"`
}

type PrivateUser struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Avatar   string    `json:"avatar"`
	Friends  []*Friend `json:"friends"`
	Info     *UserInfo `json:"info"`
	IsAdmin  bool      `json:"isAdmin"`
}

type PublicUser struct {
	ID       string                      `json:"id"`
	Name     string                      `json:"name"`
	Username string                      `json:"username"`
	Avatar   string                      `json:"avatar"`
	Friends  []*PublicUserWithoutFriends `json:"friends"`
	IsAdmin  bool                        `json:"isAdmin"`
}

type PublicUserWithoutFriends struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	IsAdmin  bool   `json:"isAdmin"`
}

type Friend struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
	Avatar   string    `json:"avatar"`
	Info     *UserInfo `json:"info"`
	IsAdmin  bool      `json:"isAdmin"`
}
