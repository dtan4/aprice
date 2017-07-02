package db

// DB represents database client interface
type DB interface {
	ImportPriceList(table string, header []string, records [][]string) error
}
