package gotable

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/liushuochen/gotable/constant"
	"github.com/liushuochen/gotable/exception"
	"github.com/liushuochen/gotable/table"
	"github.com/liushuochen/gotable/util"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

const (
	Center = table.C
	Left = table.L
	Right = table.R
	Default = table.Default
)

// Create an empty table. When duplicate values in columns, table creation fails.
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
		return nil, exception.FileDoNotExist(path)
	}
	if !util.IsCSVFile(path) {
		return nil, exception.NotARegularCSVFile(path)
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil{
		return nil, err
	}
	if len(lines) < 1 {
		return nil, fmt.Errorf("csv file %s is empty", path)
	}

	tb, err := Create(lines[0]...)
	if err != nil {
		return nil, err
	}

	rows := make([]map[string]string, 0)
	for _, line := range lines[1:] {
		row := make(map[string]string)
		for i := range line {
			row[lines[0][i]] = line[i]
		}
		rows = append(rows, row)
	}
	tb.AddRows(rows)
	return tb, nil
}

func ReadFromJSONFile(path string) (*table.Table, error) {
	if !util.IsFile(path) {
		return nil, exception.FileDoNotExist(path)
	}
	if !util.IsJsonFile(path) {
		return nil, exception.NotARegularJSONFile(path)
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	rows := make([]map[string]string, 0)
	err = json.Unmarshal(byteValue, &rows)
	if err != nil {
		return nil, exception.NotGotableJSONFormat(path)
	}

	if len(rows) == 0 { return Create() }
	columns := make([]string, 0)
	for column := range rows[0] {
		columns = append(columns, column)
	}
	tb, err := Create(columns...)
	if err != nil {
		return nil, err
	}
	tb.AddRows(rows)
	return tb, nil
}
