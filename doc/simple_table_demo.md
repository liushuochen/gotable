# Simple table demos
In this section, we have written some demos about simple table type for your reference.
Click [here](demo.md) to return to demo main page.



## Create a table
Use ```gotable.Create``` function with a column(string slice or strings) to create a table.
It returns a pointer of ```table.Table``` struct and an error.

```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	table, err := gotable.Create("China", "US", "UK")
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}
}

```



## Create a table from struct

```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

type Student struct {
	Id		string	`gotable:"id"`    // Specify the column name of the table through the struct tag `gotable`
	Name	string
}

func main() {
	table, err := gotable.CreateByStruct(&Student{})
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}
}

```



## Load data from file

Currently, csv and json file are supported.
```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	table, err := gotable.Read("cmd/demo.csv")
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

}

```



## Add row

Use table method ```AddRow``` to add a new row to the table. Method ```AddRow``` supports Map and Slice argument:
* For Map argument, you must put the data from each row into a Map and use column-data as key-value pairs. If the Map
  does not contain a column, the table sets it to the default value(see more information in 'Set Default' section). If
  the Map contains a column that does not exist, the ```AddRow``` method returns an error.

* For Slice argument, you must ensure that the slice length is equal to the column length. The ```AddRow``` method
  automatically mapping values in Slice and columns. The default value cannot be omitted and must use
  ```gotable.Default``` constant.

```go
package main

import (
  "fmt"
  "github.com/liushuochen/gotable"
)

func main() {
  table, err := gotable.Create("China", "US", "French")
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



## Add rows

Method ```AddRows``` add a list of rows. It returns a slice that
consists of adding failed rows.

```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	table, err := gotable.Create("Name", "ID", "salary")
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

	table.AddRows(rows)
	fmt.Println(table)
	// outputs:
	// +------------+-----+--------+
	// |    Name    | ID  | salary |
	// +------------+-----+--------+
	// | employee-0 | 000 | 60000  |
	// | employee-1 | 001 | 60000  |
	// | employee-2 | 002 | 60000  |
	// +------------+-----+--------+
}
```



## Print table

You can print the contents of the table instance to STDOUT using the print function in the ```fmt``` standard library.
For example ```fmt.Println```, ```fmt.Print``` and so on.

```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	table, err := gotable.Create("China", "US", "UK")
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	row := make(map[string]string)
	row["China"] = "Beijing"
	row["US"] = "Washington, D.C."
	row["UK"] = "London"
	table.AddRow(row)

	fmt.Println(table)
	// outputs:
	// +---------+------------------+--------+
	// |  China  |        US        |   UK   |
	// +---------+------------------+--------+
	// | Beijing | Washington, D.C. | London |
	// +---------+------------------+--------+

	fmt.Printf("%s", table)
	// outputs:
	// +---------+------------------+--------+
	// |  China  |        US        |   UK   |
	// +---------+------------------+--------+
	// | Beijing | Washington, D.C. | London |
	// +---------+------------------+--------+

	fmt.Print(table)
	// outputs:
	// +---------+------------------+--------+
	// |  China  |        US        |   UK   |
	// +---------+------------------+--------+
	// | Beijing | Washington, D.C. | London |
	// +---------+------------------+--------+
}

```



## Clear data

```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	table, err := gotable.Create("Name", "ID", "salary")
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	length := table.Length()
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

	table.AddRows(rows)
	fmt.Println(table)
	// outputs:
	// +------------+-----+--------+
	// |    Name    | ID  | salary |
	// +------------+-----+--------+
	// | employee-0 | 000 | 60000  |
	// | employee-1 | 001 | 60000  |
	// | employee-2 | 002 | 60000  |
	// +------------+-----+--------+

	table.Clear()
	fmt.Println("After the table data is cleared...")
	fmt.Println(table)
	// output: After the table data is cleared...
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
	table, err := gotable.Create("China", "US", "UK")
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	table.SetDefault("China", "Xi'AN")
	table.SetDefault("US", "Los Angeles")

	row := make(map[string]string)
	row["China"] = "Beijing"
	row["US"] = "Washington D.C."
	row["UK"] = "London"
	table.AddRow(row)

	// China is omitted in Map, the value of China column in the row is changed to the default value(Xi'AN).
	row2 := make(map[string]string)
	row2["US"] = "NewYork"
	row2["UK"] = "Manchester"
	table.AddRow(row2)

	// Use the gotable.Default constant to indicate that the value of US is the default(Los Angeles)
	row3 := make(map[string]string)
	row3["China"] = "Hangzhou"
	row3["US"] = gotable.Default
	row3["UK"] = "Manchester"
	table.AddRow(row3)

	// Use gotable.Default in Slice.
	// Because the value of row4[1] is gotable.Default constant, the value for column[1](US) is the default value(Los Angeles)
	row4 := []string{"Qingdao", gotable.Default, "Oxford"}
	table.AddRow(row4)

	fmt.Println(table)
	// outputs:
	// +----------+-----------------+------------+
	// |  China   |       US        |     UK     |
	// +----------+-----------------+------------+
	// | Beijing  | Washington D.C. |   London   |
	// |  Xi'AN   |     NewYork     | Manchester |
	// | Hangzhou |   Los Angeles   | Manchester |
	// | Qingdao  |   Los Angeles   |   Oxford   |
	// +----------+-----------------+------------+
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
	table, err := gotable.Create("China", "US", "UK")
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



