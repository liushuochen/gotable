# Safe table demos
In this section, we have written some demos about safe table type for your reference.
Click [here](demo.md) to return to demo main page.



## Create a safe table

Use ```gotable.CreateSafeTable``` function with a column(string slice or strings) to create a safe table.
It returns a pointer of ```table.SafeTable``` struct and an error.

```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	table, err := gotable.CreateSafeTable("China", "US", "UK")
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}
}

```



## Print table

You can print the contents of the `SafeTable` instance to STDOUT using the print function in the ```fmt``` standard library.
For example ```fmt.Println```, ```fmt.Print``` and so on.

```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	table, _ := gotable.CreateSafeTable("ID", "Name")
	row := []string{"001", "gotable"}
	_ = table.AddRow(row)
	fmt.Println(table)
	// outputs:
	// +-----+---------+
	// | ID  |  Name   |
	// +-----+---------+
	// | 001 | gotable |
	// +-----+---------+
}

```



## Add row

Use safe table method ```AddRow``` to add a new row to the table. Method ```AddRow``` supports Map and Slice.
argument:
* For Map argument, you must put the data from each row into a Map and use column-data as key-value pairs. If the Map
  does not contain a column, the table sets it to the default value(see more information in 'Set Default' section). If
  the Map contains a column that does not exist, the ```AddRow``` method returns an error.
* For Slice argument, you must ensure that the slice length is equal to the column length. Method will automatically
  mapping values in Slice and columns. The default value cannot be omitted and must use gotable.Default constant.
  
```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	table, err := gotable.CreateSafeTable("China", "US", "French")
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	// Use map
	row := make(map[string]string)
	row["China"] = "Beijing"
	row["US"] = "Washington, D.C."
	row["French"] = "Paris"
	err = table.AddRow(row)
	if err != nil {
		fmt.Println("Add value to table failed: ", err.Error())
		return
	}

	// Use Slice
	row2 := []string{"Yinchuan", "Los Angeles", "Orleans"}
	err = table.AddRow(row2)
	if err != nil {
		fmt.Println("Add value to table failed: ", err.Error())
		return
	}

	fmt.Println(table)
	// outputs:
	// +----------+------------------+---------+
	// |  China   |        US        | French  |
	// +----------+------------------+---------+
	// | Beijing  | Washington, D.C. |  Paris  |
	// | Yinchuan |   Los Angeles    | Orleans |
	// +----------+------------------+---------+
}

```



## Add a new column to a table

```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	table, err := gotable.CreateSafeTable("China", "US", "French")
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	row := make(map[string]string)
	row["China"] = "Beijing"
	row["US"] = "Washington, D.C."
	row["French"] = "Paris"
	err = table.AddRow(row)
	if err != nil {
		fmt.Println("Add value to table failed: ", err.Error())
		return
	}

	err = table.AddColumn("Japan")
	if err != nil {
		fmt.Println("Add column failed: ", err.Error())
		return
	}

	fmt.Println(table)
	// outputs:
	// +---------+------------------+--------+-------+
	// |  China  |        US        | French | Japan |
	// +---------+------------------+--------+-------+
	// | Beijing | Washington, D.C. | Paris  |       |
	// +---------+------------------+--------+-------+
}


```



## Set default value

You can use the ```SetDefault``` method to set a default value for a column. By default, the default value is an empty
string. For Map structure data, when adding a row, omitting a column indicates that the value of column in the row is
the default value. You can also use the ```gotable.Default``` constant to indicate that a column in the row is the
default value. For Slice structure data, when adding a row, you must explicitly specify the ```gotable.Default```
constant to indicate that the value for a column is the default value.

```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	tb, err := gotable.CreateSafeTable("China", "US", "UK")
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
}

```



## Check the table type is safe table

```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	table, err := gotable.CreateSafeTable("China", "US", "UK")
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	fmt.Println(table.IsSafeTable())
	// output: true
}

```



## Get default value

```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	table, err := gotable.CreateSafeTable("China", "US", "UK")
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	table.SetDefault("China", "Beijing")
	table.SetDefault("China", "Hangzhou")
	fmt.Println(table.GetDefault("China"))
	// Hangzhou
}

```



## Get default map

```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	table, err := gotable.CreateSafeTable("China", "US", "UK")
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	table.SetDefault("UK", "London")
	fmt.Println(table.GetDefaults())
	// map[China: UK:London US:]
	table.DropDefault("UK")
	fmt.Println(table.GetDefaults())
	// map[China: UK: US:]
}

```



## Drop default value

```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	table, err := gotable.CreateSafeTable("China", "US", "UK")
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	table.SetDefault("UK", "London")
	fmt.Println(table.GetDefaults())
	// map[China: UK:London US:]
	table.DropDefault("UK")
	fmt.Println(table.GetDefaults())
	// map[China: UK: US:]
}
```



## Get table length

```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	tb, err := gotable.CreateSafeTable("Name", "ID", "salary")
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



## Check table empty

```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	table, err := gotable.CreateSafeTable("China", "US", "UK")
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	if table.Empty() {
		fmt.Println("table is empty.")
	}
}

```



## Get list of columns

```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	tb, err := gotable.CreateSafeTable("China", "US", "UK")
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	fmt.Println(tb.GetColumns())
	// [China US UK]
}

```

