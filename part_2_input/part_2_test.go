package part2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestXxx(t *testing.T) {
	e := CreateCountryTable()
	require.Nil(t, e)
	//InsertIntoCountryTable(3, "UK")
	countries, e := ReadFromCountryTable()
	require.Nil(t, e)

	require.Equal(t, 1, countries[0].ID)
	require.Equal(t, "Israel", countries[0].Name)

	require.Equal(t, 2, countries[1].ID)
	require.Equal(t, "CZ", countries[1].Name)

	require.Equal(t, 3, countries[2].ID)
	require.Equal(t, "UK", countries[2].Name)
}
