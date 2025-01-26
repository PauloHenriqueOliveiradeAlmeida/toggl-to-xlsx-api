package Tests

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
	"toggl-xlsx-back/src/Domain/Entities"
	"toggl-xlsx-back/src/Infra/Track"
	"toggl-xlsx-back/src/Tests/Mocks"
)

func TestSuccessOnMakeTogglWorkspaceIntegration(test *testing.T) {
	workspaces := []Entities.WorkspaceEntity{
		{Id: 1, OrganizationId: 1, Name: "Workspace 1"},
		{Id: 2, OrganizationId: 2, Name: "Workspace 2"},
	}

	payload, _ := json.Marshal(workspaces)
	client := &http.Client{
		Transport: &Mocks.HttpTransportMock{
			Response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBuffer(payload)),
			},
		},
	}

	togglClient := Track.NewTogglClient(client)
	togglClient.SetCredentials("email@email.com", "any_password001")
	response, error := togglClient.GetWorkspaces()

	assert.Nil(test, error)
	assert.Equal(test, response, workspaces)
}

func TestErrorOnMakeTogglWorkspaceIntegration(test *testing.T) {
	client := &http.Client{
		Transport: &Mocks.HttpTransportMock{
			Response: &http.Response{
				StatusCode: http.StatusUnauthorized,
			},
		},
	}

	togglClient := Track.NewTogglClient(client)
	togglClient.SetCredentials("email@email.com", "any_password001")
	response, error := togglClient.GetWorkspaces()

	assert.NotNil(test, error)
	assert.Contains(test, error.Error(), "Unauthorized")
	assert.Nil(test, response)
}

func TestSuccessOnMakeTogglProjectIntegration(test *testing.T) {
	projects := []Entities.ProjectEntity{
		{Id: 1, WorkspaceId: 1, Name: "Project 1", ActualHours: 1, ActualSeconds: 1},
		{Id: 2, WorkspaceId: 2, Name: "Project 2", ActualHours: 2, ActualSeconds: 2},
	}

	payload, _ := json.Marshal(projects)
	client := &http.Client{
		Transport: &Mocks.HttpTransportMock{
			Response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBuffer(payload)),
			},
		},
	}

	togglClient := Track.NewTogglClient(client)
	togglClient.SetCredentials("email@email.com", "any_password001")
	response, error := togglClient.GetProjects(123)

	assert.Nil(test, error)
	assert.Equal(test, response, projects)
}

func TestErrorOnMakeTogglProjectIntegration(test *testing.T) {
	client := &http.Client{
		Transport: &Mocks.HttpTransportMock{
			Response: &http.Response{
				StatusCode: http.StatusUnauthorized,
			},
		},
	}

	togglClient := Track.NewTogglClient(client)
	togglClient.SetCredentials("email@email.com", "any_password001")
	response, error := togglClient.GetProjects(123)

	assert.NotNil(test, error)
	assert.Contains(test, error.Error(), "Unauthorized")
	assert.Nil(test, response)
}
