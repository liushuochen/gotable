// Package gotable_test used to test package gotable
package gotable_test

import (
	"github.com/liushuochen/gotable"
	"strings"
	"testing"
)

// TestVersionPrefix used to test version whether start with "gotable".
func TestVersionPrefix(t *testing.T) {
	version := gotable.Version()
	if !strings.HasPrefix(version, "gotable") {
		t.Errorf("expected version start switch gotable, but %s got", version)
	}
}
