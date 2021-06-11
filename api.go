package gotable

import (
	"fmt"
	"github.com/liushuochen/gotable/constant"
	"github.com/liushuochen/gotable/table"
)

const (
	Center = table.C
	Left = table.L
	Right = table.R
	Default = table.Default
)

// Deprecated
func CreateTable(header []string) (*table.Table, error) {
	fmt.Println("Function `CreateTable` will no longer supported." +
		" You can use the `Create` function instead of `CreateTable` function." +
		" This function will be removed in version 3.0.")
	return Create(header...)
}

func Create(columns ...string) (*table.Table, error) {
	set := &table.Set{}
	for _, column := range columns {
		err := set.Add(column)
		if err != nil {
			return nil, err
		}
	}
	tb := table.CreateTable(set)
	return tb, nil
}

func Version() string { return constant.GetVersion() }
