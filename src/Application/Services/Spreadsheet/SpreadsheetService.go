package Spreadsheet

import "bytes"

type SpreadsheetService struct {
	spreadsheet ISpreadsheet
}

func NewSpreadsheetService(spreadsheet ISpreadsheet) *SpreadsheetService {
	return &SpreadsheetService{spreadsheet: spreadsheet}
}

func (this *SpreadsheetService) ConvertManyToSpreadsheet(fileName string, headers []string, data []map[string]string) (*bytes.Buffer, error) {
	file, error := this.spreadsheet.ConvertManyToSpreadsheet(fileName, headers, data)
	if error != nil {
		return nil, error
	}

	return file, error
}
