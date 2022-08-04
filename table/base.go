// Package table define all table types methods.
// base.go contains basic methods of table types.
package table

// base struct contains common attributes to the table.
// Columns: Table columns
// border: Control the table border display(true: print table border).
// tableType: Use to record table types
// End: Used to set the ending. The default is newline "\n".
type base struct {
	partLen          int
	Columns          []*Set
	ColumnMaxLengths []map[string]int
	border           int8
	tableType        string
	End              string
}

func createTableBase(columns *Set, tableType string, border int8) *base {
	b := new(base)
	b.partLen = 1
	b.Columns = append(b.Columns, columns)
	b.ColumnMaxLengths = append(b.ColumnMaxLengths, make(map[string]int))
	b.tableType = tableType
	b.border = border
	b.End = "\n"
	return b
}

func (b *base) addTableBase(columns *Set) error {
	b.Columns = append(b.Columns, columns)
	b.ColumnMaxLengths = append(b.ColumnMaxLengths, make(map[string]int))
	b.partLen++
	return nil
}

// Type method returns a table type string.
func (b *base) Type() string {
	return b.tableType
}

// SetDefault method used to set default value for a given column name.
func (b *base) SetDefault(partNumber int, column string, defaultValue string) {
	for _, head := range b.Columns[partNumber].base {
		if head.Original() == column {
			head.SetDefault(defaultValue)
			break
		}
	}
}

// IsSimpleTable method returns a bool value indicate the table type is simpleTableType.
func (b *base) IsSimpleTable() bool {
	return b.tableType == simpleTableType
}

// IsSafeTable method returns a bool value indicate the table type is safeTableType.
func (b *base) IsSafeTable() bool {
	return b.tableType == safeTableType
}

// GetDefault method returns default value with a designated column name.
func (b *base) GetDefault(partNumber int, column string) string {
	for _, col := range b.Columns[partNumber].base {
		if col.Original() == column {
			return col.Default()
		}
	}
	return ""
}

// DropDefault method used to delete default value for designated column.
func (b *base) DropDefault(partNumber int, column string) {
	b.SetDefault(partNumber, column, "")
}

// GetDefaults method return a map that contains all default value of each columns.
// * map[column name] = default value
func (b *base) GetDefaults(partNumber int) map[string]string {
	defaults := make(map[string]string)
	if partNumber <= b.partLen {
		for _, column := range b.Columns[partNumber].base {
			defaults[column.Original()] = column.Default()
		}
		return defaults
	}
	return nil
}
