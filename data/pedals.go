package data

import "time"

// Pedal represents a guitar pedal
type Pedal struct {
	ID        int
	Brand     string
	Name      string
	Width     float32
	Height    float32
	Image     string
	UpdatedAt string
	CreatedAt string
}

var pedalList = []*Pedal{
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
