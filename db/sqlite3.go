package db

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

// SQLite3Client represents the wrapper of SQLite3 client
type SQLite3Client struct {
	db *sql.DB
}

// NewSQLite3Client creates new SQLite3Client object
func NewSQLite3Client(filename string) (*SQLite3Client, error) {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create sqlite3 client")
	}

	return &SQLite3Client{
		db: db,
	}, nil
}

// ImportPriceList imports price list to database
func (c *SQLite3Client) ImportPriceList(table string, header []string, records [][]string) error {
	headerValues := []string{}
	for _, h := range header {
		headerValues = append(headerValues, strconv.Quote(h))
	}

	queryValues := []string{}
	for range records[0] {
		queryValues = append(queryValues, "?")
	}

	// insert into aprice_price_list(h1, h2, ...) values(r1), (r2), ...;
	query := fmt.Sprintf("insert into %s(%s) values(%s);", table, strings.Join(headerValues, ", "), strings.Join(queryValues, ", "))

	for _, record := range records {
		vs := []interface{}{}

		for _, f := range record {
			vs = append(vs, f)
		}

		_, err := c.db.Exec(query, vs...)
		if err != nil {
			return errors.Wrap(err, "failed to insert record")
		}
	}

	return nil
}
