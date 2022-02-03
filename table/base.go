package table

type base struct {
	Columns   *Set
	border    bool
	tableType string
}

func createTableBase(columns *Set, tableType string, border bool) *base {
	base := new(base)
	base.Columns = columns
	base.tableType = tableType
	base.border = border
	return base
}

// Type method returns a table type string.
func (b *base) Type() string {
	return b.tableType
}
