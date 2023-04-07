package handlers

import (
	"github.com/felipesantoos/tcms/src/api/handlers/dto/request"
	"github.com/felipesantoos/tcms/src/api/handlers/dto/response"
	"github.com/felipesantoos/tcms/src/api/handlers/keys"
	"github.com/felipesantoos/tcms/src/api/handlers/params"
	"github.com/felipesantoos/tcms/src/core/interfaces/usecases"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type ProjectHandlers struct {
	projectServices usecases.ProjectManager
}

func (instance *ProjectHandlers) GetProjects(context *gin.Context) {
	projects, err := instance.projectServices.GetProjects()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{keys.Error: err.Error()})
		return
	}

	context.JSON(http.StatusOK, response.NewProjectList(projects))
}

func (instance *ProjectHandlers) GetProject(context *gin.Context) {
	projectID, err := uuid.Parse(context.Params.ByName(params.ProjectID))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{keys.Error: err.Error()})
		return
	}

	_project, err := instance.projectServices.GetProject(projectID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{keys.Error: err.Error()})
		return
	}

	context.JSON(http.StatusOK, response.NewProject(*_project))
}

func (instance *ProjectHandlers) CreateProject(context *gin.Context) {
	projectDTO := request.Project{}
	err := context.ShouldBindJSON(&projectDTO)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{keys.Error: err.Error()})
		return
	}

	_project := projectDTO.ConvertToModel()

	projectID, err := instance.projectServices.CreateProject(*_project)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{keys.Error: err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{keys.ID: projectID})
}

func (instance *ProjectHandlers) DeleteProject(context *gin.Context) {
	projectID, err := uuid.Parse(context.Params.ByName(params.ProjectID))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{keys.Error: err.Error()})
		return
	}

	err = instance.projectServices.DeleteProject(projectID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{keys.Error: err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, gin.H{})
}

func (instance *ProjectHandlers) UpdateProject(context *gin.Context) {
	projectID, err := uuid.Parse(context.Params.ByName(params.ProjectID))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{keys.Error: err.Error()})
		return
	}

	projectDTO := request.Project{}
	err = context.ShouldBindJSON(&projectDTO)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{keys.Error: err.Error()})
		return
	}

	projectDTO.ID = projectID
	_project := projectDTO.ConvertToModel()

	err = instance.projectServices.UpdateProject(*_project)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{keys.Error: err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, gin.H{})
}

func NewProjectHandlers(projectServices usecases.ProjectManager) *ProjectHandlers {
	return &ProjectHandlers{projectServices: projectServices}
}