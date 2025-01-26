package Tests

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"toggl-xlsx-back/src/Application/Errors"
	"toggl-xlsx-back/src/Application/Services/Track"
	"toggl-xlsx-back/src/Application/UseCases/Workspace"
	"toggl-xlsx-back/src/Domain/Entities"
	"toggl-xlsx-back/src/Entrypoints/Http/Controllers"
	"toggl-xlsx-back/src/Entrypoints/Http/Middlewares"
	"toggl-xlsx-back/src/Tests/Mocks"
)

func TestGetWorkspacesShouldBeSuccess(test *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(Middlewares.ErrorHandler())
	workspaces := []Entities.WorkspaceEntity{
		{Id: 1, OrganizationId: 1, Name: "Workspace 1"},
		{Id: 2, OrganizationId: 2, Name: "Workspace 2"},
	}

	trackMock := Mocks.NewTrackMock(workspaces, nil)

	controllers.NewWorkspaceController(router, Workspace.NewGetWorkspacesUseCase(Track.NewTrackService(trackMock)))

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/workspaces", nil)
	router.ServeHTTP(recorder, request)

	assert.Equal(test, http.StatusOK, recorder.Code)

	payload, _ := json.Marshal(workspaces)
	assert.Equal(test, string(payload), recorder.Body.String())
}

func TestErrorOnGetWorkspaces(test *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(Middlewares.ErrorHandler())
	error := &Errors.ServiceUnavailable{Message: "Serviço de track indisponível"}

	trackMock := Mocks.NewTrackMock([]Entities.WorkspaceEntity{}, error)
	controllers.NewWorkspaceController(router, Workspace.NewGetWorkspacesUseCase(Track.NewTrackService(trackMock)))

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/workspaces", nil)
	router.ServeHTTP(recorder, request)

	fmt.Println(recorder)
	assert.Equal(test, http.StatusServiceUnavailable, recorder.Code)

	payload, _ := json.Marshal(error)
	assert.Equal(test, string(payload), recorder.Body.String())
}
