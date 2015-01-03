package lib

import (
	"encoding/csv"
	"fmt"
	// "io"
	// "io/ioutil"
	// "bufio"
	"os"
)

type header map[int]string
type row map[string]string
type tableData map[int]row

type Table struct {
	Header header
	Data   tableData
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

	rawCSVdata, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		return Table{}, err
	}

	h := make(header)
	t := make(tableData)

	for i, each := range rawCSVdata {
		if i == 0 {
			for j, e := range each {
				h[j] = e
			}
		} else {
			r := make(row)
			for j, e := range each {
				r[h[j]] = e
			}
			t[i] = r
		}
	}

	return Table{h, t}, nil
}
