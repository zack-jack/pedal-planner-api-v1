package data

import "time"

// Pedal represents a guitar pedal
type Pedal struct {
	ID        int     `json:"id,omitempty"`
	Brand     string  `json:"brand"`
	Name      string  `json:"name"`
	Width     float32 `json:"width"`
	Height    float32 `json:"height"`
	Image     string  `json:"image"`
	UpdatedAt string  `json:"-"`
	CreatedAt string  `json:"-"`
}

type Pedals []*Pedal

func FindAllPedals() Pedals {
	return pedals
}

var pedals = Pedals{
	{
		ID:        1,
		Brand:     "Strymon",
		Name:      "El Capistan",
		Width:     4,
		Height:    4.5,
		Image:     "strymon-elcap.png",
		UpdatedAt: time.Now().UTC().String(),
		CreatedAt: time.Now().UTC().String(),
	},
	{
		ID:        2,
		Brand:     "Strymon",
		Name:      "BigSky",
		Width:     6.75,
		Height:    5.1,
		Image:     "strymon-bigsky.png",
		UpdatedAt: time.Now().UTC().String(),
		CreatedAt: time.Now().UTC().String(),
	},
}
