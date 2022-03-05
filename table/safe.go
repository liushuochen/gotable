package table

import (
	"github.com/liushuochen/gotable/exception"
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

// AddRow method only support Map argument.
// For Map argument, you must put the data from each row into a Map and use column-data as key-value pairs. If the Map
//   does not contain a column, the table sets it to the default value. If the Map contains a column that does not
//   exist, the AddRow method returns an error.
// Return error types:
//   - *exception.UnsupportedRowTypeError: It returned when the type of the argument is not supported.
//   - *exception.ColumnDoNotExistError: It returned if the argument is type of the Map but contains a nonexistent
//       column as a key.
func (s *SafeTable) AddRow(row interface{}) error {
	switch v := row.(type) {
	//case []string:
	//	return tb.addRowFromSlice(v)
	case map[string]string:
		return s.addRowFromMap(v)
	default:
		return exception.UnsupportedRowType(v)
	}
}

func (s *SafeTable) addRowFromMap(row map[string]string) error {
	for key := range row {
		if !s.Columns.Exist(key) {
			return exception.ColumnDoNotExist(key)
		}

		// add row by const `DEFAULT`
		if row[key] == Default {
			row[key] = s.Columns.Get(key).Default()
		}
	}

	// Add default value
	for _, col := range s.Columns.base {
		_, ok := row[col.Original()]
		if !ok {
			row[col.Original()] = col.Default()
		}
	}

	s.Row = append(s.Row, toSafeRow(row))
	return nil
}
