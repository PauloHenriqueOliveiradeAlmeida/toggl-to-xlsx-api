package Spreadsheet

import "bytes"

type ISpreadsheet interface {
	ConvertManyToSpreadsheet(fileName string, headers []string, data []map[string]string) (*bytes.Buffer, error)
}
