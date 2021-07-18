package gotable

import (
	"fmt"
	"github.com/liushuochen/gotable/constant"
	"github.com/liushuochen/gotable/table"
	"github.com/liushuochen/gotable/util"
	"io/ioutil"
	"reflect"
	"strings"
)

const (
	Center = table.C
	Left = table.L
	Right = table.R
	Default = table.Default
)

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

func ReadFromCSVFile(path string) (*table.Table, error) {
	if !util.IsFile(path) {
		return nil, fmt.Errorf("csv file `%s` do not exist", path)
	}
	if !util.IsCSVFile(path) {
		return nil, fmt.Errorf("not a regular csv file: %s", path)
	}

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")
	if len(lines) < 1 {
		return nil, fmt.Errorf("csv file %s is empty", path)
	}

	columns := strings.Split(lines[0], ",")
	tb, err := Create(columns...)
	if err != nil {
		return nil, err
	}

	rows := make([]map[string]string, 0)
	for _, line := range lines[1:] {
		values := strings.Split(line, ",")
		row := make(map[string]string)
		for i := range values {
			row[columns[i]] = values[i]
		}
		rows = append(rows, row)
	}
	tb.AddRows(rows)
	return tb, nil
}
