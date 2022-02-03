package table

import (
	"sync"
)

type SafeTable struct {
	*base
	Row []sync.Map
}

func CreateSafeTable(set *Set) *SafeTable {
	return &SafeTable{
		base: createTableBase(set, SafeTableType, true),
		Row:  make([]sync.Map, 0),
	}
}

// Clear the table. The table is cleared of all data.
func (s *SafeTable) Clear() {
	s.Columns.Clear()
	s.Row = make([]sync.Map, 0)
}
