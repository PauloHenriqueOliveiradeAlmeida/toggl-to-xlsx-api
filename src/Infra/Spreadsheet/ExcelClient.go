package Spreadsheet

import (
	"bytes"
	"github.com/xuri/excelize/v2"
	"strconv"
)

type ExcelClient struct{}

func NewExcelClient() *ExcelClient {
	return &ExcelClient{}
}

func (this *ExcelClient) ConvertManyToSpreadsheet(sheetName string, headers []string, data []map[string]string) (*bytes.Buffer, error) {

	file := excelize.NewFile()
	file.SetSheetName(file.GetSheetName(0), sheetName)
	err := file.SetSheetRow(sheetName, "A1", &headers)
	if err != nil {
		return nil, err
	}

	for index, row := range data {
		for columnIndex, cell := range headers {
			value := row[cell]
			if value == "" {
				value = "-"
			}
			err := file.SetCellValue(sheetName, string(rune(65+columnIndex))+strconv.Itoa(index+2), value)
			if err != nil {
				return nil, err
			}
			columnWidth, error := file.GetColWidth(sheetName, string(rune(65+columnIndex)))
			if error != nil {
				return nil, error
			}
			if int(columnWidth) < len(value) {
				file.SetColWidth(sheetName, string(rune(65+columnIndex)), string(rune(65+columnIndex)), float64(len(value)+5))
			}
		}
	}

	buffer, error := file.WriteToBuffer()
	if error != nil {
		return nil, error
	}

	return buffer, nil
}
