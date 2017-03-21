package database

import (
	"database/sql"

	_ "github.com/cznic/ql/driver" // Load cznic/ql driver
	"github.com/pkg/errors"
)

// New opens database file and create DB object
func New(filename string) (*sql.DB, error) {
	db, err := sql.Open("ql", filename)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open database file")
	}

	return db, nil
}
