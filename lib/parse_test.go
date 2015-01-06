package lib

import (
	// "fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseFile(t *testing.T) {
	assert := assert.New(t)

	table, err := ParseFile("/Users/SumitAsok/Downloads/google-contacts/other-contacts.csv")

	assert.Equal(len(table.Data[1]), len(table.Header)+1)

	// l := 0
	// for i := range table.Data[1] {
	// 	fmt.Println(l, i, "---", table.Data[1][i])
	// 	l += 1
	// }
	assert.Nil(err)
}

// Table.Data[index].WithHeader(string) = Table.Data[index]["header"] => "Value"
// Table.Data[index].WithHeaderIndex(int)
// Table.Matching(key, value) => []Row
// Table.Sort(key, ASC/DESC) => []Row
