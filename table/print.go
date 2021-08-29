package table

import (
	"fmt"
	"github.com/liushuochen/gotable/cell"
	"github.com/liushuochen/gotable/column"
)

func printGroup(
	group []map[string]cell.Cell,
	header []*column.Column,
	columnMaxLen map[string]int,
	setBorder bool,
) {
	for _, item := range group {
		for index, head := range header {
			itemLen := columnMaxLen[head.Original()] + 2
			//fmt.Println(head, itemLen)
			s := ""
			if item[head.String()].String() == "-" {
				if setBorder {
					s, _ = center(item[head.String()], itemLen, "-")
				}
			} else {
				switch head.Align() {
				case R:
					s, _ = right(item[head.String()], itemLen, " ")
				case L:
					s, _ = left(item[head.String()], itemLen, " ")
				default:
					s, _ = center(item[head.String()], itemLen, " ")
				}
			}

			icon := "|"
			if item[head.String()].String() == "-" {
				icon = "+"
			}
			if !setBorder {
				icon = ""
			}

			if index == 0 {
				s = icon + s + icon
			} else {
				s = "" + s + icon
			}
			fmt.Print(s)
		}
		fmt.Println()
	}
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func center(c cell.Cell, length int, fillchar string) (string, error) {
	if len(fillchar) != 1 {
		err := fmt.Errorf("the fill character must be exactly one" +
			" character long")
		return "", err
	}

	if c.Length() >= length {
		return c.String(), nil
	}

	result := ""
	if isEvenNumber(length - c.Length()) {
		front := ""
		for i := 0; i < ((length - c.Length()) / 2); i++ {
			front = front + fillchar
		}

		result = front + c.String() + front
	} else {
		front := ""
		for i := 0; i < ((length - c.Length() - 1) / 2); i++ {
			front = front + fillchar
		}

		behind := front + fillchar
		result = front + c.String() + behind
	}
	return result, nil
}

func left(c cell.Cell, length int, fillchar string) (string, error) {
	if len(fillchar) != 1 {
		err := fmt.Errorf("the fill character must be exactly one" +
			" character long")
		return "", err
	}

	result := c.String() + block(length - c.Length())
	return result, nil
}

func right(c cell.Cell, length int, fillchar string) (string, error) {
	if len(fillchar) != 1 {
		err := fmt.Errorf("the fill character must be exactly one" +
			" character long")
		return "", err
	}

	result := block(length - c.Length()) + c.String()
	return result, nil
}

func block(length int) string {
	result := ""
	for i := 0; i < length; i++ {
		result += " "
	}
	return result
}

func isEvenNumber(number int) bool {
	if number % 2 == 0 {
		return true
	}
	return false
}

func toRow(value map[string]string) map[string]cell.Cell {
	row := make(map[string]cell.Cell)
	for k, v := range value {
		row[k] = cell.CreateData(v)
	}
	return row
}
