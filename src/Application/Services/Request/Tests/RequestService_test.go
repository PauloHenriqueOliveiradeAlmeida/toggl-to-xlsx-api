package Tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"strings"
	"testing"
	"toggl-xlsx-back/src/Application/Services/Request"
	"toggl-xlsx-back/src/Domain/Entities"
	"toggl-xlsx-back/src/Tests/Mocks"
)

func TestSendRequestPostWithBody(test *testing.T) {
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
	request := Request.NewRequest("/workspaces", Request.POST).WithBody(struct{ name string }{
		name: "Workspace 1",
	})

	response, error := Request.Send[[]Entities.WorkspaceEntity](client, request)
	assert.Nil(test, error)
	assert.NotNil(test, request)

	assert.Equal(test, workspaces, response)
}

func TestSendRequestGetWithBody(test *testing.T) {
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
	request := Request.NewRequest("/workspaces", Request.GET).WithBody(struct{ name string }{
		name: "Workspace 1",
	})

	response, error := Request.Send[[]Entities.WorkspaceEntity](client, request)
	assert.Nil(test, error)
	assert.NotNil(test, request)

	assert.Equal(test, workspaces, response)
}

func TestSendRequestWithHeaders(test *testing.T) {
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
	request := Request.NewRequest("/workspaces", Request.POST).WithHeaders(map[string]string{
		"Authorization": "Bearer token",
	})

	response, error := Request.Send[[]Entities.WorkspaceEntity](client, request)
	assert.Nil(test, error)
	assert.NotNil(test, request)

	assert.Equal(test, workspaces, response)
}

func TestErrorOnJSONSerialization(test *testing.T) {
	client := &http.Client{
		Transport: &Mocks.HttpTransportMock{
			Response: &http.Response{},
		},
	}
	type InvalidBody struct {
		Name chan int
	}
	request := Request.NewRequest("/workspaces", Request.POST).WithBody(InvalidBody{})

	_, error := Request.Send[[]Entities.WorkspaceEntity](client, request)

	assert.NotNil(test, error)
	assert.Equal(test, "unexpected end of JSON input", error.Error())
}

func TestErrorOnMakeRequest(test *testing.T) {
	client := &http.Client{
		Transport: &Mocks.HttpTransportMock{
			Response: &http.Response{},
		},
	}
	request := Request.NewRequest("/workspaces", "??")

	_, error := Request.Send[[]Entities.WorkspaceEntity](client, request)
	assert.NotNil(test, error)
	assert.Contains(test, error.Error(), "invalid method")
}

func TestErrorOnSendRequest(test *testing.T) {
	client := &http.Client{
		Transport: &Mocks.HttpTransportMock{
			Error: errors.New("Bad Request"),
		},
	}
	request := Request.NewRequest("/workspaces", Request.GET)

	_, error := Request.Send[[]Entities.WorkspaceEntity](client, request)
	assert.NotNil(test, error)
	assert.Contains(test, error.Error(), "Bad Request")
}

func TestErrorOnReadResponse(test *testing.T) {
	client := &http.Client{
		Transport: &Mocks.HttpTransportMock{
			Response: &http.Response{
				StatusCode: http.StatusBadRequest,
				Body:       io.NopCloser(&Mocks.CorrompedResponse{Message: "Bad Request"}),
			},
		},
	}
	request := Request.NewRequest("/workspaces", Request.GET)

	_, error := Request.Send[[]Entities.WorkspaceEntity](client, request)
	assert.NotNil(test, error)
	assert.Contains(test, error.Error(), "Bad Request")
}

func TestErrorOnDesserializeResponse(test *testing.T) {
	client := &http.Client{
		Transport: &Mocks.HttpTransportMock{
			Response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(`"key": invalid json`)),
			},
		},
	}
	request := Request.NewRequest("/workspaces", Request.GET)

	_, error := Request.Send[[]Entities.WorkspaceEntity](client, request)
	assert.NotNil(test, error)
	assert.Contains(test, error.Error(), "invalid character")
}

func TestBadStatusCodeResponse(test *testing.T) {
	client := &http.Client{
		Transport: &Mocks.HttpTransportMock{
			Response: &http.Response{
				StatusCode: http.StatusBadRequest,
			},
		},
	}
	request := Request.NewRequest("/workspaces", Request.GET)

	_, error := Request.Send[[]Entities.WorkspaceEntity](client, request)
	assert.NotNil(test, error)
	assert.Contains(test, error.Error(), "Bad Request")
}
