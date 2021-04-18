package gotable

import (
	"github.com/liushuochen/gotable/constant"
	"github.com/liushuochen/gotable/table"
)

const (
	Center = table.C
	Left = table.L
	Right = table.R
)

func CreateTable(header []string) (*table.Table, error) {
	set := &table.Set{}
	for _, head := range header {
		err := set.Add(head)
		if err != nil {
			return nil, err
		}
	}
	tb := table.CreateTable(set)
	return tb, nil
}

func Version() string { return constant.GetVersion() }
