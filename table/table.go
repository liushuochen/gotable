package table

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/liushuochen/gotable/cell"
	"github.com/liushuochen/gotable/exception"
	"github.com/liushuochen/gotable/util"
)

const (
	C       = cell.AlignCenter
	L       = cell.AlignLeft
	R       = cell.AlignRight
	Default = "__DEFAULT__"
)

// Table struct:
// - Columns: Save the table columns.
// - Row: Save the list of column and value mapping.
// - border: A flag indicates whether to print table border. If the value is `true`, table will show its border.
//           Default is true(By CreateTable function).
// - tableType: Type of table.
type Table struct {
	*base
	Rows [][]map[string]cell.Cell
}

// CreateTable function returns a pointer of Table.
func CreateTable(set *Set) *Table {
	return &Table{
		base: createTableBase(set, simpleTableType, 1),
		Rows: make([][]map[string]cell.Cell, 1),
	}
}

// Clear the table. The table is cleared of all data.
func (tb *Table) Clear() {
	if tb.partLen != 1 {
		tb.Columns = append(tb.Columns[0:1])
		tb.Rows = append(tb.Rows[0:1])
		tb.partLen = 1
	}
	tb.Columns[0].Clear()
	tb.Rows[0] = make([]map[string]cell.Cell, 0)
}

func (tb *Table) AddPart(columns ...string) error {
	set, err := CreateSetFromString(columns...)
	if err != nil {
		return err
	}
	tb.Rows = append(tb.Rows, make([]map[string]cell.Cell, 0))
	return tb.base.addTableBase(set)
}

// AddColumn method used to add a new column for table. It returns an error when column has been existed.
func (tb *Table) AddColumn(column string) error {
	return tb.AddPNColumn(tb.partLen-1, column)
}

func (tb *Table) AddPNColumn(partNumber int, column string) error {
	if partNumber >= tb.partLen {
		return exception.PartNumber(tb.partLen)
	}
	err := tb.Columns[partNumber].Add(column)
	if err != nil {
		return err
	}

	// Modify exist value, add new column.
	for _, row := range tb.Rows[partNumber] {
		row[column] = cell.CreateEmptyData()
	}
	return nil
}

// AddRow method support Map and Slice argument.
// For Map argument, you must put the data from each row into a Map and use column-data as key-value pairs. If the Map
//   does not contain a column, the table sets it to the default value. If the Map contains a column that does not
//   exist, the AddRow method returns an error.
// For Slice argument, you must ensure that the slice length is equal to the column length. Method will automatically
//   map values in Slice and columns. The default value cannot be omitted and must use gotable.Default constant.
// Return error types:
//   - *exception.UnsupportedRowTypeError: It returned when the type of the argument is not supported.
//   - *exception.RowLengthNotEqualColumnsError: It returned if the argument is type of the Slice but the length is
//       different from the length of column.
//   - *exception.ColumnDoNotExistError: It returned if the argument is type of the Map but contains a nonexistent
//       column as a key.
func (tb *Table) AddRow(row interface{}) error {
	switch v := row.(type) {
	case []string:
		return tb.addRowFromSlice(v)
	case map[string]string:
		return tb.addRowFromMap(v)
	default:
		return exception.UnsupportedRowType(v)
	}
}

func (tb *Table) AddPNRow(partNumber int, row interface{}) error {
	if partNumber >= tb.partLen {
		return exception.PartNumber(tb.partLen)
	}
	switch v := row.(type) {
	case []string:
		return tb.addPNRowFromSlice(partNumber, v)
	case map[string]string:
		return tb.addPNRowFromMap(partNumber, v)
	default:
		return exception.UnsupportedRowType(v)
	}
}

func (tb *Table) addRowFromSlice(row []string) error {
	return tb.addPNRowFromSlice(tb.partLen-1, row)
}

func (tb *Table) addPNRowFromSlice(partNumber int, row []string) error {
	rowLength := len(row)
	if rowLength != tb.Columns[partNumber].Len() {
		return exception.RowLengthNotEqualColumns(rowLength, tb.Columns[partNumber].Len())
	}

	rowMap := make(map[string]string, 0)
	for i := 0; i < rowLength; i++ {
		if row[i] == Default {
			rowMap[tb.Columns[partNumber].base[i].Original()] = tb.Columns[partNumber].base[i].Default()
		} else {
			rowMap[tb.Columns[partNumber].base[i].Original()] = row[i]
		}
	}

	tb.Rows[partNumber] = append(tb.Rows[partNumber], toRow(rowMap))
	return nil
}

