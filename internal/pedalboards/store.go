package pedalboards

import (
	"context"

	store "github.com/zack-jack/pedal-tetris-api-v1/internal/platform/mysql/pedalboards"
)

// Store is the data store interface
type Store interface {
	FindAllPedalboards(ctx context.Context) ([]store.Pedalboard, error)
}
