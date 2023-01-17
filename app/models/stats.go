package models

type Cords struct {
	Lat float32 `json:"lat" validate:"min=-90,max=90"`
	Lon float32 `json:"lon" validate:"min=-180,max=180"`
}

type Geo struct {
	Cords     Cords `json:"cords" validate:"required"`
	Speed     int   `json:"speed" validate:"min=-1"`
	Direction int   `json:"direction" validate:"min=-360,max=360"`
}

type Info struct {
	Battery int `json:"battery" validate:"min=0,max=100"`
}

type Stats struct {
	Geo      Geo  `json:"geo" validate:"required"`
	Info     Info `json:"info" validate:"required"`
	IsOnline bool `json:"isOnline" validate:"omitempty"`
}