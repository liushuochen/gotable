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
