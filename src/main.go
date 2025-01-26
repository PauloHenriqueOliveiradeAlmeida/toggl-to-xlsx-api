package main

import (
	"net/http"
	"toggl-xlsx-back/src/Application/Services/Spreadsheet"
	"toggl-xlsx-back/src/Application/Services/Track"
	"toggl-xlsx-back/src/Application/UseCases/Project"
	"toggl-xlsx-back/src/Application/UseCases/TimeEntry"
	"toggl-xlsx-back/src/Application/UseCases/Workspace"
	"toggl-xlsx-back/src/Domain/Entities"
	"toggl-xlsx-back/src/Entrypoints/Http/Controllers"
	"toggl-xlsx-back/src/Entrypoints/Http/Middlewares"
	SpreadsheetImpl "toggl-xlsx-back/src/Infra/Spreadsheet"
	TrackImpl "toggl-xlsx-back/src/Infra/Track"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(Middlewares.ErrorHandler())
	router.GET("/ping", func(context *gin.Context) {
		context.IndentedJSON(http.StatusOK, "pong")
	})

	client := &http.Client{}
	Controllers.NewWorkspaceController(router, Workspace.NewGetWorkspacesUseCase(Track.NewTrackService(TrackImpl.NewTogglClient(client))))
	Controllers.NewProjectController(router, Project.NewGetProjectsUseCase(Track.NewTrackService(TrackImpl.NewTogglClient(client))))
	Controllers.NewTimeEntryController(router, TimeEntry.NewConvertTimeEntriesToSpreadsheetUseCase(
		Track.NewTrackService(TrackImpl.NewTogglClient(client)),
		Spreadsheet.NewSpreadsheetService(SpreadsheetImpl.NewExcelClient[Entities.TimeEntryEntity]()),
	))
	router.Run(":8000")
}
