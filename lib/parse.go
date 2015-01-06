package lib

import (
	"encoding/csv"
	"os"
)

type headerT map[int]string
type rowT map[string]string
type tableDataT map[int]rowT // kept as map type for make() and sort purpose

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

	// numberOfRows := len(rawCSVdata) - 1

	header := make(headerT)
	tableData := make([]Row, 0)

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
			tableData = append(tableData, Row{row})
		}
	}

	return Table{header, tableData}, nil
}
