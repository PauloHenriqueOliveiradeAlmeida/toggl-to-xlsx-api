package Tests

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"toggl-xlsx-back/src/Application/Services/Track"
	"toggl-xlsx-back/src/Application/UseCases/Workspace"
	"toggl-xlsx-back/src/Domain/Entities"
	"toggl-xlsx-back/src/Tests/Mocks"
)

func TestGetWorkspacesUseCaseShouldBeSuccess(test *testing.T) {
	trackService := Track.NewTrackService(Mocks.NewTrackMock([]Entities.WorkspaceEntity{
		{Id: 1, OrganizationId: 1, Name: "Workspace 1"},
		{Id: 2, OrganizationId: 2, Name: "Workspace 2"},
	}, nil))

	useCase := Workspace.NewGetWorkspacesUseCase(trackService)
	workspaces, error := useCase.Execute("email@email.com", "any_password001")

	assert.Nil(test, error)
	assert.Equal(test, 2, len(workspaces))
	assert.Equal(test, "Workspace 1", workspaces[0].Name)
	assert.Equal(test, "Workspace 2", workspaces[1].Name)
}

func TestGetWorkspacesUseCaseShouldBeFailure(test *testing.T) {
	trackService := Track.NewTrackService(Mocks.NewTrackMock([]Entities.WorkspaceEntity{}, errors.New("Serviço de track indisponível")))
	useCase := Workspace.NewGetWorkspacesUseCase(trackService)
	workspaces, error := useCase.Execute("email@email.com", "any_password001")

	assert.NotNil(test, error)
	assert.Equal(test, "Serviço de track indisponível", error.Error())
	assert.Nil(test, workspaces)
}
