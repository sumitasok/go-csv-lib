This library is ment to read the csv file usingthe go package "encoding/csv" and represent it in menmory in a table structure. The library also provides a way to manipulate the data in many ways, like get matching rows, sort rows.

> The library assumes that the first row of the csv file is always the headers

> And all data are string

How to read a csv and get the table

```
table, err := lib.Parse("/user/you/document/your-file.csv")
```

||name|phone|
|---|---|---|
|0|sumit|1234567890|
|1|rabi|1234567890|
|2|andrea|1234567890|

To get row count

```
table.RowCount() => 3
```

To iterate over rows

```
for i, e := range table.Rows {
}
```

you can also query row with index
index start from 0..table.RowCount() 

```
table.Row(0)
```
||name|phone|
|---|---|---|
|0|sumit|1234567890|


To get the list of headers

```
table.Headers
```

```
 => ["name", "phone", "address"]
```

To get the values as an array in the header

```
table.Values("name")
```

```
=> ["sumit", "rabi", "andrea"]
```

To sort the table rows by a header value

```
table := table.Sort("name")
```

||name|phone|
|---|---|---|
|0|andrea|1234567890|
|1|rabi|1234567890|
|2|sumit|1234567890|