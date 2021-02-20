package pedalboards

import (
	"database/sql"
	"fmt"

	"github.com/zack-jack/pedal-planner-api-v1/internal/platform/mysql"
)

// Queries
const (
	qryFindAllPedalboards = "qry-all-pedalboards"
)

// Database operation types
const (
	opSelect = "SELECT"
	opInsert = "INSERT"
	opUpdate = "UPDATE"
	opDelete = "DELETE"
)

// Table names
const (
	tblPedalboards = "pedalboard_db.pedalboards"
)

// Store is MySQL storage for storing pedalboard items
type Store struct {
	mysql.Store
	stmts map[string]*sql.Stmt
	db    *sql.DB
}

// NewPedalboardsStore initializes a mysql store with the underlying connection pool
func NewPedalboardsStore(db *sql.DB) (*Store, error) {
	unprepared := map[string]mysql.StmtMeta{
		qryFindAllPedalboards: mysql.NewStmtMeta(opSelect, tblPedalboards, `
		  SELECT
				id, brand, name, width, height, image
			FROM
				pedalboard_db.pedalboards
		`),
	}

	prepared, err := mysql.PrepareStmts(db, unprepared)
	if err != nil {
		fmt.Println("error prepared", err)
		return nil, err
	}

	s := &Store{
		db:    db,
		stmts: prepared,
	}
	s.StmtsMeta = unprepared

	return s, nil
}
