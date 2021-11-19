# API
This section describes the gotable APIs.

[Return to the home page](../README.md)

## github.com/liushuochen/gotable
### Create table
```go
func Create(columns ...string) (*table.Table, error)
```

### Create a table from struct
```go
func CreateByStruct(v interface{}) (*table.Table, error)
```

### Get version
```go
func Version() string
```

### Get version list
```go
func Versions() []string
```

### Specify default value
The ```gotable.Default``` constant replaces the default value stored in column. Refer to Set Default Values in the Demo 
section for more information.

```go
gotable.Default
```

### Load data from file
Currently，csv and json file are supported.
```go
func Read(path string) (*table.Table, error)
```

### Color control
The following constants are used in conjunction with the ```*table.SetColumnColor``` method to change the column color.
#### display type
Default display
```go
gotable.TerminalDefault
```
Fonts are highlighted
```go
gotable.Highlight
```
Underline
```go
gotable.Underline
```
Font flash
```go
gotable.Flash
```
#### color
```go
gotable.Black
gotable.Red
gotable.Green
gotable.Yellow
gotable.Blue
gotable.Purple
gotable.Cyan
gotable.Write
```

Do not set the background color
```go
gotable.NoneBackground
```

## *table.Table
### Clear data
The clear method is used to clear all data in the table, include columns and rows.
```go
func (tb *Table) Clear()
```

### Add row
Add a row to the table. Support Map and Slice. See the Demo section for more information.
```go
func (tb *Table) AddRow(row interface{}) error
```

### Add a list of rows
Method ```AddRows``` add a list of rows. It returns a slice that
consists of adding failed rows.

```go
func (tb *Table) AddRows(rows []map[string]string) []map[string]string
```

### Add column
```go
func (tb *Table) AddColumn(column string) error
```

### Print table
```*table``` implements ```fmt.Stringer``` interface, so you can use the ```fmt.Print```, ```fmt.Printf``` functions 
and so on to print the contents of the table instance.
```go
func (tb *Table) String() string
```

### Set default value
By default, the default value for all heads is an empty string.

```go
func (tb *Table) SetDefault(h string, defaultValue string)
```

### Drop default value
```go
func (tb *Table) DropDefault(h string)
```

### Get default value
Use table method ```GetDefault``` to get default value of head.
If h does not exist in the table.Header, the method returns an empty
string.

```go
func (tb *Table) GetDefault(h string) string
```

### Get default map
Use table method ```GetDefaults``` to get default map of head.

```go
func (tb *Table) GetDefaults() map[string]string
```

### Arrange: center, align left or align right
<p> By default, the table is centered. You can set a header to be left 
aligned or right aligned. See the next section for more details on how 
to use it.</p>

```go
func (tb *Table) Align(column string, mode int)
```

### Check empty
Use table method ```Empty``` to check if the table is empty.

```go
func (tb *Table) Empty() bool
```

### Get list of columns
Use table method ```GetColumns``` to get a list of columns.

```go
func (tb *Table) GetColumns() []string
```

### Get values map
Use table method ```GetValues``` to get the map that save values.
```go
func (tb *Table) GetValues() []map[string]string
```

### Check value exists
```go
func (tb *Table) Exist(value map[string]string) bool
```

### Get table length
```go
func (tb *Table) Length() int
```

### To JSON string
Use table method ```JSON``` to convert the table to JSON format.
The argument ```indent``` indicates the number of indents.
If the argument ```indent``` is less than or equal to 0, then the ```Json``` method unindents.
```go
func (tb *Table) JSON(indent int) (string, error)
```

### Save the table data to a JSON file
Use table method ```ToJsonFile``` to save the table data to a JSON file.
```go
func (tb *Table) ToJsonFile(path string, indent int) error
```

### Save the table data to a CSV file
Use table method ```ToCSVFile``` to save the table data to a CSV file.
```go
func (tb *Table) ToCSVFile(path string) error
```

### Close border
Use table method ```CloseBorder``` to close table border.
```go
func (tb *Table) CloseBorder()
```

### Open border
Use table method ```OpenBorder``` to open table border. By default, the border property is turned on.
```go
func (tb *Table) OpenBorder()
```

### Has column
Table method ```HasColumn``` determine whether the column is included.
```go
func (tb *Table) HasColumn(column string) bool
```

### Check whether the columns of the two tables are the same
Table method ```EqualColumns``` is used to check whether the columns of two tables are the same. This method returns
true if the columns of the two tables are identical (length, name, order, alignment, default), and false otherwise.
```go
func (tb *Table) EqualColumns(other *Table) bool
```

### Set column color
Table method ```SetColumnColor``` is used to set the color of a specific column. The first parameter specifies the name 
of the column to be modified. The second parameter indicates the type of font to display. Refer to the Color control 
section in this document for more information. The third and fourth parameters specify the font and background color.
```go
func (tb *Table) SetColumnColor(columnName string, display, fount, background int)
```
