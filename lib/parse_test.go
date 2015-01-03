package lib

import (
	// "fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseFile(t *testing.T) {
	assert := assert.New(t)

	ParseFile("/Users/SumitAsok/Downloads/google-contacts/other-contacts.csv")

	assert.True(true)
}