## Get default value

```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	table, err := gotable.Create("China", "US", "UK")
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



## Add a new column to a table

```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	table, err := gotable.Create("China", "US", "UK")
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	row := make(map[string]string)
	row["China"] = "Beijing"
	row["US"] = "Washington D.C."
	row["UK"] = "London"
	table.AddRow(row)
	table.AddColumn("Japan")

	fmt.Println(table)
	// outputs:
	// +---------+-----------------+--------+-------+
	// |  China  |       US        |   UK   | Japan |
	// +---------+-----------------+--------+-------+
	// | Beijing | Washington D.C. | London |       |
	// +---------+-----------------+--------+-------+
}

```



## Arrange: center, align left or align right

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
	table, err := gotable.Create("China", "US", "UK")
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	row := make(map[string]string)
	row["China"] = "Beijing"
	row["US"] = "Washington D.C."
	row["UK"] = "London"
	table.AddRow(row)
	table.Align("UK", gotable.Left)

	row2 := make(map[string]string)
	row2["US"] = "NewYork"
	row2["UK"] = "Manchester"
	table.AddRow(row2)

	fmt.Println(table)
	// outputs:
	// +---------+-----------------+------------+
	// |  China  |       US        |UK          |
	// +---------+-----------------+------------+
	// | Beijing | Washington D.C. |London      |
	// |         |     NewYork     |Manchester  |
	// +---------+-----------------+------------+
}

```



## Check empty

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



## Get list of columns

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



## Get values map

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



## Check value exists

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



## Get table length

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



## To JSON string

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

	jsonString, err := tb.JSON(4)
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
		return
	}
	fmt.Println(jsonString)
	// output: []

	tb.AddRows(rows)

	jsonString, err = tb.JSON(4)
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



## To XML string

```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	table, err := gotable.Create("Name", "ID", "salary")
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

	table.AddRows(rows)

	content := table.XML(2)
	fmt.Println(content)
	// outputs:
	// <?xml version="1.0" encoding="utf-8" standalone="yes"?>
	// <table>
	//   <row>
	//     <Name>employee-0</Name>
	//     <ID>000</ID>
	//     <salary>60000</salary>
	//   </row>
	//   <row>
	//     <Name>employee-1</Name>
	//     <ID>001</ID>
	//     <salary>60000</salary>
	//   </row>
	//   <row>
	//     <Name>employee-2</Name>
	//     <ID>002</ID>
	//     <salary>60000</salary>
	//   </row>
	// </table>
}

```



## Save the table data to a JSON file

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
	err = tb.ToJsonFile("cmd/demo.json", 4)
	if err != nil {
		fmt.Println("write json file error: ", err.Error())
		return
	}
}
```

cmd/demo.json:
```json
[
       {
              "ID": "000",
              "Name": "employee-0",
              "salary": "60000"
       },
       {
              "ID": "001",
              "Name": "employee-1",
              "salary": "60000"
       },
       {
              "ID": "002",
              "Name": "employee-2",
              "salary": "60000"
       }
]
```



## Save the table data to a CSV file

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
	err = tb.ToCSVFile("cmd/demo.csv")
	if err != nil {
		fmt.Println("write csv file error: ", err.Error())
		return
	}
}

```



## Close border

```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	table, err := gotable.Create("China", "US", "UK")
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	table.SetDefault("US", "Los Angeles")

	row := make(map[string]string)
	row["China"] = "Beijing"
	row["US"] = "Washington D.C."
	row["UK"] = "London"
	table.AddRow(row)

	row2 := make(map[string]string)
	row2["China"] = "Xi'AN"
	row2["US"] = "NewYork"
	row2["UK"] = "Manchester"
	table.AddRow(row2)

	row3 := make(map[string]string)
	row3["China"] = "Hangzhou"
	row3["US"] = gotable.Default
	row3["UK"] = "Manchester"
	table.AddRow(row3)

	// close border
	table.CloseBorder()
	table.Align("China", gotable.Left)
	table.Align("US", gotable.Left)
	table.Align("UK", gotable.Left)

	fmt.Println(table)
	// outputs:
	//  China    US              UK             
	//  Beijing  Washington D.C. London     
	//  Xi'AN    NewYork         Manchester 
	//  Hangzhou Los Angeles     Manchester 
}

```



