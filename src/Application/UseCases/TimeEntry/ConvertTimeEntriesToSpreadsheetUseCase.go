package TimeEntry

import (
	"bytes"
	"time"
	"toggl-xlsx-back/src/Application/Errors"
	"toggl-xlsx-back/src/Application/Services/Spreadsheet"
	"toggl-xlsx-back/src/Application/Services/Track"
	"toggl-xlsx-back/src/Domain/Entities"
)

type ConvertTimeEntriesToSpreadsheetUseCase struct {
	trackService       *Track.TrackService
	spreadSheetService *Spreadsheet.SpreadsheetService[Entities.TimeEntryEntity]
}

func NewConvertTimeEntriesToSpreadsheetUseCase(trackService *Track.TrackService, spreadSheetService *Spreadsheet.SpreadsheetService[Entities.TimeEntryEntity]) *ConvertTimeEntriesToSpreadsheetUseCase {
	return &ConvertTimeEntriesToSpreadsheetUseCase{trackService: trackService, spreadSheetService: spreadSheetService}
}

func (this *ConvertTimeEntriesToSpreadsheetUseCase) Execute(email string, password string, projectId int, startDate time.Time, endDate time.Time) (*bytes.Buffer, error) {
	this.trackService.SetCredentials(email, password)
	timeEntries, error := this.trackService.GetTimeEntriesByProjectId(projectId, startDate, endDate)
	if error != nil {
		return nil, &Errors.ServiceUnavailable{Message: "Serviço de track indisponível"}
	}

	spreadSheet, error := this.spreadSheetService.ConvertManyToSpreadsheet(time.Now().Month().String(), []string{"Date", "Duration", "Description"}, timeEntries)
	if error != nil {
		return nil, error
	}

	return spreadSheet, nil
}
