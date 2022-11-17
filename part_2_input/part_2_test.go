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

func TestLogDB(t *testing.T) {
	db, _ := OpenDB()
	CreateTables(db)
	defer db.Close()
	InsertLog(db, "today is go day", time.Now())
}