func (tb *Table) addRowFromMap(row map[string]string) error {
	return tb.addPNRowFromMap(tb.partLen-1, row)
}

func (tb *Table) addPNRowFromMap(partNumber int, row map[string]string) error {
	for key := range row {
		if !tb.Columns[partNumber].Exist(key) {
			return exception.ColumnDoNotExist(key)
		}

		// add row by const `DEFAULT`
		if row[key] == Default {
			row[key] = tb.Columns[partNumber].Get(key).Default()
		}
	}

	// Add default value
	for _, col := range tb.Columns[partNumber].base {
		_, ok := row[col.Original()]
		if !ok {
			row[col.Original()] = col.Default()
		}
	}

	tb.Rows[partNumber] = append(tb.Rows[partNumber], toRow(row))
	return nil
}

// AddRows used to add a slice of rows maps. It returns a slice of map which add failed.
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

func (tb *Table) SetColumnMaxLength(partNumber int, column string, maxlength int) error {
	if partNumber >= tb.partLen {
		return exception.PartNumber(tb.partLen)
	}
	if !tb.Columns[partNumber].Exist(column) {
		return exception.ColumnDoNotExist(column)
	}
	tb.ColumnMaxLengths[partNumber][column] = maxlength
	//fmt.Println(column, tb.ColumnMaxLengths[column])
	return nil
}

// String method used to implement fmt.Stringer.
func (tb *Table) String() string {
	tag := make(map[string]cell.Cell)
	taga := make([]map[string]cell.Cell, 0)
	border := ""
	switch tb.border {
	case 0:
	case 1:
		border = "-"
	case 2:
		border = "="
	case 3:
		border = "~"
	case 4:
		border = "+"
	}

	for pn, columns := range tb.Columns {
		for _, h := range columns.base {
			if length, exist := tb.ColumnMaxLengths[pn][h.Original()]; !(exist && length > h.Length()) {
				tb.ColumnMaxLengths[pn][h.Original()] = h.Length()
			}
			fmt.Println(h.Original(), tb.ColumnMaxLengths[pn][h.Original()], h.Length())
			tag[h.String()] = cell.CreateData(border)
		}
	}

	//confirm the maxLength
	for pn, row := range tb.Rows {
		for _, data := range row {
			for _, h := range tb.Columns[pn].base {
				maxLength := max(h.Length(), data[h.Original()].Length())
				maxLength = max(maxLength, tb.ColumnMaxLengths[pn][h.Original()])
				tb.ColumnMaxLengths[pn][h.Original()] = maxLength
				//fmt.Println(h.Original(), tb.ColumnMaxLengths[h.Original()])
			}
		}
	}

	content := ""
	taga = append(taga, tag)

	for pn, rows := range tb.Rows {
		// print first line

		if tb.border > 0 {
			// tb.printGroup(taga, columnMaxLength)
			content += tb.printGroup(pn, taga)
		}

		// print table head
		icon := "|"
		if tb.border == 0 {
			icon = " "
		}
		for index, head := range tb.Columns[pn].base {
			itemLen := tb.ColumnMaxLengths[pn][head.Original()]
			if tb.border > 0 {
				itemLen += 2
			}
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

			content += s
		}

		if tb.border > 0 {
			content += "\n"
		}

		// input tableValue
		tableValue := taga
		if !tb.Empty() {
			for _, row := range rows {
				value := make(map[string]cell.Cell)
				for key := range row {
					col := tb.Columns[pn].Get(key)
					value[col.String()] = row[key]
				}
				tableValue = append(tableValue, value)
			}

		}

		content += tb.printGroup(pn, tableValue)
	}

	content += tb.printGroup(tb.partLen-1, taga)

	return tb.end(content)
}

func (tb *Table) end(content string) string {
	content = content[:len(content)-1]
	content += tb.End
	return content
}

// Empty method is used to determine whether the table is empty.
func (tb *Table) Empty() bool {
	return tb.Length() == 0
}

func (tb *Table) Length() int {
	l := 0
	for _, rows := range tb.Rows {
		l += len(rows)
	}
	return l
}

func (tb *Table) PartLength() int {
	return tb.partLen
}

func (tb *Table) GetColumns() []string {
	return tb.GetPNColumns(tb.partLen - 1)
}

func (tb *Table) GetPNColumns(partNumber int) []string {
	if partNumber >= tb.partLen {
		return nil
	}
	columns := make([]string, 0)
	for _, col := range tb.Columns[partNumber].base {
		columns = append(columns, col.Original())
	}
	return columns
}

