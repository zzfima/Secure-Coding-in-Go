package part2

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCountryDB(t *testing.T) {
	e := CreateCountryTable()
	require.Nil(t, e)
	//InsertIntoCountryTable(3, "UK")
	countries, e := ReadFromCountryTable()
	require.Nil(t, e)

	require.Equal(t, 3, countries[0].ID)
	require.Equal(t, "UK", countries[0].Name)
}

func TestCreateAndInsertLogDB(t *testing.T) {
	db, _ := OpenDB()
	CreateTables(db)
	defer db.Close()
	InsertLog(db, "today is go day", time.Now())
	//executed query:
	//"INSERT INTO logs (time, message) VALUES ('2022-11-17 06:24:40', 'Hallow!');"
}

func TestInsertInjectionLogDB(t *testing.T) {
	db, _ := OpenDB()
	defer db.Close()
	InsertLog(db, "Hacked!'); DROP TABLE logs; --", time.Now())
	//executed query:
	//"INSERT INTO logs (time, message) VALUES ('2022-11-17 06:26:00', 'Hacked!'); DROP TABLE logs; --');"
}

func TestInsertFixedLogDB(t *testing.T) {
	db, e := OpenDB()
	require.Nil(t, e)
	defer db.Close()
	e = InsertLogFixed(db, "Hacked!'); DROP TABLE logs; --", time.Now())
	require.Nil(t, e)
}

func TestRunServerTimeoutsDDOS(t *testing.T) {
	RunServerTimeoutsDDOS()
}
