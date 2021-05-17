package table

import (
	"encoding/json"
	"fmt"
	"github.com/liushuochen/gotable/header"
)

const (
	C = header.AlignCenter
	L = header.AlignLeft
	R = header.AlignRight
)

type Table struct {
	Header 	*Set
	Value  	[]map[string]string
	border	bool
}

func CreateTable(set *Set) *Table {
	return &Table{
		Header: set,
		Value: make([]map[string]string, 0),
		border: true,
	}
}

func (tb *Table) AddHead(newHead string) error {
	err := tb.Header.Add(newHead)
	if err != nil {
		return err
	}

	// modify exit value, add new column.
	for _, data := range tb.Value {
		data[newHead] = ""
	}

	return nil
}

func (tb *Table) SetDefault(h string, defaultValue string) {
	for _, head := range tb.Header.base {
		if head.Name == h {
			head.SetDefault(defaultValue)
			break
		}
	}
}

func (tb *Table) GetDefault(h string) string {
	for _, head := range tb.Header.base {
		if head.Name == h {
			return head.Default()
		}
	}
	return ""
}

func (tb *Table) AddValue(newValue map[string]string) error {
	return tb.addValue(newValue)
}

func (tb *Table) addValue(newValue map[string]string) error {
	for key := range newValue {
		if !tb.Header.Exist(key) {
			err := fmt.Errorf("invalid value %s", key)
			return err
		}
	}

	for _, head := range tb.Header.base {
		_, ok := newValue[head.Name]
		if !ok {
			newValue[head.Name] = head.Default()
		}
	}

	tb.Value = append(tb.Value, newValue)
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
	tag := make(map[string]string)
	taga := make([]map[string]string, 0)
	for _, h := range tb.Header.base {
		columnMaxLength[h.Name] = len(h.Name)
		tag[h.Name] = "-"
	}

	for _, data := range tb.Value {
		for _, h := range tb.Header.base {
			maxLength := max(len(h.Name), len(data[h.Name]))
			maxLength = max(maxLength, columnMaxLength[h.Name])
			columnMaxLength[h.Name] = maxLength
		}
	}

	// print first line
	taga = append(taga, tag)
	if tb.border {
		printGroup(taga, tb.Header.base, columnMaxLength, tb.border)
	}

	// print table head
	icon := "|"
	if !tb.border { icon = "" }
	for index, head := range tb.Header.base {
		itemLen := columnMaxLength[head.Name] + 2
		s := ""
		switch head.Align() {
		case R:
			s, _ = right(head.Name, itemLen, " ")
		case L:
			s, _ = left(head.Name, itemLen, " ")
		default:
			s, _ = center(head.Name, itemLen, " ")
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
	tableValue = append(tableValue, tb.Value...)
	tableValue = append(tableValue, tag)
	printGroup(tableValue, tb.Header.base, columnMaxLength, tb.border)
}

func (tb *Table) Empty() bool {
	if len(tb.Value) == 0 {
		return true
	}
	return false
}

func (tb *Table) GetHeaders() []string {
	result := make([]string, 0)
	for _, head := range tb.Header.base {
		result = append(result, head.Name)
	}
	return result
}

func (tb *Table) GetValues() []map[string]string {
	values := make([]map[string]string, 0)
	for _, tableValueMap := range tb.Value {
		originValueMap := make(map[string]string)
		for key := range tableValueMap {
			originValueMap[key] = tableValueMap[key]
		}
		values = append(values, originValueMap)
	}
	return values
}

func (tb *Table) Exist(head string, value interface{}) bool {
	headExit := false
	for _, headInHeader := range tb.Header.base {
		if head == headInHeader.Name {
			headExit = true
			break
		}
	}
	if !headExit {
		return headExit
	}

	find := false
	for _, data := range tb.Value {
		if data[head] == value {
			find = true
			break
		}
	}
	return find
}

func (tb *Table) Json() (string, error) {
	data := make([]map[string]interface{}, 0)
	for _, v := range tb.Value {
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
	for _, h := range tb.Header.base {
		if h.Name == head {
			h.SetAlign(mode)
			return
		}
	}
}
