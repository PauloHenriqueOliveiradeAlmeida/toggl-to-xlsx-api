package Track

import (
	"time"
	"toggl-xlsx-back/src/Domain/Entities"
)

type ITrack interface {
	SetCredentials(email string, password string)
	GetWorkspaces() ([]Entities.WorkspaceEntity, error)
	GetProjects(workspaceId int) ([]Entities.ProjectEntity, error)
	GetTimeEntries(startDate time.Time, endDate time.Time) ([]Entities.TimeEntryEntity, error)
}
