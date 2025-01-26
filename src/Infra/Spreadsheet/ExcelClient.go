package Spreadsheet

import (
	"bytes"
	"strconv"

	"github.com/xuri/excelize/v2"
)

type ExcelClient[T interface{}] struct{}

func NewExcelClient[T interface{}]() *ExcelClient[T] {
	return &ExcelClient[T]{}
}

func (this *ExcelClient[T]) ConvertManyToSpreadsheet(sheetName string, headers []string, data []T) (*bytes.Buffer, error) {

	file := excelize.NewFile()
	file.SetSheetName(file.GetSheetName(0), sheetName)
	err := file.SetSheetRow(sheetName, "A1", &headers)
	if err != nil {
		return nil, err
	}

	for index, row := range data {
		for columnIndex, cell := range headers {
			err := file.SetCellValue(sheetName, strconv.Itoa(index+2)+string('A'+uint8(columnIndex)), cell)
			if err != nil {
				return nil, err
			}
		}
	}

	buffer, error := file.WriteToBuffer()
	if error != nil {
		return nil, error
	}

	return buffer, nil
}
