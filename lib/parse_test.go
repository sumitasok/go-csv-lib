package lib

import (
	// "fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseFile(t *testing.T) {
	assert := assert.New(t)

	table, err := ParseFile("/Users/SumitAsok/Downloads/google-contacts/other-contacts.csv")

	assert.Equal(len(table.Data[1]), len(table.Header))
	assert.Nil(err)
}
