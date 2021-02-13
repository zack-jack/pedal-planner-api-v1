package data

import "time"

// Pedalboard represents a pedalboard
type Pedalboard struct {
	ID        int     `json:"id,omitempty"`
	Brand     string  `json:"brand"`
	Name      string  `json:"name"`
	Width     float32 `json:"width"`
	Height    float32 `json:"height"`
	Image     string  `json:"image"`
	UpdatedAt string  `json:"-"`
	CreatedAt string  `json:"-"`
}

type Pedalboards []*Pedalboard

func FindAllPedalboards() Pedalboards {
	return pedalboards
}

var pedalboards = Pedalboards{
	{
		ID:        1,
		Brand:     "Creation Music Co",
		Name:      "Elevation 24x12.5",
		Width:     24,
		Height:    12.5,
		Image:     "creation-elevation24125.png",
		UpdatedAt: time.Now().UTC().String(),
		CreatedAt: time.Now().UTC().String(),
	},
	{
		ID:        2,
		Brand:     "Pedaltrain",
		Name:      "Classic 1",
		Width:     22,
		Height:    12.5,
		Image:     "pedaltrain-classic1.png",
		UpdatedAt: time.Now().UTC().String(),
		CreatedAt: time.Now().UTC().String(),
	},
}
