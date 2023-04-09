package handlers

import (
	"github.com/felipesantoos/tcms/src/api/handlers/dto/request"
	"github.com/felipesantoos/tcms/src/api/handlers/dto/response"
	"github.com/felipesantoos/tcms/src/api/handlers/keys"
	"github.com/felipesantoos/tcms/src/api/handlers/params"
	"github.com/felipesantoos/tcms/src/core/filters"
	"github.com/felipesantoos/tcms/src/core/interfaces/usecases"
	"github.com/felipesantoos/tcms/src/core/models/project"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type ProjectHandlers struct {
	projectServices usecases.ProjectManager
}

func (instance *ProjectHandlers) GetProjects(context *gin.Context) {
	name := context.Query(params.Name)
	projectFilters := filters.ProjectFilters{Name: name}

	projects, err := instance.projectServices.GetProjects(projectFilters)
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

	projectBuilder := project.NewBuilder()
	projectBuilder.Name(projectDTO.Name).Description(projectDTO.Description)
	_project, err := projectBuilder.Build()
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{keys.Error: err.Error()})
		return
	}

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

	projectBuilder := project.NewBuilder()
	projectBuilder.ID(projectID).Name(projectDTO.Name).Description(projectDTO.Description)
	_project, err := projectBuilder.Build()
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{keys.Error: err.Error()})
		return
	}

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
