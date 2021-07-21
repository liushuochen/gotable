package exception

import "fmt"

type NotARegularCSVFileError struct {
	*baseError
	filename	string
}

func NotARegularCSVFile(path string) *NotARegularCSVFileError {
	message := fmt.Sprintf("not a regular csv file: %s", path)
	err := &NotARegularCSVFileError{createBaseError(message), path}
	return err
}

func (err *NotARegularCSVFileError) Filename() string {
	return err.filename
}