func (tb *Table) GetValues() []map[string]string {
	return tb.GetPNValues(tb.partLen - 1)
}

func (tb *Table) GetPNValues(partNumber int) []map[string]string {
	if partNumber >= tb.partLen {
		return nil
	}
	values := make([]map[string]string, 0)
	for _, value := range tb.Rows[partNumber] {
		ms := make(map[string]string)
		for k, v := range value {
			ms[k] = v.String()
		}
		values = append(values, ms)
	}
	return values
}

func (tb *Table) PNExist(partNumber int, value map[string]string) bool {
	if partNumber >= tb.partLen {
		return false
	}
	for _, row := range tb.Rows[partNumber] {
		exist := true
		for key := range value {
			v, ok := row[key]
			if !ok || v.String() != value[key] {
				exist = false
				break
			}
		}
		if exist {
			return exist
		}
	}
	return false
}

func (tb *Table) Exist(value map[string]string) bool {
	exist := false
	for pn := 0; pn <= tb.partLen; pn++ {
		exist = tb.PNExist(pn, value)
		if exist {
			return exist
		}
	}
	return false
} ///done

func (tb *Table) json(indent int) ([]byte, error) {
	data := make([]map[string]string, 0)
	for _, rows := range tb.Rows {
		for _, row := range rows {
			element := make(map[string]string)
			for col, value := range row {
				element[col] = value.String()
			}
			data = append(data, element)
		}
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

// The JSON method returns the JSON string corresponding to the gotable. The indent argument represents the indent
// value. If index is less than zero, the JSON method treats it as zero.
func (tb *Table) JSON(indent int) (string, error) {
	bytes, err := tb.json(indent)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// The XML method returns the XML format string corresponding to the gotable. The indent argument represents the indent
// value. If index is less than zero, the XML method treats it as zero.
func (tb *Table) XML(indent int) string {
	if indent < 0 {
		indent = 0
	}
	indents := make([]string, indent)
	for i := 0; i < indent; i++ {
		indents[i] = " "
	}
	indentString := strings.Join(indents, "")

	contents := []string{"<?xml version=\"1.0\" encoding=\"utf-8\" standalone=\"yes\"?>", "<table>"}
	for _, rows := range tb.Rows {
		for _, row := range rows {
			contents = append(contents, indentString+"<row>")
			for name := range row {
				line := indentString + indentString + fmt.Sprintf("<%s>%s</%s>", name, row[name], name)
				contents = append(contents, line)
			}
			contents = append(contents, indentString+"</row>")
		}
	}
	contents = append(contents, "</table>")
	content := strings.Join(contents, "\n")
	return content
}

func (tb *Table) CloseBorder() {
	tb.border = 0
}

func (tb *Table) OpenBorder() {
	tb.border = 1
}

func (tb *Table) SetBorder(border int8) {
	tb.border = border
}

func (tb *Table) Align(column string, mode int) {
	tb.PNAlign(tb.partLen-1, column, mode)
}

func (tb *Table) PNAlign(partNumber int, column string, mode int) {
	if partNumber >= tb.partLen {
		return
	}
	for _, h := range tb.Columns[partNumber].base {
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
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

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
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	writer := csv.NewWriter(file)

	contents := make([][]string, 0)
	for pn := 0; pn < tb.partLen; pn++ {
		columns := tb.GetPNColumns(pn)
		contents = append(contents, columns)
		for _, value := range tb.GetPNValues(pn) {
			content := make([]string, 0)
			for _, col := range columns {
				content = append(content, value[col])
			}
			contents = append(contents, content)
		}
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
	return tb.HasPNColumn(tb.partLen-1, column)
}

func (tb *Table) HasPNColumn(partNumber int, column string) bool {
	if partNumber >= tb.partLen {
		return false
	}
	for index := range tb.Columns[partNumber].base {
		if tb.Columns[partNumber].base[index].Original() == column {
			return true
		}
	}
	return false
}

func (tb *Table) EqualColumns(other *Table) bool {
	if tb.partLen == other.partLen {
		for pn := 0; pn < tb.partLen; pn++ {
			if !tb.Columns[pn].Equal(other.Columns[pn]) {
				return false
			}
		}
		return true
	}
	return false
}

func (tb *Table) SetColumnColor(columnName string, display, fount, background int) {
	background += 10
	for _, col := range tb.Columns[tb.partLen-1].base {
		if col.Original() == columnName {
			col.SetColor(display, fount, background)
			break
		}
	}
}
