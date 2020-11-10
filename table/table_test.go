package table

import (
	"reflect"
	"strings"
	"testing"

	"github.com/TuoAiTang/gotable/color"
	"github.com/stretchr/testify/assert"
)

type MyStruct struct {
	Name       string
	Experience int32
	Salary     float64
}

var (
	testHeader = []string{"Name", "Experience", "Salary"}
	plainData  = []map[string]Sequence{
		{
			"Name":       DefaultSequence("Alice"),
			"Experience": DefaultSequence("Three year."),
			"Salary":     DefaultSequence("2300.00"),
		},
		{
			"Name":       DefaultSequence("Bob"),
			"Experience": DefaultSequence("Ten year."),
			"Salary":     DefaultSequence("900.00"),
		},
		{
			"Name":       DefaultSequence("Coco"),
			"Experience": DefaultSequence("One year."),
			"Salary":     DefaultSequence("9000.00"),
		},
	}
	colorData = []map[string]Sequence{
		{
			"Name":       color.ColorfulString(color.CYAN, "Alice"),
			"Experience": color.ColorfulString(color.CYAN, "Three year."),
			"Salary":     color.ColorfulString(color.CYAN, "2300.00"),
		},
		{
			"Name":       color.ColorfulString(color.BLUE, "Bob"),
			"Experience": color.ColorfulString(color.BLUE, "Ten year."),
			"Salary":     color.ColorfulString(color.BLUE, "900.00"),
		},
		{
			"Name":       color.ColorfulString(color.CYAN, "Coco"),
			"Experience": color.ColorfulString(color.CYAN, "One year."),
			"Salary":     color.ColorfulString(color.CYAN, "9000.00"),
		},
	}
	structSliceData = []interface{}{
		MyStruct{
			Name:       "京东方A",
			Experience: 3,
			Salary:     2300.00,
		},
		MyStruct{
			Name:       "券商ETF",
			Experience: 10,
			Salary:     900.00,
		},
		MyStruct{
			Name:       "圆通",
			Experience: 1,
			Salary:     9000.00,
		},
	}
)

func TestTable(t *testing.T) {
	tbl, err := CreateTable(testHeader)
	assert.Nil(t, err)

	for _, value := range plainData {
		err := tbl.AddValue(value)
		assert.Nil(t, err)
	}

	tbl.PrintTable()
}

func TestColorTable(t *testing.T) {
	tbl, err := CreateTable(testHeader)
	assert.Nil(t, err)

	for _, value := range colorData {
		err := tbl.AddValue(value)
		assert.Nil(t, err)
	}

	tbl.PrintTable()
}

func TestStructTable(t *testing.T) {
	controller := func(field string, val reflect.Value) color.Color {
		switch field {
		case "Name":
			if strings.Contains(val.String(), "c") {
				return color.CYAN
			}
		case "Experience":
			if val.Int() > 5 {
				return color.MAGENTA
			}
		case "Salary":
			if val.Float() < 1000 {
				return color.YELLOW
			}
		}
		return ""
	}

	tbl, err := CreateTableFromStruct(MyStruct{}, WithColorController(controller))
	assert.Nil(t, err)

	err = tbl.AddValuesFromSlice(structSliceData)
	assert.Nil(t, err)

	tbl.PrintTable()
}
