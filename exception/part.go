package exception

import "fmt"

type PartNumberError struct {
	*baseError
}

func PartNumber(partLen int) *PartNumberError {
	err := &PartNumberError{createBaseError("Part length now is " + fmt.Sprint(partLen))}
	return err
}
