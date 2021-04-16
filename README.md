# gotable

## Introduction
Print table in console

## Effect
### Normal table
![](https://tuocheng.oss-cn-beijing.aliyuncs.com/gotable_test_plain.png)
### Color table
![](https://tuocheng.oss-cn-beijing.aliyuncs.com/gotable_test_color.png)

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

- Add head
```go
func (tb *Table) AddHead(newHead string) error
```

- Print table
```go
func (tb *Table) PrintTable()
```

- Set default value
```go
func (tb *Table) SetDefault(h string, defaultValue string)
```

## Demo
### Create a table
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
	tb.AddValue(value)

	tb.PrintTable()
}

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
	value["US"] = "DC"
	value["UK"] = "London"
	tb.AddValue(value)

	tb.PrintTable()
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
	value["US"] = "DC"
	value["UK"] = "London"
	tb.AddValue(value)
	tb.AddHead("Japan")

	tb.PrintTable()
}

```
