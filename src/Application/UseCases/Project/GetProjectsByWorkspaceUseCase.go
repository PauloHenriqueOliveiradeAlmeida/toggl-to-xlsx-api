package Project

import (
	"toggl-xlsx-back/src/Application/Errors"
	"toggl-xlsx-back/src/Application/Services/Track"
	"toggl-xlsx-back/src/Domain/Entities"
)

type GetProjectsUseCase struct {
	trackService *Track.TrackService
}

func NewGetProjectsUseCase(trackService *Track.TrackService) *GetProjectsUseCase {
	return &GetProjectsUseCase{trackService: trackService}
}

func (this *GetProjectsUseCase) Execute(email string, password string, workspaceId int) ([]Entities.ProjectEntity, error) {
	this.trackService.SetCredentials(email, password)
	projects, error := this.trackService.GetProjectsByWorkspaceId(workspaceId)

	if error != nil {
		return nil, &Errors.ServiceUnavailable{Message: "Serviço de track indisponível"}
	}

	return projects, nil

}
