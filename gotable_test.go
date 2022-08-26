// Package gotable_test used to test package gotable
package gotable_test

import (
	"github.com/liushuochen/gotable/exception"
	"strings"
	"testing"

	"github.com/liushuochen/gotable"
)

// Check version string whether start with "gotable".
func TestVersionPrefix(t *testing.T) {
	version := gotable.Version()
	if !strings.HasPrefix(version, "gotable") {
		t.Errorf("expected version start switch gotable, but %s got", version)
	}
}

// Test the value of gotable.TerminalDefault, gotable.Highlight, gotable.Underline and gotable.Flash.
func TestValueOfColorDisplay(t *testing.T) {
	if gotable.TerminalDefault != 0 {
		t.Errorf("expected gotable.TerminalDefault is 0, but %d got", gotable.TerminalDefault)
	}

	if gotable.Highlight != 1 {
		t.Errorf("expected gotable.Highlight is 1, but %d got", gotable.Highlight)
	}

	if gotable.Underline != 4 {
		t.Errorf("expected gotable.Underline is 4, but %d got", gotable.Underline)
	}

	if gotable.Flash != 5 {
		t.Errorf("expected gotable.Flash is 5, but %d got", gotable.Flash)
	}
}

// The value of color control
func TestValueOfColorControllers(t *testing.T) {
	if gotable.Black != 30 {
		t.Errorf("expected gotable.Black is 30, but %d got", gotable.Black)
	}

	if gotable.Red != 31 {
		t.Errorf("expected gotable.Red is 31, but %d got", gotable.Red)
	}

	if gotable.Green != 32 {
		t.Errorf("expected gotable.Green is 32, but %d got", gotable.Green)
	}

	if gotable.Yellow != 33 {
		t.Errorf("expected gotable.Yellow is 33, but %d got", gotable.Yellow)
	}

	if gotable.Blue != 34 {
		t.Errorf("expected gotable.Blue is 34, but %d got", gotable.Blue)
	}

	if gotable.Purple != 35 {
		t.Errorf("expected gotable.Purple is 35, but %d got", gotable.Purple)
	}

	if gotable.Cyan != 36 {
		t.Errorf("expected gotable.Cyan is 36, but %d got", gotable.Cyan)
	}

	if gotable.Write != 37 {
		t.Errorf("expected gotable.Write is 37, but %d got", gotable.Write)
	}

	if gotable.NoneBackground != 0 {
		t.Errorf("expected gotable.NoneBackground is 0, but %d got", gotable.NoneBackground)
	}
}

// Create a simple table.
func TestCreateSimpleTable(t *testing.T) {
	columns := []string{"country", "capital"}
	_, err := gotable.Create(columns...)
	if err != nil {
		t.Errorf("expected err is nil, but %s got", err.Error())
		return
	}
}

// Create a simple table with duplicate columns.
func TestCreateSimpleTableWithDuplicateColumn(t *testing.T) {
	columns := []string{"name", "name"}
	_, err := gotable.Create(columns...)
	if err == nil {
		t.Errorf("expected err is not nil, but nil got")
	}
}

// Create a simple table without column.
func TestCreateSimpleTableWithoutColumn(t *testing.T) {
	_, err := gotable.Create()
	switch err.(type) {
	case *exception.ColumnsLengthError:
	default:
		t.Errorf("expected err is ColumnsLengthError, but %T got", err)
	}
}

// Create a safe table.
func TestCreateSafeTable(t *testing.T) {
	columns := []string{"country", "capital"}
	_, err := gotable.CreateSafeTable(columns...)
	if err != nil {
		t.Errorf("expected err is nil, but %s got", err.Error())
		return
	}
}

// Create a safe table with duplicate columns.
func TestCreateSafeTableWithDuplicateColumn(t *testing.T) {
	columns := []string{"name", "name"}
	_, err := gotable.CreateSafeTable(columns...)
	if err == nil {
		t.Errorf("expected err is not nil, but nil got")
	}
}

// Create a safe table without column.
func TestCreateSafeTableWithoutColumn(t *testing.T) {
	_, err := gotable.CreateSafeTable()
	switch err.(type) {
	case *exception.ColumnsLengthError:
	default:
		t.Errorf("expected err is ColumnsLengthError, but %T got", err)
	}
}

// Create table using struct.
func TestCreateTableByStruct(t *testing.T) {
	type Student struct {
		Name string `gotable:"name"`
		Age  string `gotable:"age"`
	}

	stu := &Student{
		Name: "Bob",
		Age:  "12",
	}

	_, err := gotable.CreateByStruct(stu)
	if err != nil {
		t.Errorf("expected err is nil, but %s got.", err.Error())
	}
}

// Create table using empty struct.
func TestCreateTableByEmptyStruct(t *testing.T) {
	type Student struct{}

	stu := new(Student)
	_, err := gotable.CreateByStruct(stu)
	switch err.(type) {
	case *exception.ColumnsLengthError:
	default:
		t.Errorf("expected err is ColumnsLengthError, but %T got", err)
	}
}
