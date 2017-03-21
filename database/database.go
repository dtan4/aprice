package database

import (
	"database/sql"

	_ "github.com/cznic/ql/driver" // Load cznic/ql driver
	"github.com/pkg/errors"
)

// DB represents abstract database object
type DB struct {
	conn *sql.DB
}

// New opens database file and return new database connection
func New(filename string) (*DB, error) {
	conn, err := sql.Open("ql", filename)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open database file")
	}

	return &DB{
		conn: conn,
	}, nil
}

// Close closes database connection
func (d *DB) Close() error {
	return d.conn.Close()
}
