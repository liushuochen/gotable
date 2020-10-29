// Copyright TCatTime
// Developed in 2019.11
// Print table which data is ASCII code in CLI
// ...
// table.go provides functions and methods for external calls

package gotable

import (
	"errors"
	"reflect"

	"github.com/TuoAiTang/gotable/constant"
	"github.com/TuoAiTang/gotable/table"
)


func CreateTable(header []string) (*table.Table, error) {
	set := table.Set{}
	for _, head := range header {
		err := set.Add(head)
		if err != nil {
			return nil, err
		}
	}

	tb := &table.Table{
		Header: set,
		Value:  make([]map[string]table.Sequence, 0),
	}
	return tb, nil
}

func PrintSlice(items []interface{}, controller table.ColorController) error {
	if len(items) == 0 {
		return errors.New("数组为空")
	}

	t, err := createTable(items[0], controller)
	if err != nil {
		return err
	}

	err = t.FillItems(items)
	if err != nil {
		return err
	}

	t.PrintTable()

	return nil
}

func Version() string { return constant.Version }


func createTable(meta interface{}, colorController table.ColorController) (*table.Table, error) {
	typ := reflect.TypeOf(meta)
	var header []string
	for i := 0; i < typ.NumField(); i++ {
		header = append(header, typ.Field(i).Name)
	}
	t, err := CreateTable(header)
	if err != nil {
		return nil, err
	}
	t.ColorController = colorController
	return t, nil
}