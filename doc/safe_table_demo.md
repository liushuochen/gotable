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
}

```
