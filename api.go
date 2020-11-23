package gotable

import (
	"github.com/liushuochen/gotable/constant"
	"reflect"
)

func CreateTable(header []string, options ...Option) (*Table, error) {
	set := &Set{}
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

func Version() string { return constant.GetVersion() }
