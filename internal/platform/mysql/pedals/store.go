package pedals

import (
	"database/sql"
	"fmt"

	"github.com/zack-jack/pedal-tetris-api-v1/internal/platform/mysql"
)

// Queries
const (
	qryFindAllPedals = "qry-all-pedals"
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
	tblPedals = "pedal_data.pedals"
)

// Store is MySQL storage for storing deal items
type Store struct {
	mysql.Store
	stmts map[string]*sql.Stmt
	db    *sql.DB
}

// NewPedalsStore initializes a mysql store with the underlying connection pool
func NewPedalsStore(db *sql.DB) (*Store, error) {
	unprepared := map[string]mysql.StmtMeta{
		qryFindAllPedals: mysql.NewStmtMeta(opSelect, tblPedals, `
		  SELECT
				id, brand, name, width, height, image
			FROM
				pedal_data.pedals
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
