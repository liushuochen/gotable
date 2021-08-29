package table

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/liushuochen/gotable/cell"
	"github.com/liushuochen/gotable/column"
	"github.com/liushuochen/gotable/exception"
	"github.com/liushuochen/gotable/util"
	"os"
	"strings"
)

const (
	C = column.AlignCenter
	L = column.AlignLeft
	R = column.AlignRight
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

// Clear the table. The table is cleared of all data.
func (tb *Table) Clear() {
	tb.Columns.Clear()
	tb.Row = make([]map[string]cell.Cell, 0)
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

func (tb *Table) SetDefault(h string, defaultValue string) {
	for _, head := range tb.Columns.base {
		if head.Original() == h {
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
		if head.Original() == h {
			return head.Default()
		}
	}
	return ""
}

func (tb *Table) GetDefaults() map[string]string {
	defaults := make(map[string]string)
	for _, h := range tb.Columns.base {
		defaults[h.Original()] = h.Default()
	}
	return defaults
}

func (tb *Table)  AddRow(row map[string]string) error {
	for key := range row {
		if !tb.Columns.Exist(key) {
			return exception.ColumnDoNotExist(key)
		}

		// add row by const `DEFAULT`
		if row[key] == Default {
			row[key] = tb.Columns.Get(key).Default()
		}
	}

	// Add default value
	for _, col := range tb.Columns.base {
		_, ok := row[col.Original()]
		if !ok {
			row[col.Original()] = col.Default()
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

// PrintTable method used to print table data in STDOUT
func (tb *Table) PrintTable() {
	columnMaxLength := make(map[string]int)
	tag := make(map[string]cell.Cell)
	taga := make([]map[string]cell.Cell, 0)
	for _, h := range tb.Columns.base {
		columnMaxLength[h.Original()] = h.Length()
		tag[h.String()] = cell.CreateData("-")
	}

	for _, data := range tb.Row {
		for _, h := range tb.Columns.base {
			maxLength := max(h.Length(), data[h.Original()].Length())
			maxLength = max(maxLength, columnMaxLength[h.Original()])
			columnMaxLength[h.Original()] = maxLength
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
		itemLen := columnMaxLength[head.Original()] + 2
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
	if !tb.Empty() {
		for _, row := range tb.Row {
			value := make(map[string]cell.Cell)
			for key := range row {
				col := tb.Columns.Get(key)
				value[col.String()] = row[key]
			}
			tableValue = append(tableValue, value)
		}
		tableValue = append(tableValue, tag)
	}
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
	for _, col := range tb.Columns.base {
		columns = append(columns, col.Original())
	}
	return columns
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
		for col, value := range row {
			element[col] = value.String()
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
		if h.Original() == column {
			h.SetAlign(mode)
			return
		}
	}
}

func (tb *Table) ToJsonFile(path string, indent int) error {
	if !util.IsJsonFile(path) {
		return fmt.Errorf("%s: not a regular json file", path)
	}

	bytes, err := tb.json(indent)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}

func (tb *Table) ToCSVFile(path string) error {
	if !util.IsCSVFile(path) {
		return exception.NotARegularCSVFile(path)
	}
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := csv.NewWriter(file)

	contents := make([][]string, 0)
	columns := tb.GetColumns()
	contents = append(contents, columns)
	for _, value := range tb.GetValues() {
		content := make([]string, 0)
		for _, col := range columns {
			content = append(content, value[col])
		}
		contents = append(contents, content)
	}

	err = writer.WriteAll(contents)
	if err != nil {
		return err
	}
	writer.Flush()
	err = writer.Error()
	if err != nil {
		return err
	}
	return nil
}

func (tb *Table) HasColumn(column string) bool {
	for index := range tb.Columns.base {
		if tb.Columns.base[index].Original() == column {
			return true
		}
	}
	return false
}

func (tb *Table) EqualColumns(other *Table) bool {
	return tb.Columns.Equal(other.Columns)
}

func (tb *Table) SetColumnColor(columnName string, display, fount, background int) {
	background += 10
	for _, col := range tb.Columns.base {
		if col.Original() == columnName {
			col.SetColor(display, fount, background)
			break
		}
	}
}
