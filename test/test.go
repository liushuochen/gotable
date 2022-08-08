package main

import (
	"fmt"

	"github.com/liushuochen/gotable"
)

func main() {
	fmt.Println("sss")
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

	// row3 := []string{"+", "+", "+"}
	// err = table.AddRow(row3)
	// if err != nil {
	// 	fmt.Println("Add value to table failed: ", err.Error())
	// 	return
	// }

	// Add new part with new columns
	table.AddPart("name", "salary")

	rows := make([]map[string]string, 0)
	for i := 0; i < 3; i++ {
		row := make(map[string]string)
		row["name"] = fmt.Sprintf("employee-%d", i)
		row["salary"] = "60000"
		rows = append(rows, row)
	}
	table.AddRows(rows)

	// Add row to previous part
	row3 := []string{"WuHan", "April", "Blank"}
	err = table.AddPNRow(0, row3)
	if err != nil {
		fmt.Println("Add value to table failed: ", err.Error())
		return
	}

	table.AddPart("name", "salary")

	rows = make([]map[string]string, 0)
	for i := 0; i < 3; i++ {
		row := make(map[string]string)
		row["name"] = fmt.Sprintf("employee-%d", i)
		row["salary"] = "60000"
		rows = append(rows, row)
	}
	table.AddRows(rows)

	table.SetBorder(4) //0" " 1"-" 2"=" 3"~" 4"+"

	//get columns and maps in specific part
	fmt.Println(table.GetPNColumns(1))
	fmt.Println(table.GetPNValues(0))

	//change any column length in any part
	table.SetColumnMaxLength(1, "name", 20)
	table.SetColumnMaxLength(1, "salary", 14)
	table.SetColumnMaxLength(2, "name", 20)
	table.SetColumnMaxLength(2, "salary", 7)

	fmt.Println(table)

	//table.CloseBorder()

	//fmt.Println(table)
	// outputs:
	// +----------+------------------+---------+
	// |  China   |        US        | French  |
	// +----------+------------------+---------+
	// | Beijing  | Washington, D.C. |  Paris  |
	// | Yinchuan |   Los Angeles    | Orleans |
	// +----------+------------------+---------+
}
