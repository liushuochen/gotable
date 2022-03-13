package table

// base struct contains common attributes to the table.
// Columns: Table columns
// border: Control the table border display(true: print table border).
// tableType: Use to record table types
// End: Used to set the ending. The default is newline "\n".
type base struct {
	Columns   *Set
	border    bool
	tableType string
	End       string
}

func createTableBase(columns *Set, tableType string, border bool) *base {
	b := new(base)
	b.Columns = columns
	b.tableType = tableType
	b.border = border
	b.End = "\n"
	return b
}

// Type method returns a table type string.
func (b *base) Type() string {
	return b.tableType
}
