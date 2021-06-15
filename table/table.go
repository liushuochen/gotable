package table

import (
	"encoding/json"
	"fmt"
	"github.com/liushuochen/gotable/cell"
	"github.com/liushuochen/gotable/header"
	"github.com/liushuochen/gotable/util"
)

const (
	C = header.AlignCenter
	L = header.AlignLeft
	R = header.AlignRight
	Default = "__DEFAULT__"
)

type Table struct {
	Columns *Set
	Row  	[]map[string]cell.Cell
	border	bool
}

func CreateTable(set *Set) *Table {
	return &Table{
		Columns: set,
		Row: make([]map[string]cell.Cell, 0),
		border: true,
	}
}

func (tb *Table) AddColumn(column string) error {
	err := tb.Columns.Add(column)
	if err != nil {
		return err
	}

	// modify exist value, add new column.
	for _, row := range tb.Row {
		row[column] = cell.CreateEmptyData()
	}
	return nil
}

// Deprecated
func (tb *Table) AddHead(newHead string) error {
	util.DeprecatedTips("AddHead", "AddColumn", "3.0", "method")
	return tb.AddColumn(newHead)
}

func (tb *Table) SetDefault(h string, defaultValue string) {
	for _, head := range tb.Columns.base {
		if head.String() == h {
			head.SetDefault(defaultValue)
			break
		}
	}
}

func (tb *Table) DropDefault(h string) {
	tb.SetDefault(h, "")
}

func (tb *Table) GetDefault(h string) string {
	for _, head := range tb.Columns.base {
		if head.String() == h {
			return head.Default()
		}
	}
	return ""
}

func (tb *Table) GetDefaults() map[string]string {
	defaults := make(map[string]string)
	for _, h := range tb.Columns.base {
		defaults[h.String()] = h.Default()
	}
	return defaults
}

// Deprecated
func (tb *Table) AddValue(newValue map[string]string) error {
	util.DeprecatedTips("AddValue", "AddRow", "3.0", "method")
	return tb.addValue(newValue)
}

func (tb *Table) addValue(newValue map[string]string) error {
	for key := range newValue {
		if !tb.Columns.Exist(key) {
			err := fmt.Errorf("invalid value %s", key)
			return err
		}

		// add value by const `DEFAULT`
		if newValue[key] == Default {
			newValue[key] = tb.Columns.Get(key).Default()
		}
	}

	for _, head := range tb.Columns.base {
		_, ok := newValue[head.String()]
		if !ok {
			newValue[head.String()] = head.Default()
		}
	}

	tb.Row = append(tb.Row, toRow(newValue))
	return nil
}

func (tb *Table) AddRow(row map[string]string) error {
	for key := range row {
		if !tb.Columns.Exist(key) {
			return fmt.Errorf("invalid value %s", key)
		}

		// add row by const `DEFAULT`
		if row[key] == Default {
			row[key] = tb.Columns.Get(key).Default()
		}
	}

	for _, column := range tb.Columns.base {
		_, ok := row[column.String()]
		if !ok {
			row[column.String()] = column.Default()
		}
	}

	tb.Row = append(tb.Row, toRow(row))
	return nil
}

func (tb *Table) AddValues(values []map[string]string) []map[string]string {
	failure := make([]map[string]string, 0)
	for _, value := range values {
		err := tb.AddValue(value)
		if err != nil {
			failure = append(failure, value)
		}
	}
	return failure
}

func (tb *Table) PrintTable() {
	if tb.Empty() {
		fmt.Println("table is empty.")
		return
	}

	columnMaxLength := make(map[string]int)
	tag := make(map[string]cell.Cell)
	taga := make([]map[string]cell.Cell, 0)
	for _, h := range tb.Columns.base {
		columnMaxLength[h.String()] = h.Length()
		tag[h.String()] = cell.CreateData("-")
	}

	for _, data := range tb.Row {
		for _, h := range tb.Columns.base {
			maxLength := max(h.Length(), data[h.String()].Length())
			maxLength = max(maxLength, columnMaxLength[h.String()])
			columnMaxLength[h.String()] = maxLength
		}
	}

	// print first line
	taga = append(taga, tag)
	if tb.border {
		printGroup(taga, tb.Columns.base, columnMaxLength, tb.border)
	}

	// print table head
	icon := "|"
	if !tb.border { icon = "" }
	for index, head := range tb.Columns.base {
		itemLen := columnMaxLength[head.String()] + 2
		s := ""
		switch head.Align() {
		case R:
			s, _ = right(head, itemLen, " ")
		case L:
			s, _ = left(head, itemLen, " ")
		default:
			s, _ = center(head, itemLen, " ")
		}
		if index == 0 {
			s = icon + s + icon
		} else {
			s = "" + s + icon
		}

		fmt.Print(s)
	}

	if tb.border {
		fmt.Println()
	}

	// print value
	tableValue := taga
	tableValue = append(tableValue, tb.Row...)
	tableValue = append(tableValue, tag)
	printGroup(tableValue, tb.Columns.base, columnMaxLength, tb.border)
}

func (tb *Table) Empty() bool {
	return tb.Length() == 0
}

func (tb *Table) Length() int {
	return len(tb.Row)
}

func (tb *Table) GetHeaders() []string {
	result := make([]string, 0)
	for _, head := range tb.Columns.base {
		result = append(result, head.String())
	}
	return result
}

func (tb *Table) GetValues() []map[string]string {
	values := make([]map[string]string, 0)
	for _, value := range tb.Row {
		ms := make(map[string]string)
		for k, v := range value {
			ms[k] = v.String()
		}
		values = append(values, ms)
	}
	return values
}

func (tb *Table) Exist(value map[string]string) bool {
	for _, row := range tb.Row {
		exist := true
		for key := range value {
			v, ok := row[key]
			if !ok || v.String() != value[key] {
				exist = false
				break
			}
		}
		if exist { return exist }
	}
	return false
}

func (tb *Table) Json() (string, error) {
	data := make([]map[string]interface{}, 0)
	for _, v := range tb.Row {
		element := make(map[string]interface{})
		for head, value := range v {
			element[head] = value
		}
		data = append(data, element)
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (tb *Table) CloseBorder() {
	tb.border = false
}

func (tb *Table) OpenBorder() {
	tb.border = true
}

func (tb *Table) Align(head string, mode int) {
	for _, h := range tb.Columns.base {
		if h.String() == head {
			h.SetAlign(mode)
			return
		}
	}
}
