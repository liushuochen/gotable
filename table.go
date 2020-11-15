package gotable

import (
	"fmt"
	"github.com/liushuochen/gotable/color"
	"reflect"
)

type Table struct {
	Header Set
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
	return color.WHITE
}

//Sequence sequence for print
type Sequence interface {
	Val() string
	// actual length except invisible rune
	Len() int
}

type DefaultSequence string

func (s DefaultSequence) Val() string {
	return string(s)
}

func (s DefaultSequence) Len() int {
	return len(s)
}

func CreateTable(header []string, options ...Option) (*Table, error) {
	set := Set{}
	for _, head := range header {
		err := set.Add(head)
		if err != nil {
			return nil, err
		}
	}

	tb := &Table {
		Header: set,
		Value:  make([]map[string]Sequence, 0),
	}

	opts := &Options {
		ColorController: defaultController,
	}

	for _, do := range options {
		do(opts)
	}

	tb.Opts = opts
	return tb, nil
}

func CreateTableFromStruct(meta interface{}, options ...Option) (*Table, error) {
	typ := reflect.TypeOf(meta)
	var header []string
	for i := 0; i < typ.NumField(); i++ {
		header = append(header, typ.Field(i).Name)
	}
	return CreateTable(header, options...)
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

func (tb *Table) AddValue(newValue map[string]Sequence) error {
	for key := range newValue {
		if tb.Header.Exist(key) {
			value := reflect.ValueOf(newValue[key])
			clr := tb.Opts.ColorController(key, value)
			newValue[key] = color.ColorfulString(clr, value)
			continue
		} else {
			err := fmt.Errorf("invalid value %s", key)
			return err
		}
	}

	tb.Value = append(tb.Value, newValue)
	return nil
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
		columnMaxLength[header] = 0
		tag[header] = DefaultSequence("-")
	}

	for _, data := range tb.Value {
		for _, header := range tb.Header.base {
			maxLength := max(len(header), data[header].Len())
			maxLength = max(maxLength, columnMaxLength[header])
			columnMaxLength[header] = maxLength
		}
	}

	// print first line
	taga = append(taga, tag)
	printGroup(taga, tb.Header.base, columnMaxLength)

	// print table head
	for index, head := range tb.Header.base {
		itemLen := columnMaxLength[head] + 4
		s, _ := center(DefaultSequence(head), itemLen, " ")
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
	return tb.Header.base
}

func (tb *Table) GetValues() []map[string]Sequence {
	return tb.Value
}

// TODO if value in table, return true.
func (tb *Table) Exist(head string, value interface{}) bool {
	headExit := false
	for _, headInHeader := range tb.Header.base {
		if head == headInHeader {
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

func printGroup(group []map[string]Sequence, header []string, columnMaxLen map[string]int) {
	for _, item := range group {
		for index, head := range header {
			itemLen := columnMaxLen[head] + 4
			s := ""
			if item[head].Val() == "-" {
				s, _ = center(item[head], itemLen, "-")
			} else {
				s, _ = center(item[head], itemLen, " ")
			}

			icon := "|"
			if item[head].Val() == "-" {
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

func center(str Sequence, length int, fillchar string) (string, error) {
	if len(fillchar) != 1 {
		err := fmt.Errorf("the fill character must be exactly one" +
			" character long")
		return "", err
	}

	if str.Len() >= length {
		return str.Val(), nil
	}

	result := ""
	if isEvenNumber(length - str.Len()) {
		front := ""
		for i := 0; i < ((length - str.Len()) / 2); i++ {
			front = front + fillchar
		}

		result = front + str.Val() + front
	} else {
		front := ""
		for i := 0; i < ((length - str.Len() - 1) / 2); i++ {
			front = front + fillchar
		}

		behind := front + fillchar
		result = front + str.Val() + behind
	}
	return result, nil
}

func isEvenNumber(number int) bool {
	if number%2 == 0 {
		return true
	}
	return false
}
