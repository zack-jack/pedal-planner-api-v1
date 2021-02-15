package pedalboards

import (
	"context"

	"github.com/pkg/errors"
)

// FindAllPedalboards gets all pedalboards
func (s *Store) FindAllPedalboards(ctx context.Context) ([]Pedalboard, error) {
	row, err := s.stmts[qryFindAllPedalboards].QueryContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to fetch pedals")
	}

	var pedalboards []Pedalboard
	for row.Next() {
		var pedalboard Pedalboard
		err = row.Scan(
			&pedalboard.ID,
			&pedalboard.Brand,
			&pedalboard.Name,
			&pedalboard.Width,
			&pedalboard.Height,
			&pedalboard.Image,
		)
		if err != nil {
			return nil, errors.Wrap(err, "unable to scan pedals")
		}

		pedalboards = append(pedalboards, pedalboard)
	}

	return pedalboards, nil
}
