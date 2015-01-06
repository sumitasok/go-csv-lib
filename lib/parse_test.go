package lib

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	inputCsvFile = "/Users/SumitAsok/Downloads/google-contacts/other-contacts.csv"
)

func TestParseFile(t *testing.T) {
	assert := assert.New(t)

	table, err := ParseFile(inputCsvFile)

	assert.Equal(len(table.Row(1).Data), len(table.Header)+1) // doesn't make sense
	// printTableRowData(table, 1)

	assert.NotNil(table.Row(1))
	assert.IsType(Row{}, table.Row(1))
	assert.Nil(err)
}

// Table.Data[index].WithHeader(string) = Table.Data[index]["header"] => "Value"
// Table.Data[index].WithHeaderIndex(int)
// Table.Matching(key, value) => []Row

// Table.Sort(key, ASC/DESC) => []Row

func printTableRowData(table Table, rowIndex int) {
	// use this to print the data
	iterator := 0

	for column := range table.Row(rowIndex).Data {
		fmt.Println(iterator, column, "---", table.Row(rowIndex).Value(column))
		iterator += 1
	}
}
