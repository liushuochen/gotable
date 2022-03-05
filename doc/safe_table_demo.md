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

Use safe table method ```AddRow``` to add a new row to the table. Currently, method ```AddRow``` only supports Map 
argument:
* For Map argument, you must put the data from each row into a Map and use column-data as key-value pairs. If the Map
  does not contain a column, the table sets it to the default value(see more information in 'Set Default' section). If
  the Map contains a column that does not exist, the ```AddRow``` method returns an error.
  
```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	table, err := gotable.CreateSafeTable("version", "description")
	if err != nil {
		fmt.Println("Create safe table failed: ", err.Error())
		return
	}

	row := make(map[string]string)
	row["version"] = "v1.0"
	row["description"] = "test"
	table.AddRow(row)
}

```
