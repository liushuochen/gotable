package gotable

import (
	"github.com/liushuochen/gotable/constant"
	"github.com/liushuochen/gotable/table"
	"github.com/liushuochen/gotable/util"
	"reflect"
	"strings"
)

const (
	Center = table.C
	Left = table.L
	Right = table.R
	Default = table.Default
)

// Deprecated
func CreateTable(header []string) (*table.Table, error) {
	util.DeprecatedTips("CreateTable", "Create", "3.0", "function")
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

func CreateByStruct(v interface{}) (*table.Table, error) {
	set := &table.Set{}
	s := reflect.TypeOf(v).Elem()
	for i := 0; i < s.NumField(); i++ {
		field := s.Field(i)
		name := field.Tag.Get("gotable")
		if name == "" {
			name = field.Name
		}

		err := set.Add(name)
		if err != nil {
			return nil, err
		}
	}
	tb := table.CreateTable(set)
	return tb, nil
}

func Version() string {
	return "gotable " + strings.Join(constant.GetVersions(), ".")
}

func Versions() []string { return constant.GetVersions() }
