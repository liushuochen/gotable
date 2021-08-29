# Gotable Demo
In this section, we have written some demos for your reference.
Click [here](../README.md) to return to home page.

## Create a table
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
	tb, err := gotable.CreateByStruct(&Student{})
	if err != nil {
		fmt.Println("Create table failed.")
		return
	}
}

```

## Load data from CSV file
```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	table, err := gotable.ReadFromCSVFile("cmd/demo.csv")
	if err != nil {
		fmt.Println("read failed: ", err.Error())
		return
	}
}

```

## Load data from JSON file
```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func main() {
	table, err := gotable.ReadFromJSONFile("cmd/demo.json")
	if err != nil {
		fmt.Println("[ERROR] ", err.Error())
		return
	}
}

```

## Add row
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

## Print table
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

## Clear data
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
	tb.PrintTable()

	tb.Clear()
	fmt.Println("After the table data is cleared...")
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
After the table data is cleared...

```

## Set default value
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

## Drop default value
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

## Get default value
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

## Open border
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
	
	// Underline the `salary` column with a white font and no background color
	tb.SetColumnColor("salary", gotable.Underline, gotable.Write, gotable.NoneBackground)
	tb.PrintTable()
}

```

