package Track

import (
	"time"
	"toggl-xlsx-back/src/Domain/Entities"
)

type TrackService struct {
	track ITrack
}

func NewTrackService(track ITrack) *TrackService {
	return &TrackService{
		track: track,
	}
}

func (this *TrackService) SetCredentials(email string, password string) {
	this.track.SetCredentials(email, password)
}

func (this *TrackService) GetWorkspaces() ([]Entities.WorkspaceEntity, error) {
	workspaces, error := this.track.GetWorkspaces()

	if error != nil {
		return nil, error
	}

	return workspaces, nil
}

func (this *TrackService) GetProjectsByWorkspaceId(workspaceId int) ([]Entities.ProjectEntity, error) {
	projects, error := this.track.GetProjects(workspaceId)

	if error != nil {
		return nil, error
	}

	return projects, nil
}

func (this *TrackService) GetTimeEntriesByProjectId(projectId int, startDate time.Time, endDate time.Time) ([]Entities.TimeEntryEntity, error) {
	timeEntries, error := this.track.GetTimeEntries(startDate, endDate)
	if error != nil {
		return nil, error
	}

	var timeEntriesByProjectId []Entities.TimeEntryEntity
	for _, timeEntry := range timeEntries {
		if timeEntry.IsFromProject(projectId) {
			timeEntriesByProjectId = append(timeEntriesByProjectId, timeEntry)
		}
	}

	return timeEntries, nil
}
