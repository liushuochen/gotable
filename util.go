package gotable

import "fmt"

func printGroup(group []map[string]Sequence, header []string, columnMaxLen map[string]int) {
	for _, item := range group {
		for index, head := range header {
			itemLen := columnMaxLen[head] + 4
			s := ""
			if item[head].Val() == "-" {
				s, _ = center(item[head], itemLen, "-")
			} else {
				s, _ = center(item[head], itemLen, " ")
			}

			icon := "|"
			if item[head].Val() == "-" {
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
		return str.Val(), nil
	}

	result := ""
	if isEvenNumber(length - str.Len()) {
		front := ""
		for i := 0; i < ((length - str.Len()) / 2); i++ {
			front = front + fillchar
		}

		result = front + str.Val() + front
	} else {
		front := ""
		for i := 0; i < ((length - str.Len() - 1) / 2); i++ {
			front = front + fillchar
		}

		behind := front + fillchar
		result = front + str.Val() + behind
	}
	return result, nil
}

func isEvenNumber(number int) bool {
	if number%2 == 0 {
		return true
	}
	return false
}