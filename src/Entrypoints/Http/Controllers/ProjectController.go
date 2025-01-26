package Controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"toggl-xlsx-back/src/Application/Errors"
	"toggl-xlsx-back/src/Application/UseCases/Project"
)

type ProjectController struct {
	router             *gin.Engine
	getProjectsUseCase *Project.GetProjectsUseCase
}

func NewProjectController(router *gin.Engine, getProjectsUseCase *Project.GetProjectsUseCase) *ProjectController {
	controller := ProjectController{
		router:             router,
		getProjectsUseCase: getProjectsUseCase,
	}

	controller.makeRoutes()
	return &controller
}

func (this *ProjectController) makeRoutes() {
	this.router.GET("projects/:workspaceId", this.GetProjects)
}

func (this *ProjectController) GetProjects(context *gin.Context) {
	workspaceId, haveWorkspaceId := context.Params.Get("workspaceId")
	if !haveWorkspaceId {
		context.Error(&Errors.BadRequest{Message: "workspaceId is required"})
		return
	}

	convertedWorkspaceId, error := strconv.Atoi(workspaceId)
	if error != nil {
		context.Error(&Errors.BadRequest{Message: "workspaceId must be a number"})
		return
	}

	username, password, _ := context.Request.BasicAuth()
	projects, error := this.getProjectsUseCase.Execute(username, password, convertedWorkspaceId)
	if error != nil {
		context.Error(error)
		return
	}

	context.JSON(http.StatusOK, projects)
}
