# gotable

## Introduction
Print table in console

## reference
Please refer to guide: [gotable guide](https://blog.csdn.net/TCatTime/article/details/103068260#%E8%8E%B7%E5%8F%96gotable)

## API
### github.com/liushuochen/gotable
- Create table
```go
func CreateTable(header []string) (*table.Table, error)
```

- Get version
```go
func Version() string
```

### *table.Table
- Add value
```go
func (tb *Table) AddValue(newValue map[string]string) error
```

- Add a list of values
Method ```AddValues``` add a list of values. It returns a slice that 
consists of adding failed values.

```go
func (tb *Table) AddValues(values []map[string]string) []map[string]string
```

- Add head
```go
func (tb *Table) AddHead(newHead string) error
```

- Print table
```go
func (tb *Table) PrintTable()
```

- Set default value
<p>By default, the default value for all heads is an empty string.</p>

```go
func (tb *Table) SetDefault(h string, defaultValue string)
```

- Get default value
Use table method ```GetDefault``` to get default value of head. 
If h does not exist in the table.Header, the method returns an empty 
string.

```go
func (tb *Table) GetDefault(h string) string
```

- Arrange: center, align left or align right
<p> By default, the table is centered. You can set a header to be left 
aligned or right aligned. See the next section for more details on how 
to use it.</p>

```go
func (tb *Table) Align(head string, mode int)
```

- Check empty
Use table method ```Empty``` to check if the table is empty.

```go
func (tb *Table) Empty() bool
```

- Get list of head
Use table method ```GetHeaders``` to get a list of heads.

```go
func (tb *Table) GetHeaders() []string
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
	headers := []string{"China", "US", "UK"}
	tb, err := gotable.CreateTable(headers)
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}
}

```

### Add value
Use table method ```AddValue``` to add a new value in the table.

```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	headers := []string{"China", "US", "UK"}
	tb, err := gotable.CreateTable(headers)
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	value := make(map[string]string)
	value["China"] = "Beijing"
	value["US"] = "DC"
	value["UK"] = "London"
	err = tb.AddValue(value)
    if err != nil {
        fmt.Println("Add value to table failed: ", err.Error())
        return
    }
}

```

### Add values
Method ```AddValues``` add a list of values. It returns a slice that 
consists of adding failed values.

```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	headers := []string{"Name", "ID", "salary"}
	tb, err := gotable.CreateTable(headers)
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	values := make([]map[string]string, 0)
	for i := 0; i < 3; i++ {
		value := make(map[string]string)
		value["Name"] = fmt.Sprintf("employee-%d", i)
		value["ID"] = fmt.Sprintf("00%d", i)
		value["salary"] = "60000"
		values = append(values, value)
	}

	tb.AddValues(values)
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
	headers := []string{"China", "US", "UK"}
	tb, err := gotable.CreateTable(headers)
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	value := make(map[string]string)
	value["China"] = "Beijing"
	value["US"] = "Washington, D.C."
	value["UK"] = "London"
	tb.AddValue(value)

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
	headers := []string{"China", "US", "UK"}
	tb, err := gotable.CreateTable(headers)
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	tb.SetDefault("China", "Xi'AN")

	value := make(map[string]string)
	value["China"] = "Beijing"
	value["US"] = "Washington D.C."
	value["UK"] = "London"
	tb.AddValue(value)

	value2 := make(map[string]string)
	value2["US"] = "NewYork"
	value2["UK"] = "Manchester"
	tb.AddValue(value2)

	tb.PrintTable()
}

```

execute result:
```text
+---------+-----------------+------------+
|  China  |       US        |     UK     |
+---------+-----------------+------------+
| Beijing | Washington D.C. |   London   |
|  Xi'AN  |     NewYork     | Manchester |
+---------+-----------------+------------+

```

### Get default value
```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	headers := []string{"China", "US", "UK"}
	tb, err := gotable.CreateTable(headers)
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

### Add a new head to a table
```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	headers := []string{"China", "US", "UK"}
	tb, err := gotable.CreateTable(headers)
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	value := make(map[string]string)
	value["China"] = "Beijing"
	value["US"] = "Washington D.C."
	value["UK"] = "London"
	tb.AddValue(value)
	tb.AddHead("Japan")

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
	headers := []string{"China", "US", "UK"}
	tb, err := gotable.CreateTable(headers)
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	value := make(map[string]string)
	value["China"] = "Beijing"
	value["US"] = "Washington D.C."
	value["UK"] = "London"
	tb.AddValue(value)
	tb.Align("UK", gotable.Left)

	value2 := make(map[string]string)
	value2["US"] = "NewYork"
	value2["UK"] = "Manchester"
	tb.AddValue(value2)

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
	headers := []string{"China", "US", "UK"}
	tb, err := gotable.CreateTable(headers)
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	if tb.Empty() {
		fmt.Println("table is empty.")
	}
}
```

### Get list of heads
```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	headers := []string{"China", "US", "UK"}
	tb, err := gotable.CreateTable(headers)
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	fmt.Println(tb.GetHeaders())
	// [China US UK]
}
```
