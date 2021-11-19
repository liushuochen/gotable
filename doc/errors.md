# Gotable error types

In this section, we introduce the error types defined in gotable. By default, we will still return the original
```Error``` interface. The following code demonstrates how to properly handle errors returned by ```gotable```.

```go
package main

import (
	"fmt"
	"github.com/liushuochen/gotable"
	"github.com/liushuochen/gotable/exception"
)

func main() {
	table, err := gotable.ReadFromJSONFile("cmd/fun.csv")
	if err != nil {
		switch err.(type) {
		case *exception.FileDoNotExistError:
			exp, _ := err.(*exception.FileDoNotExistError)
			fmt.Printf("file %s dot exit: %s", exp.Filename(), err.Error())
		default:
			fmt.Println(err.Error())
		}
		return
	}
	fmt.Println(table)
}

```

[Return to the home page](../README.md)

## FileDoNotExistError
This error type indicates that the filename was not found in the server. It has a public method
```*FileDoNotExistError.Filename() string``` that returns the wrong filename.

## NotARegularCSVFileError
This error type indicates that the given filename is not a valid csv. It has a public method
```*NotARegularCSVFileError.Filename() string``` that returns the wrong CSV filename.

## NotARegularJSONFileError
This error type indicates that the given filename is not a valid JSON. It has a public method
```*NotARegularJSONFileError.Filename() string``` that returns the wrong JSON filename.

## NotGotableJSONFormatError
This error type indicates that the data format stored in the JSON file can not be parsed as a table.
It has a public method ```*NotGotableJSONFormatError.Filename() string``` that returns the wrong JSON filename.

## UnsupportedRowTypeError
This error type indicates that the row data structure is not support. It has a public method 
```*UnsupportedRowTypeError.Type() string``` that returns the wrong type name.

## ColumnsLengthError
This error type indicates that column's length not greater than 0.

## ColumnDoNotExistError
A nonexistent column was found while adding a row. It has a public method ```*ColumnDoNotExistError.Name() string``` 
that returns the nonexistent column name.

## RowLengthNotEqualColumnsError
This error is raised when adding a row from a Slice when the length of the Slice is not equal with the length of the 
table column.

## UnSupportedFileTypeError
When the file type read is not supported. It has a public mnethod ```*UnSupportedFileTypeError.Filename() string``` 
that returns the wrong filename.
