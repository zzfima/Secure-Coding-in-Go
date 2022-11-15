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

// CreateCountryTable : create Country table
func CreateCountryTable() error {
	db, e := sql.Open("sqlite3", "./cars.db")
	if e != nil {
		return e
	}
	defer db.Close()

	stmt, e := db.Prepare("CREATE TABLE IF NOT EXISTS Country (ID INTEGER PRIMARY KEY, Name TEXT)")
	if e != nil {
		return e
	}
	stmt.Exec()

	return nil
}

// InsertIntoCountryTable : create Country table
func InsertIntoCountryTable(id int, name string) error {
	db, e := sql.Open("sqlite3", "./cars.db")
	if e != nil {
		return e
	}
	defer db.Close()

	stmt, e := db.Prepare("INSERT INTO Country (ID, Name) VALUES (?, ?)")
	if e != nil {
		return e
	}
	stmt.Exec(id, name)

	return nil
}

// Country describes country entity
type Country struct {
	ID   int
	Name string
}

// ReadFromCountryTable : read countries from db
func ReadFromCountryTable() ([]Country, error) {
	countries := []Country{}
	db, e := sql.Open("sqlite3", "./cars.db")
	if e != nil {
		return nil, e
	}
	defer db.Close()

	rows, e := db.Query("SELECT * FROM Country")
	if e != nil {
		return nil, e
	}
	defer rows.Close()

	var id int
	var name string
	for rows.Next() {
		rows.Scan(&id, &name)
		countries = append(countries, Country{id, name})
	}

	return countries, nil
}
