// Copyright TCatTime
// Developed in 2019.11
// Print table which data is ASCII code in CLI
// ...
// table.go provides functions and methods for external calls

package gotable

import (
	"gotable/constant"
	"gotable/table"
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
		Value:  make([]map[string]string, 0),
	}
	return tb, nil
}

func Version() string { return constant.Version }