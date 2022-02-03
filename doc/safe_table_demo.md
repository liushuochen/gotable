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
