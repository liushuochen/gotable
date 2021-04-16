package table

import (
	"fmt"
	"github.com/liushuochen/gotable/header"
)

func printGroup(
	group []map[string]string,
	header []*header.Header,
	columnMaxLen map[string]int,
	setBorder bool,
) {
	for _, item := range group {
		for index, head := range header {
			itemLen := columnMaxLen[head.Name] + 4
			s := ""
			if item[head.Name] == "-" {
				if setBorder {
					s, _ = center(item[head.Name], itemLen, "-")
				}
			} else {
				s, _ = center(item[head.Name], itemLen, " ")
			}

			icon := "|"
			if item[head.Name] == "-" {
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

func center(str string, length int, fillchar string) (string, error) {
	if len(fillchar) != 1 {
		err := fmt.Errorf("the fill character must be exactly one" +
			" character long")
		return "", err
	}

	if len(str) >= length {
		return str, nil
	}

	result := ""
	if isEvenNumber(length - len(str)) {
		front := ""
		for i := 0; i < ((length - len(str)) / 2); i++ {
			front = front + fillchar
		}

		result = front + str + front
	} else {
		front := ""
		for i := 0; i < ((length - len(str) - 1) / 2); i++ {
			front = front + fillchar
		}

		behind := front + fillchar
		result = front + str + behind
	}
	return result, nil
}

func isEvenNumber(number int) bool {
	if number%2 == 0 {
		return true
	}
	return false
}
