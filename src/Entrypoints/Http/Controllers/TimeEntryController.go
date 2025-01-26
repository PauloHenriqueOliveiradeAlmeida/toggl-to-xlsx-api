package Controllers

import (
	"net/http"
	"strconv"
	"time"
	"toggl-xlsx-back/src/Application/Errors"
	"toggl-xlsx-back/src/Application/UseCases/TimeEntry"

	"github.com/gin-gonic/gin"
)

type TimeEntryController struct {
	router                                 *gin.Engine
	convertTimeEntriesToSpreadsheetUseCase *TimeEntry.ConvertTimeEntriesToSpreadsheetUseCase
}

func NewTimeEntryController(router *gin.Engine, convertTimeEntryToSpreadsheetUseCase *TimeEntry.ConvertTimeEntriesToSpreadsheetUseCase) *TimeEntryController {
	controller := TimeEntryController{
		router:                                 router,
		convertTimeEntriesToSpreadsheetUseCase: convertTimeEntryToSpreadsheetUseCase,
	}

	controller.makeRoutes()
	return &controller
}

func (this *TimeEntryController) makeRoutes() {
	this.router.GET("/time-entries/:projectId/spreadsheet", this.ConvertTimeEntriesToSpreadsheet)
}

func (this *TimeEntryController) ConvertTimeEntriesToSpreadsheet(context *gin.Context) {
	projectId, haveProjectId := context.Params.Get("projectId")
	if !haveProjectId {
		context.Error(&Errors.BadRequest{Message: "projectId is required"})
		return
	}

	convertedProjectId, error := strconv.Atoi(projectId)
	if error != nil {
		context.Error(&Errors.BadRequest{Message: "projectId must be a number"})
		return
	}

	startDate := context.Query("start-date")
	endDate := context.Query("end-date")
	if startDate == "" || endDate == "" {
		context.Error(&Errors.BadRequest{Message: "startDate and endDate are required"})
		return
	}

	convertedStartDate, error := time.Parse("2006-01-02", startDate)
	if error != nil {
		context.Error(&Errors.BadRequest{Message: "startDate is a invalid date"})
		return
	}

	convertedEndDate, error := time.Parse("2006-01-02", endDate)
	if error != nil {
		context.Error(&Errors.BadRequest{Message: "endDate is a invalid date"})
		return
	}

	username, password, _ := context.Request.BasicAuth()
	spreadsheet, error := this.convertTimeEntriesToSpreadsheetUseCase.Execute(username, password, convertedProjectId, convertedStartDate, convertedEndDate)
	if error != nil {
		context.Error(error)
		return
	}

	context.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", spreadsheet.Bytes())
}
