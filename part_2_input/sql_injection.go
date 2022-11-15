package part2

import (
	"database/sql"
	"fmt"
	"time"

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

// CreateDB create db
func CreateDB() (*sql.DB, error) {
	db, e := sql.Open("sqlite3", "./Logs.db")
	if e != nil {
		return nil, e
	}

	stmt, e := db.Prepare("CREATE TABLE IF NOT EXISTS Log (ID INTEGER PRIMARY KEY, Message TEXT, Time TIME)")
	if e != nil {
		return nil, e
	}
	stmt.Exec()

	return db, nil
}

// InsertLog insert log into db: message and time
func InsertLog(db *sql.DB, message string, time time.Time) error {
	stmt, e := db.Prepare("INSERT INTO Log (Message, Time) VALUES (?, ?)")
	if e != nil {
		return e
	}
	stmt.Exec(message, time)

	return nil
}
