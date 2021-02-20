package pedalboards

import (
	"context"

	"github.com/pkg/errors"
	store "github.com/zack-jack/pedal-planner-api-v1/internal/platform/mysql/pedalboards"
)

// Service is provides functionality for pedalboards
type Service struct {
	pedalboardsStore Store
}

// New creates a new pedalboards service instance
func New(pedalboardsStore Store) (*Service, error) {
	return &Service{
		pedalboardsStore: pedalboardsStore,
	}, nil
}

// PedalboardPublic is a pedal we return from the API
type PedalboardPublic struct {
	ID     int     `json:"id"`
	Brand  string  `json:"brand"`
	Name   string  `json:"name"`
	Width  float32 `json:"width"`
	Height float32 `json:"height"`
	Image  string  `json:"image"`
}

// FindAllPedalboards gets all pedals
func (s *Service) FindAllPedalboards(ctx context.Context) ([]PedalboardPublic, error) {
	pedals, err := s.pedalboardsStore.FindAllPedalboards(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to fetch pedalboards")
	}

	retPedalboards := getPublicPedalboards(pedals)

	return retPedalboards, nil
}

// Removing fields we don't need to return to client
func getPublicPedalboards(pedalboards []store.Pedalboard) []PedalboardPublic {
	publicPedalboards := make([]PedalboardPublic, len(pedalboards))
	for i, pedalboard := range pedalboards {
		publicPedalboards[i] = getPublicPedalboard(pedalboard)
	}
	return publicPedalboards
}

// Remove fields from pedalboard to return to client
func getPublicPedalboard(pedalboard store.Pedalboard) PedalboardPublic {
	return PedalboardPublic{
		ID:     pedalboard.ID,
		Brand:  pedalboard.Brand,
		Name:   pedalboard.Name,
		Width:  pedalboard.Width,
		Height: pedalboard.Height,
		Image:  pedalboard.Image,
	}
}
