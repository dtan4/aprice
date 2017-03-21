package database

import (
	"database/sql"

	_ "github.com/cznic/ql/driver" // Load cznic/ql driver
	"github.com/pkg/errors"
)

// New opens database file and return new database connection
func New(filename string) (*sql.DB, error) {
	conn, err := sql.Open("ql", filename)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open database file")
	}

	return conn, nil
}
