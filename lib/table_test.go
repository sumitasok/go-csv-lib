package lib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	tName  = "Rabi"
	table1 = Table{
		Header: headerT{
			0: "name",
			1: "phone",
			2: "address",
		},
		Data: []Row{
			{rowT{
				"name":    "Sumit",
				"phone":   "1234567890",
				"address": "India",
			}},
			{rowT{
				"name":    tName,
				"phone":   "1234567890",
				"address": "Nepal",
			}},
			{rowT{
				"name":    "Andrea",
				"phone":   "82364192783",
				"address": "Canada",
			}},
		},
	}
)

func TestTableRow(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(tName, table1.Row(1).Value("name"))

	assert.True(true)
}

func TestTableRowCount(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(3, table1.RowCount())
}

func TestTableSort(t *testing.T) {
	assert := assert.New(t)

	_, err := table1.Sort("camel")
	assert.Error(err)
	assert.EqualError(err, errHeaderNotFound.Error())
	table, err2 := table1.Sort("address")
	assert.Equal("Canada", table.Row(0).Value("address"))
	assert.Equal("Andrea", table.Row(0).Value("name"))
	assert.NoError(err2)

	assert.True(true)
}

func TestTableValues(t *testing.T) {
	assert := assert.New(t)

	values := table1.Values("name")

	assert.Equal(3, len(values))
	assert.Equal("Andrea", values[0])
	assert.True(true)
}

func TestTableHeaders(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(3, len(table1.Headers()))
}

func TestContains(t *testing.T) {
	assert := assert.New(t)

	assert.True(Contains(table1.Headers(), "name"))
	assert.False(Contains(table1.Headers(), "camel"))

	assert.True(true)
}
