package gotable

import (
	"encoding/json"
	"fmt"
	"github.com/liushuochen/gotable/color"
	"reflect"
)

type Table struct {
	Header *Set
	Value  []map[string]Sequence
	Opts   *Options
}

type Options struct {
	ColorController ColorController
}

type Option func(t *Options)

func WithColorController(controller ColorController) Option {
	return func(o *Options) {
		o.ColorController = controller
	}
}

type ColorController func(field string, val reflect.Value) color.Color

func defaultController(field string, val reflect.Value) color.Color {
	return ""
}

// Sequence sequence for print
type Sequence interface {
	Value() string

	// Actual length except invisible rune
	Len() int

	// Origin value
	OriginValue() string
}

type DefaultSequence string

func (s DefaultSequence) Value() string {
	return string(s)
}

func (s DefaultSequence) Len() int {
	return len(s)
}

func (s DefaultSequence) OriginValue() string {
	return s.Value()
}

func (tb *Table) AddHead(newHead string) error {
	err := tb.Header.Add(newHead)
	if err != nil {
		return err
	}

	// modify exit value, add new column.
	for _, data := range tb.Value {
		data[newHead] = nil
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

func (tb *Table) AddValue(newValue map[string]Sequence) error {
	for key := range newValue {
		value := reflect.ValueOf(newValue[key])
		clr := tb.Opts.ColorController(key, value)
		newValue[key] = color.ColorfulString(clr, value)
	}

	return tb.addValue(newValue)
}

func (tb *Table) addValue(newValue map[string]Sequence) error {
	for key := range newValue {
		if tb.Header.Exist(key) {
			continue
		} else {
			err := fmt.Errorf("invalid value %s", key)
			return err
		}
	}

	for _, head := range tb.Header.base {
		_, ok := newValue[head.Name]
		if !ok {
			newValue[head.Name] = DefaultSequence(head.Default())
		}
	}

	tb.Value = append(tb.Value, newValue)
	return nil
}

func (tb *Table) AddValuesFromSlice(items []interface{}) error {
	structToMap := func(item interface{}) map[string]Sequence {
		typ := reflect.TypeOf(item)
		val := reflect.ValueOf(item)
		mp := make(map[string]Sequence, val.NumField())
		for i := 0; i < val.NumField(); i++ {
			fieldName := typ.Field(i).Name
			fieldVal := val.FieldByName(fieldName)
			clr := tb.Opts.ColorController(fieldName, fieldVal)
			mp[fieldName] = color.ColorfulString(clr, fieldVal)
		}
		return mp
	}

	for _, item := range items {
		err := tb.addValue(structToMap(item))
		if err != nil {
			return err
		}
	}

	return nil
}

func (tb *Table) PrintTable() {
	if tb.Empty() {
		fmt.Println("table is empty.")
		return
	}

	columnMaxLength := make(map[string]int)
	tag := make(map[string]Sequence)
	taga := make([]map[string]Sequence, 0)
	for _, header := range tb.Header.base {
		columnMaxLength[header.Name] = len(header.Name)
		tag[header.Name] = DefaultSequence("-")
	}

	for _, data := range tb.Value {
		for _, header := range tb.Header.base {
			maxLength := max(len(header.Name), data[header.Name].Len())
			maxLength = max(maxLength, columnMaxLength[header.Name])
			columnMaxLength[header.Name] = maxLength
		}
	}

	// print first line
	taga = append(taga, tag)
	printGroup(taga, tb.Header.base, columnMaxLength)

	// print table head
	for index, head := range tb.Header.base {
		itemLen := columnMaxLength[head.Name] + 4
		s, _ := center(DefaultSequence(head.Name), itemLen, " ")
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
	tableValue = append(tableValue, tb.Value...)
	tableValue = append(tableValue, tag)
	printGroup(tableValue, tb.Header.base, columnMaxLength)
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

func (tb *Table) GetColoredValues() []map[string]Sequence {
	return tb.Value
}

func (tb *Table) GetValues() []map[string]string {
	values := make([]map[string]string, 0)
	for _, tableValueMap := range tb.Value {
		originValueMap := make(map[string]string)
		for key := range tableValueMap {
			originValueMap[key] = tableValueMap[key].OriginValue()
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
			element[head] = value.OriginValue()
		}
		data = append(data, element)
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}