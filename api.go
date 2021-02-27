package gotable

import (
	"github.com/liushuochen/gotable/constant"
	table "github.com/liushuochen/gotable/table"
	"reflect"
)

func CreateTable(header []string, options ... table.Option) (*table.Table, error) {
	set := &table.Set{}
	for _, head := range header {
		err := set.Add(head)
		if err != nil {
			return nil, err
		}
	}

	tb := &table.Table {
		Header: set,
		Value:  make([]map[string]table.Sequence, 0),
	}

	opts := &table.Options {
		ColorController: table.DefaultController,
	}

	for _, do := range options {
		do(opts)
	}

	tb.Opts = opts
	return tb, nil
}

func CreateTableFromStruct(meta interface{}, options ...table.Option) (*table.Table, error) {
	typ := reflect.TypeOf(meta)
	var header []string
	for i := 0; i < typ.NumField(); i++ {
		header = append(header, typ.Field(i).Name)
	}
	return CreateTable(header, options...)
}

func Version() string { return constant.GetVersion() }

func CreateEmptyValueMap() map[string]table.Sequence {
	return make(map[string]table.Sequence)
}

func CreateValue(value string) table.TableValue {
	return table.TableValue(value)
}

func WithColorController(controller table.ColorController) table.Option {
	return table.WithColorController(controller)
}
