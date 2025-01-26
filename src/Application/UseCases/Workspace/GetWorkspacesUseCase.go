package Workspace

import (
	"toggl-xlsx-back/src/Application/Errors"
	"toggl-xlsx-back/src/Application/Services/Track"
	"toggl-xlsx-back/src/Domain/Entities"
)

type GetWorkspacesUseCase struct {
	trackService *Track.TrackService
}

func NewGetWorkspacesUseCase(trackService *Track.TrackService) *GetWorkspacesUseCase {
	return &GetWorkspacesUseCase{trackService: trackService}
}

func (this *GetWorkspacesUseCase) Execute(email string, password string) ([]Entities.WorkspaceEntity, error) {
	this.trackService.SetCredentials(email, password)
	workspaces, error := this.trackService.GetWorkspaces()

	if error != nil {
		return nil, &Errors.ServiceUnavailable{Message: "Serviço de track indisponível"}
	}

	return workspaces, nil

}
