package Mocks

import (
	"time"
	"toggl-xlsx-back/src/Domain/Entities"
)

type TrackMock[Response interface{}] struct {
	response Response
	error    error
}

func NewTrackMock[Response any](response Response, error error) *TrackMock[Response] {
	return &TrackMock[Response]{
		response: response,
		error:    error,
	}
}

func (this *TrackMock[Response]) SetCredentials(email string, password string) {}

func (this *TrackMock[Response]) GetWorkspaces() ([]Entities.WorkspaceEntity, error) {
	workspaces, _ := any(this.response).([]Entities.WorkspaceEntity)
	return workspaces, this.error
}

func (this *TrackMock[Response]) GetProjects(workspaceId int) ([]Entities.ProjectEntity, error) {
	projects, _ := any(this.response).([]Entities.ProjectEntity)
	return projects, this.error
}

func (this *TrackMock[Response]) GetProject(workspaceId int, projectId int) (Entities.ProjectEntity, error) {
	project, _ := any(this.response).(Entities.ProjectEntity)
	return project, this.error
}

func (this *TrackMock[Response]) GetTimeEntries(startDate time.Time, endDate time.Time) ([]Entities.TimeEntryEntity, error) {
	timeEntries, _ := any(this.response).([]Entities.TimeEntryEntity)
	return timeEntries, this.error
}
