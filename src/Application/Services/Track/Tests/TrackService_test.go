package Tests

import (
	"errors"
	"testing"
	"toggl-xlsx-back/src/Application/Services/Track"
	"toggl-xlsx-back/src/Domain/Entities"
	"toggl-xlsx-back/src/Tests/Mocks"

	"github.com/stretchr/testify/assert"
)

func TestGetWorkspacesShouldBeSuccess(test *testing.T) {
	trackMock := Mocks.NewTrackMock([]Entities.WorkspaceEntity{
		{Id: 1, OrganizationId: 1, Name: "Workspace 1"},
		{Id: 2, OrganizationId: 2, Name: "Workspace 2"},
	}, nil)

	trackService := Track.NewTrackService(trackMock)
	workspaces, error := trackService.GetWorkspaces()

	assert.Nil(test, error)

	assert.Len(test, workspaces, 2)
	assert.Equal(test, "Workspace 1", workspaces[0].Name)
	assert.Equal(test, "Workspace 2", workspaces[1].Name)
}

func TestGetWorkspacesShouldBeFailure(test *testing.T) {
	trackMock := Mocks.NewTrackMock([]Entities.WorkspaceEntity{}, errors.New("Service unavailable"))

	trackService := Track.NewTrackService(trackMock)
	workspaces, error := trackService.GetWorkspaces()

	assert.NotNil(test, error)
	assert.Equal(test, "Service unavailable", error.Error())
	assert.Nil(test, workspaces)
}
