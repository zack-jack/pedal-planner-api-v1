package main

import (
	"database/sql"
	"time"

	"github.com/pkg/errors"
	"github.com/zack-jack/pedal-planner-api-v1/internal/pedalboards"
	"github.com/zack-jack/pedal-planner-api-v1/internal/pedals"
	pedalboardsmysql "github.com/zack-jack/pedal-planner-api-v1/internal/platform/mysql/pedalboards"
	pedalsmysql "github.com/zack-jack/pedal-planner-api-v1/internal/platform/mysql/pedals"
)

func setupPedalsStore(dsn string, maxOpenConnections int) (pedals.Store, error) {
	db, err := setupDB(dsn, maxOpenConnections)
	if err != nil {
		return nil, errors.Wrap(err, "set up db")
	}

	store, err := pedalsmysql.NewPedalsStore(db)
	if err != nil {
		return nil, errors.Wrap(err, "new sql store")
	}

	return store, nil
}

func setupPedalboardsStore(dsn string, maxOpenConnections int) (pedalboards.Store, error) {
	db, err := setupDB(dsn, maxOpenConnections)
	if err != nil {
		return nil, errors.Wrap(err, "set up db")
	}

	store, err := pedalboardsmysql.NewPedalboardsStore(db)
	if err != nil {
		return nil, errors.Wrap(err, "new sql store")
	}

	return store, nil
}

func setupDB(dsn string, maxOpenConnections int) (*sql.DB, error) {
	// the values for max open conns, max idle conns, and conn max lifetime are derived from here:
	// https://making.pusher.com/production-ready-connection-pooling-in-go/
	if maxOpenConnections < 0 {
		return nil, errors.New("invalid max open connections")
	}
	maxIdleConns := maxOpenConnections / 2
	connMaxLifetime := time.Minute * 30

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, errors.Wrap(err, "open sql db")
	}
	db.SetMaxOpenConns(maxOpenConnections)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxLifetime(connMaxLifetime)

	return db, nil
}
