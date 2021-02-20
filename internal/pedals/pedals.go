package pedals

import (
	"context"

	"github.com/pkg/errors"
	store "github.com/zack-jack/pedal-planner-api-v1/internal/platform/mysql/pedals"
)

// Service is provides functionality for pedals
type Service struct {
	pedalsStore Store
}

// New creates a new pedals service instance
func New(pedalsStore Store) (*Service, error) {
	return &Service{
		pedalsStore: pedalsStore,
	}, nil
}

// PedalPublic is a pedal we return from the API
type PedalPublic struct {
	ID     int     `json:"id"`
	Brand  string  `json:"brand"`
	Name   string  `json:"name"`
	Width  float32 `json:"width"`
	Height float32 `json:"height"`
	Image  string  `json:"image"`
}

// FindAllPedals gets all pedals
func (s *Service) FindAllPedals(ctx context.Context) ([]PedalPublic, error) {
	pedals, err := s.pedalsStore.FindAllPedals(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to fetch pedals")
	}

	retPedals := getPublicPedals(pedals)

	return retPedals, nil
}

// Removing fields we don't need to return to client
func getPublicPedals(pedals []store.Pedal) []PedalPublic {
	publicPedals := make([]PedalPublic, len(pedals))
	for i, pedal := range pedals {
		publicPedals[i] = getPublicPedal(pedal)
	}
	return publicPedals
}

// Remove fields from pedal to return to client
func getPublicPedal(pedal store.Pedal) PedalPublic {
	return PedalPublic{
		ID:     pedal.ID,
		Brand:  pedal.Brand,
		Name:   pedal.Name,
		Width:  pedal.Width,
		Height: pedal.Height,
		Image:  pedal.Image,
	}
}
