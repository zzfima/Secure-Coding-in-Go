package part2

import (
	"database/sql"
	_ "embed"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var (
	//go:embed schema.sql
	schemaSQL string
	//go:embed insert.sql
	insertSQL string
)

// OpenDB open db
func OpenDB() (*sql.DB, error) {
	db, e := sql.Open("sqlite3", "./Logs.db")
	return db, e
}

// CreateTables create SQL table by using local schemaSQL
func CreateTables(db *sql.DB) error {
	_, e := db.Exec(schemaSQL)
	return e
}

// InsertLog insert log into db: message and time
func InsertLog(db *sql.DB, message string, time time.Time) error {
	ts := time.Format("2006-01-02 15:04:05")
	sqlQuery := fmt.Sprintf(insertSQL, ts, message)
	res, e := db.Exec(sqlQuery)
	fmt.Println(res)
	return e
}
