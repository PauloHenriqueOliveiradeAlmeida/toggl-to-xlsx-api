package Spreadsheet

import "bytes"

type ISpreadsheet[T interface{}] interface {
	ConvertManyToSpreadsheet(fileName string, headers []string, data []T) (*bytes.Buffer, error)
}
