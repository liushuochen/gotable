# gotable

## Introduction
Print table in console

## Reference
Please refer to guide: [gotable guide](https://blog.csdn.net/TCatTime/article/details/103068260#%E8%8E%B7%E5%8F%96gotable)

## Supported character set
* ASCII
* Chinese characters

## API
### github.com/liushuochen/gotable
- Create table
```go
func Create(columns ...string) (*table.Table, error)
```

- Get version
```go
func Version() string
```

- Specify default value
The ```gotable.Default``` constant replaces the default value stored 
in head. Refer to Set Default Values in the Demo section for more 
information.

```go
gotable.Default
```

### *table.Table
- Add row
```go
func (tb *Table) AddRow(row map[string]string) error
```

- Add a list of rows
Method ```AddRows``` add a list of rows. It returns a slice that 
consists of adding failed rows.

```go
func (tb *Table) AddRows(rows []map[string]string) []map[string]string
```

- Add column
```go
func (tb *Table) AddColumn(column string) error
```

- Print table
```go
func (tb *Table) PrintTable()
```

- Set default value
By default, the default value for all heads is an empty string.

```go
func (tb *Table) SetDefault(h string, defaultValue string)
```

- Drop default value
```go
func (tb *Table) DropDefault(h string)
```

- Get default value
Use table method ```GetDefault``` to get default value of head. 
If h does not exist in the table.Header, the method returns an empty 
string.

```go
func (tb *Table) GetDefault(h string) string
```

- Get default map
Use table method ```GetDefaults``` to get default map of head. 

```go
func (tb *Table) GetDefaults() map[string]string
```

- Arrange: center, align left or align right
<p> By default, the table is centered. You can set a header to be left 
aligned or right aligned. See the next section for more details on how 
to use it.</p>

```go
func (tb *Table) Align(column string, mode int)
```

- Check empty
Use table method ```Empty``` to check if the table is empty.

```go
func (tb *Table) Empty() bool
```

- Get list of columns
Use table method ```GetColumns``` to get a list of columns.

```go
func (tb *Table) GetColumns() []string
```

- Get values map
Use table method ```GetValues``` to get the map that save values.

```go
func (tb *Table) GetValues() []map[string]string
```

- Check value exists
```go
func (tb *Table) Exist(value map[string]string) bool
```

- Get table length
```go
func (tb *Table) Length() int
```

- To json string
Use table method ```Json``` to convert the table to JSON format.
The argument ```indent``` indicates the number of indents.
If the argument ```indent``` is less than or equal to 0, then the ```Json``` method unindents.

```go
func (tb *Table) Json(indent int) (string, error)
```

- Close border
Use table method ```CloseBorder``` to close table border.

```go
func (tb *Table) CloseBorder()
```

- Open border
Use table method ```OpenBorder``` to open table border. By default, the border property is turned on.

```go
func (tb *Table) OpenBorder()
```


## Demo
### Create a table
Use ```gotable.CreateTable``` function with a header(```[]string```) to create a table.
It returns a pointer to a ```table.Table``` struct and an error.

```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	tb, err := gotable.Create("China", "US", "UK")
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}
}

```

### Add row
Use table method ```AddRow``` to add a new value in the table.

```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	tb, err := gotable.Create("China", "US", "UK")
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	row := make(map[string]string)
	row["China"] = "Beijing"
	row["US"] = "DC"
	row["UK"] = "London"
	err = tb.AddRow(row)
	if err != nil {
		fmt.Println("Add value to table failed: ", err.Error())
		return
	}
}

```

### Add rows
Method ```AddRows``` add a list of rows. It returns a slice that 
consists of adding failed rows.

```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	tb, err := gotable.Create("Name", "ID", "salary")
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	rows := make([]map[string]string, 0)
	for i := 0; i < 3; i++ {
		row := make(map[string]string)
		row["Name"] = fmt.Sprintf("employee-%d", i)
		row["ID"] = fmt.Sprintf("00%d", i)
		row["salary"] = "60000"
		rows = append(rows, row)
	}

	tb.AddRows(rows)
	tb.PrintTable()
}

```

execute result:
```text
+------------+-----+--------+
|    Name    | ID  | salary |
+------------+-----+--------+
| employee-0 | 000 | 60000  |
| employee-1 | 001 | 60000  |
| employee-2 | 002 | 60000  |
+------------+-----+--------+

```

### Print table
Table method ```PrintTable``` will print content of this table in STDOUT.

```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	tb, err := gotable.Create("China", "US", "UK")
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	row := make(map[string]string)
	row["China"] = "Beijing"
	row["US"] = "Washington, D.C."
	row["UK"] = "London"
	tb.AddRow(row)

	tb.PrintTable()
}

```

execute result:
```text
+---------+-----------------+--------+
|  China  |       US        |   UK   |
+---------+-----------------+--------+
| Beijing | Washington D.C. | London |
+---------+-----------------+--------+
```

### Set default value
```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	tb, err := gotable.Create("China", "US", "UK")
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	tb.SetDefault("China", "Xi'AN")
	tb.SetDefault("US", "Los Angeles")

	row := make(map[string]string)
	row["China"] = "Beijing"
	row["US"] = "Washington D.C."
	row["UK"] = "London"
	tb.AddRow(row)

	row2 := make(map[string]string)
	row2["US"] = "NewYork"
	row2["UK"] = "Manchester"
	tb.AddRow(row2)

	row3 := make(map[string]string)
	row3["China"] = "Hangzhou"
	// use gotable.Default
	row3["US"] = gotable.Default
	row3["UK"] = "Manchester"
	tb.AddRow(row3)

	tb.PrintTable()
}

```

execute result:
```text
+----------+-----------------+------------+
|  China   |       US        |     UK     |
+----------+-----------------+------------+
| Beijing  | Washington D.C. |   London   |
|  Xi'AN   |     NewYork     | Manchester |
| Hangzhou |   Los Angeles   | Manchester |
+----------+-----------------+------------+

```

### Drop default value
```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	tb, err := gotable.Create("China", "US", "UK")
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	tb.SetDefault("UK", "London")
	fmt.Println(tb.GetDefaults())
	// map[China: UK:London US:]
	tb.DropDefault("UK")
	fmt.Println(tb.GetDefaults())
	// map[China: UK: US:]
}

```

### Get default value
```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	tb, err := gotable.Create("China", "US", "UK")
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	tb.SetDefault("China", "Beijing")
	tb.SetDefault("China", "Hangzhou")
	fmt.Println(tb.GetDefault("China"))
	// Hangzhou
}
```

### Get default map
```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	tb, err := gotable.Create("China", "US", "UK")
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	tb.SetDefault("UK", "London")
	defaults := tb.GetDefaults()
	fmt.Println(defaults)
	// map[China: UK:London US:]
}

```

### Add a new column to a table
```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	tb, err := gotable.Create("China", "US", "UK")
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	row := make(map[string]string)
	row["China"] = "Beijing"
	row["US"] = "Washington D.C."
	row["UK"] = "London"
	tb.AddRow(row)
	tb.AddColumn("Japan")

	tb.PrintTable()
}

```

execute result:
```text
+---------+-----------------+--------+-------+
|  China  |       US        |   UK   | Japan |
+---------+-----------------+--------+-------+
| Beijing | Washington D.C. | London |       |
+---------+-----------------+--------+-------+

```

### Arrange: center, align left or align right
To change the arrangement, there are three constants
(```gotable.Center```, ```gotable.Left```, ```gotable.Right```) to 
choose from. By default, all arrangements is ```gotable.Center```.

```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	tb, err := gotable.Create("China", "US", "UK")
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	row := make(map[string]string)
	row["China"] = "Beijing"
	row["US"] = "Washington D.C."
	row["UK"] = "London"
	tb.AddRow(row)
	tb.Align("UK", gotable.Left)

	row2 := make(map[string]string)
	row2["US"] = "NewYork"
	row2["UK"] = "Manchester"
	tb.AddRow(row2)

	tb.PrintTable()
}

```

execute result:
```text
+---------+-----------------+------------+
|  China  |       US        |UK          |
+---------+-----------------+------------+
| Beijing | Washington D.C. |London      |
|         |     NewYork     |Manchester  |
+---------+-----------------+------------+

```

### Check empty
```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	tb, err := gotable.Create("China", "US", "UK")
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	if tb.Empty() {
		fmt.Println("table is empty.")
	}
}
```

### Get list of columns
```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	tb, err := gotable.Create("China", "US", "UK")
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	fmt.Println(tb.GetColumns())
	// [China US UK]
}
```

### Get values map
```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	tb, err := gotable.Create("China", "US", "UK")
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	tb.SetDefault("UK", "---")
	row := make(map[string]string)
	row["China"] = "Beijing"
	row["US"] = "Washington, D.C."
	row["UK"] = "London"
	_ = tb.AddRow(row)

	row2 := make(map[string]string)
	row2["China"] = "Hangzhou"
	row2["US"] = "NewYork"
	_ = tb.AddRow(row2)
	fmt.Println(tb.GetValues())
	// [map[China:Beijing UK:London US:Washington, D.C.] map[China:Hangzhou UK:--- US:NewYork]]
}

```

### Check value exists
```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	tb, err := gotable.Create("Name", "ID", "salary")
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	rows := make([]map[string]string, 0)
	for i := 0; i < 3; i++ {
		row := make(map[string]string)
		row["Name"] = fmt.Sprintf("employee-%d", i)
		row["ID"] = fmt.Sprintf("00%d", i)
		row["salary"] = "60000"
		rows = append(rows, row)
	}

	tb.AddRows(rows)

	row := make(map[string]string)
	row["salary"] = "60000"
	// check salary="60000" exists: true
	fmt.Println(tb.Exist(row))

	row["Name"] = "employee-5"
	// check salary="60000" && Name="employee-5" exists: false
	// The value of "employee-5" in Name do not exist
	fmt.Println(tb.Exist(row))

	row2 := make(map[string]string)
	row2["s"] = "60000"
	// check s="60000" exists: false
	// The table do not has a key named 's'
	fmt.Println(tb.Exist(row2))
}
```

### Get table length
```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	tb, err := gotable.Create("Name", "ID", "salary")
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	length := tb.Length()
	fmt.Printf("Before insert values, the value of length is: %d\n", length)
	// Before insert values, the value of length is: 0

	rows := make([]map[string]string, 0)
	for i := 0; i < 3; i++ {
		row := make(map[string]string)
		row["Name"] = fmt.Sprintf("employee-%d", i)
		row["ID"] = fmt.Sprintf("00%d", i)
		row["salary"] = "60000"
		rows = append(rows, row)
	}

	tb.AddRows(rows)

	length = tb.Length()
	fmt.Printf("After insert values, the value of length is: %d\n", length)
	// After insert values, the value of length is: 3
}

```

 ### To JSON string
```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	tb, err := gotable.Create("Name", "ID", "salary")
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	rows := make([]map[string]string, 0)
	for i := 0; i < 3; i++ {
		row := make(map[string]string)
		row["Name"] = fmt.Sprintf("employee-%d", i)
		row["ID"] = fmt.Sprintf("00%d", i)
		row["salary"] = "60000"
		rows = append(rows, row)
	}

	jsonString, err := tb.Json(4)
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
		return
	}
	fmt.Println(jsonString)
	// output: []

	tb.AddRows(rows)

	jsonString, err = tb.Json(4)
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
		return
	}
	fmt.Println(jsonString)
	// output:
	// [
	//       {
	//              "ID": "000",
	//              "Name": "employee-0",
	//              "salary": "60000"
	//       },
	//       {
	//              "ID": "001",
	//              "Name": "employee-1",
	//              "salary": "60000"
	//
	//
	//              "ID": "002",
	//              "Name": "employee-2",
	//              "salary": "60000"
	//       }
	//]
}
```

### Close border
```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	tb, err := gotable.Create("China", "US", "UK")
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	tb.SetDefault("US", "Los Angeles")

	row := make(map[string]string)
	row["China"] = "Beijing"
	row["US"] = "Washington D.C."
	row["UK"] = "London"
	tb.AddRow(row)

	row2 := make(map[string]string)
	row2["China"] = "Xi'AN"
	row2["US"] = "NewYork"
	row2["UK"] = "Manchester"
	tb.AddRow(row2)

	row3 := make(map[string]string)
	row3["China"] = "Hangzhou"
	row3["US"] = gotable.Default
	row3["UK"] = "Manchester"
	tb.AddRow(row3)

	// close border
	tb.CloseBorder()
	tb.Align("China", gotable.Left)
	tb.Align("US", gotable.Left)
	tb.Align("UK", gotable.Left)

	tb.PrintTable()
}
```

execute result:
```text
China     US               UK          
Beijing   Washington D.C.  London      
Xi'AN     NewYork          Manchester  
Hangzhou  Los Angeles      Manchester

```

### Open border
```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	tb, err := gotable.Create("China", "US", "UK")
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	tb.SetDefault("US", "Los Angeles")

	row := make(map[string]string)
	row["China"] = "Beijing"
	row["US"] = "Washington D.C."
	row["UK"] = "London"
	tb.AddRow(row)

	row2 := make(map[string]string)
	row2["China"] = "Xi'AN"
	row2["US"] = "NewYork"
	row2["UK"] = "Manchester"
	tb.AddRow(row2)

	row3 := make(map[string]string)
	row3["China"] = "Hangzhou"
	row3["US"] = gotable.Default
	row3["UK"] = "Manchester"
	tb.AddRow(row3)

	// close border
	tb.CloseBorder()

	// open border again
	tb.OpenBorder()

	tb.PrintTable()
}

```

execute result:
```text
+----------+-----------------+------------+
|  China   |       US        |     UK     |
+----------+-----------------+------------+
| Beijing  | Washington D.C. |   London   |
|  Xi'AN   |     NewYork     | Manchester |
| Hangzhou |   Los Angeles   | Manchester |
+----------+-----------------+------------+

```
