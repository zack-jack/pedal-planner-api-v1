package pedals

import (
	"context"

	store "github.com/zack-jack/pedal-planner-api-v1/internal/platform/mysql/pedals"
)

// Store is the data store interface
type Store interface {
	FindAllPedals(ctx context.Context) ([]store.Pedal, error)
}
