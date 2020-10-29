package gotable

import (
	"reflect"
	"testing"

	"github.com/TuoAiTang/gotable/color"
	"github.com/TuoAiTang/gotable/table"
	"github.com/stretchr/testify/assert"
)

var (
	testHeader = []string{"Name", "Experience", "Salary"}
	testData = []map[string]table.Sequence{
		{
			"Name":       table.DefaultSequence("Alice"),
			"Experience": table.DefaultSequence("Three year."),
			"Salary":     table.DefaultSequence("2300.00"),
		},
		{
			"Name":       table.DefaultSequence("Bob"),
			"Experience": table.DefaultSequence("Ten year."),
			"Salary":     table.DefaultSequence("900.00"),
		},
		{
			"Name":       table.DefaultSequence("Coco"),
			"Experience": table.DefaultSequence("One year."),
			"Salary":     table.DefaultSequence("9000.00"),
		},
	}
	testDataWithColor = []map[string]table.Sequence{
		{
			"Name":       color.ColorfulString(color.CYAN,"Alice"),
			"Experience": color.ColorfulString(color.CYAN,"Three year."),
			"Salary":     color.ColorfulString(color.CYAN,"2300.00"),
		},
		{
			"Name":       color.ColorfulString(color.BLUE,"Bob"),
			"Experience": color.ColorfulString(color.BLUE,"Ten year."),
			"Salary":     color.ColorfulString(color.BLUE,"900.00"),
		},
		{
			"Name":       color.ColorfulString(color.CYAN,"Coco"),
			"Experience": color.ColorfulString(color.CYAN,"One year."),
			"Salary":     color.ColorfulString(color.CYAN,"9000.00"),
		},
	}
)

func TestTable(t *testing.T) {
	tbl, err := CreateTable(testHeader)
	assert.Nil(t, err)

	for _, value := range testData {
		err := tbl.AddValue(value)
		assert.Nil(t, err)
	}

	tbl.PrintTable()
}

func TestColorTable(t *testing.T) {
	tbl, err := CreateTable(testHeader)
	assert.Nil(t, err)

	for _, value := range testDataWithColor {
		err := tbl.AddValue(value)
		assert.Nil(t, err)
	}

	tbl.PrintTable()
}

func TestColorTableStruct(t *testing.T) {
	type Sal struct {
		Name string
		Salary string
		Experience string
	}

	array := make([]interface{}, 0)
	for i := 0; i < 10; i++ {
		array = append(array, Sal{
			Name:       "name",
			Salary:     "sa",
			Experience: "exp",
		})
	}

	err := PrintSlice(array, table.DefaultController)
	assert.Nil(t, err)

	type Libra struct {
		Name string
		Sex bool
		Rank int
	}

	array = make([]interface{}, 0)
	for i := 0; i < 10; i++ {
		array = append(array, Libra{
			"name",
			i % 2 == 0,
			i*3,
		})
	}

	err = PrintSlice(array, func(field string, val reflect.Value) color.Color {
		switch field {
		case "Rank":
			if val.Int() > 18 {
				return color.CYAN
			}
		case "Sex":
			if val.Bool() {
				return color.CYAN
			}
		}
		return ""
	})
	assert.Nil(t, err)
}