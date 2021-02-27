package gotable

import (
	"fmt"
	"github.com/liushuochen/gotable/header"
)

func printGroup(group []map[string]Sequence, header []*header.Header, columnMaxLen map[string]int) {
	for _, item := range group {
		for index, head := range header {
			itemLen := columnMaxLen[head.Name] + 4
			s := ""
			if item[head.Name].Value() == "-" {
				s, _ = center(item[head.Name], itemLen, "-")
			} else {
				s, _ = center(item[head.Name], itemLen, " ")
			}

			icon := "|"
			if item[head.Name].Value() == "-" {
				icon = "+"
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

func center(str Sequence, length int, fillchar string) (string, error) {
	if len(fillchar) != 1 {
		err := fmt.Errorf("the fill character must be exactly one" +
			" character long")
		return "", err
	}

	if str.Len() >= length {
		return str.Value(), nil
	}

	result := ""
	if isEvenNumber(length - str.Len()) {
		front := ""
		for i := 0; i < ((length - str.Len()) / 2); i++ {
			front = front + fillchar
		}

		result = front + str.Value() + front
	} else {
		front := ""
		for i := 0; i < ((length - str.Len() - 1) / 2); i++ {
			front = front + fillchar
		}

		behind := front + fillchar
		result = front + str.Value() + behind
	}
	return result, nil
}

func isEvenNumber(number int) bool {
	if number%2 == 0 {
		return true
	}
	return false
}