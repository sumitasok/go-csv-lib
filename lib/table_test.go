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
		Data: tableDataT{
			0: rowT{
				"name":    "Sumit",
				"phone":   "1234567890",
				"address": "India",
			},
			1: rowT{
				"name":    tName,
				"phone":   "1234567890",
				"address": "Nepal",
			},
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
