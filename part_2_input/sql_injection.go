package part2

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var (
	schemaSQL string
	insertSQL string
)

// CreateTables create SQL table by using local schemaSQL
func CreateTables(db *sql.DB) error {
	r, e := db.Exec(schemaSQL)
	fmt.Println(r)
	return e
}
