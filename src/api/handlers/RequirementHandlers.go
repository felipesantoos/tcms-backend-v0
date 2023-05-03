package handlers

import (
	"github.com/felipesantoos/tcms/src/api/handlers/dto/request"
	"github.com/felipesantoos/tcms/src/api/handlers/dto/response"
	"github.com/felipesantoos/tcms/src/api/handlers/keys"
	"github.com/felipesantoos/tcms/src/api/handlers/params"
	"github.com/felipesantoos/tcms/src/core/filters"
	"github.com/felipesantoos/tcms/src/core/interfaces/usecases"
	"github.com/felipesantoos/tcms/src/core/models/requirement"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type RequirementHandlers struct {
	requirementServices usecases.RequirementManager
}

func (instance *RequirementHandlers) GetRequirements(context *gin.Context) {
	name := context.Query(params.Name)
	requirementFilters := filters.RequirementFilters{Name: name}

	requirements, err := instance.requirementServices.GetRequirements(requirementFilters)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{keys.Error: err.Error()})
		return
	}

	context.JSON(http.StatusOK, response.NewRequirementList(requirements))
}

func (instance *RequirementHandlers) GetRequirement(context *gin.Context) {
	requirementID, err := uuid.Parse(context.Params.ByName(params.RequirementID))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{keys.Error: err.Error()})
		return
	}

	_requirement, err := instance.requirementServices.GetRequirement(requirementID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{keys.Error: err.Error()})
		return
	}

	context.JSON(http.StatusOK, response.NewRequirement(*_requirement))
}

func (instance *RequirementHandlers) CreateRequirement(context *gin.Context) {
	requirementDTO := request.Requirement{}
	err := context.ShouldBindJSON(&requirementDTO)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{keys.Error: err.Error()})
		return
	}

	requirementBuilder := requirement.NewBuilder()
	requirementBuilder.Name(requirementDTO.Name).Description(requirementDTO.Description).
		ProjectID(requirementDTO.ProjectID)
	_requirement, validationError := requirementBuilder.Build()
	if validationError != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{keys.Error: validationError.Error()})
		return
	}

	requirementID, err := instance.requirementServices.CreateRequirement(*_requirement)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{keys.Error: err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{keys.ID: requirementID})
}

func (instance *RequirementHandlers) DeleteRequirement(context *gin.Context) {
	requirementID, err := uuid.Parse(context.Params.ByName(params.RequirementID))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{keys.Error: err.Error()})
		return
	}

	err = instance.requirementServices.DeleteRequirement(requirementID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{keys.Error: err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, gin.H{})
}

func (instance *RequirementHandlers) UpdateRequirement(context *gin.Context) {
	requirementID, err := uuid.Parse(context.Params.ByName(params.RequirementID))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{keys.Error: err.Error()})
		return
	}

	requirementDTO := request.Requirement{}
	err = context.ShouldBindJSON(&requirementDTO)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{keys.Error: err.Error()})
		return
	}

	requirementBuilder := requirement.NewBuilder()
	requirementBuilder.ID(requirementID).Name(requirementDTO.Name).Description(requirementDTO.Description).
		ProjectID(requirementDTO.ProjectID)
	_requirement, validationError := requirementBuilder.Build()
	if validationError != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{keys.Error: validationError.Error()})
		return
	}

	err = instance.requirementServices.UpdateRequirement(*_requirement)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{keys.Error: err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, gin.H{})
}

func NewRequirementHandlers(requirementServices usecases.RequirementManager) *RequirementHandlers {
	return &RequirementHandlers{requirementServices: requirementServices}
}
