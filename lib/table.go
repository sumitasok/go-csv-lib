package lib

type Table struct {
	Header headerT
	Data   tableDataT
}

func (table Table) Row(index int) Row {
	return Row{table.Data[index]}
}

func (table Table) RowCount() int {
	return len(table.Data)
}

type Row struct {
	Data rowT
}

func (row Row) Value(header string) string {
	return row.Data[header]
}
