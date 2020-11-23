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

	value := make(map[string]gotable.Sequence)
	value["US"] = gotable.DefaultSequence("DC")
	value["UK"] = gotable.DefaultSequence("London")
	value["China"] = gotable.DefaultSequence("Beijing")
	tb.AddValue(value)

	tb.PrintTable()
}
```

### Create a color table
```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
	"github.com/liushuochen/gotable/color"
	"reflect"
	"strings"
)

type MyStruct struct {
	Name       string
	Experience int32
	Salary     float64
}

func main() {
	con := func (field string, val reflect.Value) color.Color {
		switch field {
		case "China":
			return color.RED
		case "US":
			return color.BLUE
		default:
			return ""
		}
	}

	headers := []string{"China", "US", "UK"}
	tb, err := gotable.CreateTable(headers, gotable.WithColorController(con))
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	value := make(map[string]gotable.Sequence)
	value["China"] = gotable.DefaultSequence("Beijing")
	value["US"] = gotable.DefaultSequence("DC")
	value["UK"] = gotable.DefaultSequence("London")
	tb.AddValue(value)

	tb.PrintTable()

	controller := func(field string, val reflect.Value) color.Color {
		switch field {
		case "Name":
			if strings.Contains(val.String(), "c") {
				return color.CYAN
			}
		case "Experience":
			if val.Int() > 5 {
				return color.MAGENTA
			}
		case "Salary":
			if val.Float() < 1000 {
				return color.YELLOW
			}
		}
		return ""
	}

	tbl, _ := gotable.CreateTableFromStruct(
		MyStruct{},
		gotable.WithColorController(controller),
		)

	structSliceData := []interface{}{
		MyStruct{
			Name:       "Mike",
			Experience: 3,
			Salary:     2300.00,
		},
		MyStruct{
			Name:       "Sum",
			Experience: 10,
			Salary:     900.00,
		},
		MyStruct{
			Name:       "Bob",
			Experience: 1,
			Salary:     9000.00,
		},
	}
	tbl.AddValuesFromSlice(structSliceData)

	tbl.PrintTable()
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

	value := make(map[string]gotable.Sequence)
	value["US"] = gotable.DefaultSequence("DC")
	value["UK"] = gotable.DefaultSequence("London")
	tb.AddValue(value)

	tb.PrintTable()
}
```

## 存在问题
- 单元格为中文时 表格对不齐
![](https://tuocheng.oss-cn-beijing.aliyuncs.com/gotable_chi_issue.png)
