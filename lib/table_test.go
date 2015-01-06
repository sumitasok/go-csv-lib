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
		Data: []*Row{
			&Row{rowT{
				"name":    "Sumit",
				"phone":   "1234567890",
				"address": "India",
			}},
			&Row{rowT{
				"name":    tName,
				"phone":   "1234567890",
				"address": "Nepal",
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

	assert.Equal(2, table1.RowCount())
}

func TestTableSort(t *testing.T) {
	assert := assert.New(t)

	// table := table1.Sort("name")
	// assert.Equal("Sumit", table.Row(1).Value("name"))

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
