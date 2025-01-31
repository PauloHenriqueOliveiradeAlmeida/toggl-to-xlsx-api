package Track

import (
	"encoding/base64"
	"net/http"
	"os"
	"strconv"
	"time"
	"toggl-xlsx-back/src/Application/Services/Request"
	"toggl-xlsx-back/src/Domain/Entities"
)

type TogglClient struct {
	email    string
	password string
	client   *http.Client
}

func NewTogglClient(client *http.Client) *TogglClient {
	return &TogglClient{
		client: client,
	}
}

func (this *TogglClient) SetCredentials(email string, password string) {
	this.email = email
	this.password = password
}

func (this *TogglClient) GetWorkspaces() ([]Entities.WorkspaceEntity, error) {
	request := this.buildGetRequest(os.Getenv("TRACK_API_BASE_URL") + "/me/workspaces")
	response, error := Request.Send[[]Entities.WorkspaceEntity](this.client, request)
	return response, error
}

func (this *TogglClient) GetProjects(workspaceId int) ([]Entities.ProjectEntity, error) {
	request := this.buildGetRequest(os.Getenv("TRACK_API_BASE_URL") + "/workspaces/" + strconv.Itoa(workspaceId) + "/projects")
	response, error := Request.Send[[]Entities.ProjectEntity](this.client, request)
	return response, error
}

func (this *TogglClient) GetTimeEntries(startDate time.Time, endDate time.Time) ([]Entities.TimeEntryEntity, error) {
	request := this.buildGetRequest(os.Getenv("TRACK_API_BASE_URL") + "/me/time_entries?start_date=" + startDate.Format("2006-01-02") + "&end_date=" + endDate.Format("2006-01-02"))
	response, error := Request.Send[[]Entities.TimeEntryEntity](this.client, request)
	return response, error
}

func (this *TogglClient) buildGetRequest(endpoint string) *Request.Request {
	basicToken := base64.StdEncoding.EncodeToString([]byte(this.email + ":" + this.password))
	return Request.NewRequest(endpoint, Request.GET).WithHeaders(map[string]string{
		"Authorization": "Basic " + basicToken,
	})
}
