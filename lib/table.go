package lib

import (
	"errors"
)

var (
	errHeaderNotFound = errors.New("Header not found")
)

type Table struct {
	Header headerT
	Data   []*Row
}

func (table Table) Row(index int) Row {
	return *table.Data[index]
}

func (table Table) RowCount() int {
	return len(table.Data)
}

func (table Table) Headers() []string {
	headers := make([]string, 0)
	for _, header := range table.Header {
		headers = append(headers, header)
	}
	return headers
}

// func (t Table) Sort(header string) (Table, error) {
// 	// if
// }

type Row struct {
	Data rowT
}

func (row Row) Value(header string) string {
	return row.Data[header]
}

func Contains(array []string, element string) bool {
	for _, arrayElem := range array {
		if arrayElem == element {
			return true
		}
	}
	return false
}
