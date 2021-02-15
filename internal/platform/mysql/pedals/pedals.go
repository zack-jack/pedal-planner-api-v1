package pedals

import (
	"context"

	"github.com/pkg/errors"
)

// FindAllPedals gets all pedals
func (s *Store) FindAllPedals(ctx context.Context) ([]Pedal, error) {
	row, err := s.stmts[qryFindAllPedals].QueryContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to fetch pedals")
	}

	var pedals []Pedal
	for row.Next() {
		var pedal Pedal
		err = row.Scan(
			&pedal.ID,
			&pedal.Brand,
			&pedal.Name,
			&pedal.Width,
			&pedal.Height,
			&pedal.Image,
		)
		if err != nil {
			return nil, errors.Wrap(err, "unable to scan pedals")
		}

		pedals = append(pedals, pedal)
	}

	return pedals, nil
}
