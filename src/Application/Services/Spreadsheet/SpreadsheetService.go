package Spreadsheet

import "bytes"

type SpreadsheetService[T interface{}] struct {
	spreadsheet ISpreadsheet[T]
}

func NewSpreadsheetService[T interface{}](spreadsheet ISpreadsheet[T]) *SpreadsheetService[T] {
	return &SpreadsheetService[T]{spreadsheet: spreadsheet}
}

func (this *SpreadsheetService[T]) ConvertManyToSpreadsheet(fileName string, headers []string, data []T) (*bytes.Buffer, error) {
	file, error := this.spreadsheet.ConvertManyToSpreadsheet(fileName, headers, data)
	if error != nil {
		return nil, error
	}

	return file, error
}
