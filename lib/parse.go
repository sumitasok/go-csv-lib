package lib

import (
	"encoding/csv"
	// "fmt"
	"os"
)

type headerT map[int]string
type rowT map[string]string
type tableDataT map[int]rowT

type Table struct {
	Header headerT
	Data   tableDataT
}

func (table Table) Row(index int) rowT {
	return table.Data[index]
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ParseFile(filePath string) (Table, error) {
	f, err := os.Open(filePath)
	check(err)

	defer f.Close()

	reader := csv.NewReader(f)

	reader.FieldsPerRecord = -1

	rawCSVdata, errReader := reader.ReadAll()

	if errReader != nil {
		return Table{}, errReader
	}

	header := make(headerT)
	tableData := make(tableDataT)

	for rowIndex, each := range rawCSVdata {
		if rowIndex == 0 { // means, first row is considered as headers always
			for columnIndex, headerName := range each {
				header[columnIndex] = headerName
			}
		} else {
			row := make(rowT)
			for columnIndex, columnValue := range each {
				headerName := header[columnIndex]
				row[headerName] = columnValue
			}
			tableData[rowIndex] = row
		}
	}

	return Table{header, tableData}, nil
}
