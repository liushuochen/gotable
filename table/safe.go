package table

import (
	"github.com/liushuochen/gotable/cell"
	"github.com/liushuochen/gotable/exception"
	"sync"
)

type SafeTable struct {
	*base
	Row []sync.Map
}

// CreateSafeTable returns a pointer of SafeTable.
func CreateSafeTable(set *Set) *SafeTable {
	return &SafeTable{
		base: createTableBase(set, safeTableType, true),
		Row:  make([]sync.Map, 0),
	}
}

// Clear the table. The table is cleared of all data.
func (s *SafeTable) Clear() {
	s.Columns.Clear()
	s.Row = make([]sync.Map, 0)
}

// AddColumn method used to add a new column for table. It returns an error when column has been exist.
func (s *SafeTable) AddColumn(column string) error {
	err := s.Columns.Add(column)
	if err != nil {
		return err
	}

	// Modify exist value, add new column.
	for _, row := range s.Row {
		row.Store(column, cell.CreateEmptyData())
	}
	return nil
}

// AddRow method support Map and Slice argument.
// For Map argument, you must put the data from each row into a Map and use column-data as key-value pairs. If the Map
//   does not contain a column, the table sets it to the default value. If the Map contains a column that does not
//   exist, the AddRow method returns an error.
// For Slice argument, you must ensure that the slice length is equal to the column length. Method will automatically
//   mapping values in Slice and columns. The default value cannot be omitted and must use gotable.Default constant.
// Return error types:
//   - *exception.UnsupportedRowTypeError: It returned when the type of the argument is not supported.
//   - *exception.RowLengthNotEqualColumnsError: It returned if the argument is type of the Slice but the length is
//       different from the length of column.
//   - *exception.ColumnDoNotExistError: It returned if the argument is type of the Map but contains a nonexistent
//       column as a key.
func (s *SafeTable) AddRow(row interface{}) error {
	switch v := row.(type) {
	case []string:
		return s.addRowFromSlice(v)
	case map[string]string:
		return s.addRowFromMap(v)
	default:
		return exception.UnsupportedRowType(v)
	}
}

// AddRows used to add a slice of rows map. It returns a slice of map which add failed.
func (s *SafeTable) AddRows(rows []map[string]string) []map[string]string {
	failure := make([]map[string]string, 0)
	for _, row := range rows {
		err := s.AddRow(row)
		if err != nil {
			failure = append(failure, row)
		}
	}
	return failure
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

func (s *SafeTable) addRowFromSlice(row []string) error {
	rowLength := len(row)
	if rowLength != s.Columns.Len() {
		return exception.RowLengthNotEqualColumns(rowLength, s.Columns.Len())
	}

	rowMap := make(map[string]string, 0)
	for i := 0; i < rowLength; i++ {
		if row[i] == Default {
			rowMap[s.Columns.base[i].Original()] = s.Columns.base[i].Default()
		} else {
			rowMap[s.Columns.base[i].Original()] = row[i]
		}
	}

	s.Row = append(s.Row, toSafeRow(rowMap))
	return nil
}
