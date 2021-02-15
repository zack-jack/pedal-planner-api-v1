package mysql

import (
	"database/sql"
	"errors"
)

var (
	// ErrNotFound represents that a value is not found in the data store
	ErrNotFound = errors.New("not found")

	// ErrAlreadyExists represents that a value already exists in the data store
	ErrAlreadyExists = errors.New("already exists")
)

// Store is meant to be an embedded type that other mysql stores embed to use
// the data segment method
type Store struct {
	StmtsMeta map[string]StmtMeta
}

// StmtMeta represents metadata for a prepared statement
type StmtMeta struct {
	ParameterizedQuery string
	TableName          string
	Operation          string
}

// NewStmtMeta returns a StmtMeta value
func NewStmtMeta(operation, tableName, query string) StmtMeta {
	return StmtMeta{
		ParameterizedQuery: query,
		TableName:          tableName,
		Operation:          operation,
	}
}

// PrepareStmts will attempt to prepare each unprepared
// query on the database. If one fails, the function returns
// with an error.
func PrepareStmts(db *sql.DB, unprepared map[string]StmtMeta) (map[string]*sql.Stmt, error) {
	prepared := map[string]*sql.Stmt{}
	for k, v := range unprepared {
		stmt, err := db.Prepare(v.ParameterizedQuery)
		if err != nil {
			return nil, err
		}
		prepared[k] = stmt
	}

	return prepared, nil
}
