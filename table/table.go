package table

import (
	"encoding/json"
	"fmt"
	"github.com/liushuochen/gotable/cell"
	"github.com/liushuochen/gotable/header"
	"github.com/liushuochen/gotable/util"
	"strings"
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
// TODO: removed in 3.0
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
// TODO: removed in 3.0
func (tb *Table) AddValue(newValue map[string]string) error {
	util.DeprecatedTips("AddValue", "AddRow", "3.0", "method")
	return tb.addValue(newValue)
}

// Deprecated
// TODO: removed in 3.0
func (tb *Table) addValue(newValue map[string]string) error {
	return tb.AddRow(newValue)
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

func (tb *Table) AddRows(rows []map[string]string) []map[string]string {
	failure := make([]map[string]string, 0)
	for _, row := range rows {
		err := tb.AddRow(row)
		if err != nil {
			failure = append(failure, row)
		}
	}
	return failure
}

// Deprecated
// TODO: removed in 3.0
func (tb *Table) AddValues(values []map[string]string) []map[string]string {
	util.DeprecatedTips("AddValues", "AddRows", "3.0", "method")
	return tb.AddRows(values)
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

func (tb *Table) GetColumns() []string {
	columns := make([]string, 0)
	for _, column := range tb.Columns.base {
		columns = append(columns, column.String())
	}
	return columns
}

// Deprecated
// TODO: removed in 3.0
func (tb *Table) GetHeaders() []string {
	util.DeprecatedTips("GetHeaders", "GetColumns", "3.0", "method")
	return tb.GetColumns()
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

func (tb *Table) json(indent int) ([]byte, error) {
	data := make([]map[string]string, 0)
	for _, row := range tb.Row {
		element := make(map[string]string)
		for column, value := range row {
			element[column] = value.String()
		}
		data = append(data, element)
	}

	if indent < 0 {
		indent = 0
	}
	elems := make([]string, 0)
	for i := 0; i < indent; i++ {
		elems = append(elems, " ")
	}

	return json.MarshalIndent(data, "", strings.Join(elems, " "))
}

func (tb *Table) Json(indent int) (string, error) {
	bytes, err := tb.json(indent)
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

func (tb *Table) Align(column string, mode int) {
	for _, h := range tb.Columns.base {
		if h.String() == column {
			h.SetAlign(mode)
			return
		}
	}
}
