package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"tcms/src/api/handlers/dto/request"
	"tcms/src/api/handlers/dto/response"
	"tcms/src/core/interfaces/useCases"
)

type ProjectHandlers struct {
	services useCases.ProjectManager
}

func (instance ProjectHandlers) CreateProject(context echo.Context) error {
	createProjectDTO := request.CreateProjectDTO{}

	projectInstance, validationError := createProjectDTO.ConvertToDomain()
	if validationError != nil {
		return getHttpHandledErrorResponse(context, validationError)
	}

	id, createError := instance.services.CreateProject(*projectInstance)
	if createError != nil {
		return getHttpHandledErrorResponse(context, createError)
	}

	return context.JSON(http.StatusCreated, response.EntityCreated{ID: id})
}
