package exception

type ColumnsLengthError struct {
	*baseError
}

func ColumnsLength() *ColumnsLengthError {
	err := &ColumnsLengthError{createBaseError("columns length must more than zero")}
	return err
}