## Open border

```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	table, err := gotable.Create("China", "US", "UK")
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	table.SetDefault("US", "Los Angeles")

	row := make(map[string]string)
	row["China"] = "Beijing"
	row["US"] = "Washington D.C."
	row["UK"] = "London"
	table.AddRow(row)

	row2 := make(map[string]string)
	row2["China"] = "Xi'AN"
	row2["US"] = "NewYork"
	row2["UK"] = "Manchester"
	table.AddRow(row2)

	row3 := make(map[string]string)
	row3["China"] = "Hangzhou"
	row3["US"] = gotable.Default
	row3["UK"] = "Manchester"
	table.AddRow(row3)

	// close border
	table.CloseBorder()

	// open border again
	table.OpenBorder()

	fmt.Println(table)
	// outputs:
	// +----------+-----------------+------------+
	// |  China   |       US        |     UK     |
	// +----------+-----------------+------------+
	// | Beijing  | Washington D.C. |   London   |
	// |  Xi'AN   |     NewYork     | Manchester |
	// | Hangzhou |   Los Angeles   | Manchester |
	// +----------+-----------------+------------+
}

```



## Has column

```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

type Stu struct {
	Name	string	`gotable:"name"`
	Sex		string	`gotable:"sex"`
	Age		int		`gotable:"age"`
}

func main() {
	tb, err := gotable.CreateByStruct(&Stu{})
	if err != nil {
		fmt.Println("[ERROR] ", err.Error())
		return
	}


	if tb.HasColumn("age") {
		fmt.Println("table has column age")
	}
}

```



## Check whether the columns of the two tables are the same

```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	tb1, _ := gotable.Create("id", "name")
	tb2, _ := gotable.Create("name", "id")
	fmt.Println(tb1.EqualColumns(tb2))
	// output: false
	// reason: tb1 and tb2 have different columns order

	tb1, _ = gotable.Create("id", "name")
	tb2, _ = gotable.Create("id", "name", "sex")
	fmt.Println(tb1.EqualColumns(tb2))
	// output: false
	// reason: tb1 and tb2 have different columns

	tb1, _ = gotable.Create("id", "name")
	tb2, _ = gotable.Create("id", "name")
	tb2.SetDefault("id", "001")
	fmt.Println(tb1.EqualColumns(tb2))
	// output: false
	// reason: tb1 and tb2 have different column default. (tb1: "", "";    tb2: "001", "")

	tb1, _ = gotable.Create("id", "name")
	tb2, _ = gotable.Create("id", "name")
	tb2.Align("id", gotable.Left)
	fmt.Println(tb1.EqualColumns(tb2))
	// output: false
	// reason: tb1 and tb2 have different column alignment.
	// (tb1: gotable.Center, gotable.Center;    tb2: gotable.Left, gotable.Center)

	tb1, _ = gotable.Create("id", "name")
	tb2, _ = gotable.Create("id", "name")
	fmt.Println(tb1.EqualColumns(tb2))
	// output: true
}

```



## Set columns color

```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	table, err := gotable.Create("Name", "ID", "salary")
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

	table.AddRows(rows)

	// Underline the `salary` column with a white font and no background color
	table.SetColumnColor("salary", gotable.Underline, gotable.Write, gotable.NoneBackground)
	fmt.Println(table)
}

```



## Get table type

```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	table, err := gotable.Create("version", "description")
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	fmt.Println(table.Type())
	// output: simple
}

```



## Custom ending string

```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	table, err := gotable.Create("version", "description")
	if err != nil {
		fmt.Println("Create safe table failed: ", err.Error())
		return
	}

	row := make(map[string]string)
	row["version"] = "v1.0"
	row["description"] = "test"
	table.AddRow(row)

	table.End = ""
	fmt.Println(table)
	fmt.Println("I am a new line after printing table.")
	// outputs:
	// +---------+-------------+ 
	// | version | description | 
	// +---------+-------------+ 
	// |  v1.0   |    test     | 
	// +---------+-------------+ 
	// I am a new line after printing table.
}

```



## Check the table type is simple table

```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	table, err := gotable.Create("China", "US", "UK")
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	fmt.Println(table.IsSimpleTable())
	// output: true
}

```

