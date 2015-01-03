package lib

import (
	"encoding/csv"
	"fmt"
	// "io"
	// "io/ioutil"
	// "bufio"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ParseFile(filePath string) {
	f, err := os.Open(filePath)
	check(err)

	defer f.Close()

	reader := csv.NewReader(f)

	reader.FieldsPerRecord = -1

	rawCSVdata, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, each := range rawCSVdata {
		fmt.Printf("email : %s and timestamp : %s\n", each[0], each[1])
	}
}
