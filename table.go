// Copyright TCatTime
// Developed in 2019.11
// Print table which data is ASCII code in CLI
// ...
// table.go provides functions and methods for external calls

package gotable

import (
	"fmt"
	"gotable/constant"
)

type Table struct {
	header 	Set
	value 	[]map[string]string
}

func (tb *Table) AddHead(newHead string) error {
	err := tb.header.Add(newHead)
	if err != nil {
		return err
	}

	// modify exit value, add new column.
	for _, data := range tb.value {
		data[newHead] = ""
	}

	return nil
}

func (tb *Table) AddValue(newValue map[string]string) error {
	for key, _ := range newValue {
		if tb.header.Exit(key) {
			continue
		} else {
			err := fmt.Errorf("invalid value %s", key)
			return err
		}
	}

	tb.value = append(tb.value, newValue)
	return nil
}

func (tb *Table) PrintTable() {
	if tb.Empty() {
		fmt.Println("table is empty.")
		return
	}

	columnMaxLength := make(map[string]int)
	tag := make(map[string]string)
	taga := make([]map[string]string, 0)
	for _, header := range tb.header.base {
		columnMaxLength[header] = 0
		tag[header] = "-"
	}

	for _, data := range tb.value {
		for _, header := range tb.header.base {
			maxLength := max(len(header), len(data[header]))
			maxLength = max(maxLength, columnMaxLength[header])
			columnMaxLength[header] = maxLength
		}
	}

	// print first line
	taga = append(taga, tag)
	printGroup(taga, tb.header.base, columnMaxLength)

	// print table head
	for index, head := range tb.header.base {
		itemLen := columnMaxLength[head] + 4
		s, _ := center(head, itemLen, " ")
		if index == 0 {
			s = "|" + s + "|"
		} else {
			s = "" + s + "|"
		}

		fmt.Print(s)
	}
	fmt.Println()

	// print value
	tableValue := taga
	tableValue = append(tableValue, tb.value...)
	tableValue = append(tableValue, tag)
	printGroup(tableValue, tb.header.base, columnMaxLength)
}

func (tb *Table) Empty() bool {
	if len(tb.value) == 0 {
		return true
	}
	return false
}

func (tb *Table) GetHeaders() []string {
	return tb.header.base
}

func (tb *Table) GetValues() []map[string]string {
	return tb.value
}

// TODO if value in table, return true.
func (tb *Table) Exit(head string, value interface{}) bool {
	headExit := false
	for _, headInHeader := range tb.header.base {
		if head == headInHeader {
			headExit = true
			break
		}
	}
	if !headExit {
		return headExit
	}

	find := false
	for _, data := range tb.value {
		if data[head] == value {
			find = true
			break
		}
	}
	return find
}

func printGroup(group []map[string]string, header []string, columnMaxLen map[string]int) {
	for _, item := range group {
		for index, head := range header {
			itemLen := columnMaxLen[head] + 4
			s := ""
			if item[head] == "-" {
				s, _ = center(item[head], itemLen, "-")
			} else {
				s, _ = center(item[head], itemLen, " ")
			}

			icon := "|"
			if item[head] == "-" {
				icon = "+"
			}

			if index == 0 {
				s = icon + s + icon
			} else {
				s = "" + s + icon
			}
			fmt.Print(s)
		}
		fmt.Println()
	}
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func center(str string, length int, fillchar string) (string, error) {
	if len(fillchar) != 1 {
		err := fmt.Errorf("the fill character must be exactly one" +
			" character long")
		return "", err
	}

	if len(str) >= length {
		return str, nil
	}

	result := ""
	if isEvenNumber(length - len(str)) {
		front := ""
		for i := 0; i < ((length - len(str)) / 2); i++ {
			front = front + fillchar
		}

		result = front + str + front
	} else {
		front := ""
		for i := 0; i < ((length - len(str) - 1) / 2); i++ {
			front = front + fillchar
		}

		behind := front + fillchar
		result = front + str + behind
	}
	return result, nil
}

func isEvenNumber(number int) bool {
	if number % 2 == 0 {
		return true
	}
	return false
}

func CreateTable(header []string) (*Table, error) {
	set := Set{}
	for _, head := range header {
		err := set.Add(head)
		if err != nil {
			return nil, err
		}
	}

	tb := &Table{
		header: set,
		value:  make([]map[string]string, 0),
	}
	return tb, nil
}

func Version() string {
	return constant.Version
}