// Package gotable_test used to test package gotable
package gotable_test

import (
	"strings"
	"testing"

	"github.com/liushuochen/gotable"
)

// TestVersionPrefix used to test version whether start with "gotable".
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

// Test the value of color control
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

// Test creating a simple table.
func TestCreateTable(t *testing.T) {
	columns := []string{"country", "capital"}
	_, err := gotable.Create(columns...)
	if err != nil {
		t.Errorf("expected err is nil, but %s got", err.Error())
	}
}
