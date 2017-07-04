package db

// DB represents database client interface
type DB interface {
	CreateTable(table string, header []string) error
	ImportPriceList(table string, header []string, records [][]string) error
	TableExists(table string) (bool, error)
}
