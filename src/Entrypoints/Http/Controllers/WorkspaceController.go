package Controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"toggl-xlsx-back/src/Application/UseCases/Workspace"
)

type WorkspaceController struct {
	router               *gin.Engine
	getWorkspacesUseCase *Workspace.GetWorkspacesUseCase
}

func NewWorkspaceController(router *gin.Engine, getWorkspacesUseCase *Workspace.GetWorkspacesUseCase) *WorkspaceController {
	controller := WorkspaceController{
		router:               router,
		getWorkspacesUseCase: getWorkspacesUseCase,
	}

	controller.makeRoutes()
	return &controller
}

func (this *WorkspaceController) makeRoutes() {
	group := this.router.Group("/workspaces")
	{
		group.GET("", this.GetWorkspaces)
	}
}

func (this *WorkspaceController) GetWorkspaces(context *gin.Context) {
	username, password, _ := context.Request.BasicAuth()
	workspaces, error := this.getWorkspacesUseCase.Execute(username, password)
	if error != nil {
		context.Error(error)
		return
	}

	context.JSON(http.StatusOK, workspaces)
}
