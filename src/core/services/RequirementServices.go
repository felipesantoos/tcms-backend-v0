package services

import (
	"github.com/felipesantoos/tcms/src/core/filters"
	"github.com/felipesantoos/tcms/src/core/interfaces/repository"
	"github.com/felipesantoos/tcms/src/core/interfaces/usecases"
	"github.com/felipesantoos/tcms/src/core/models/requirement"
	"github.com/google/uuid"
)

var _ usecases.RequirementManager = &RequirementServices{}

type RequirementServices struct {
	requirementRepository repository.RequirementLoader
}

func (instance *RequirementServices) GetRequirements(requirementFilters filters.RequirementFilters) (
	[]requirement.Requirement, error) {
	return instance.requirementRepository.GetRequirements(requirementFilters)
}

func (instance *RequirementServices) GetRequirement(requirementID uuid.UUID) (*requirement.Requirement, error) {
	return instance.requirementRepository.GetRequirement(requirementID)
}

func (instance *RequirementServices) CreateRequirement(_requirement requirement.Requirement) (*uuid.UUID, error) {
	return instance.requirementRepository.CreateRequirement(_requirement)
}

func (instance *RequirementServices) DeleteRequirement(requirementID uuid.UUID) error {
	return instance.requirementRepository.DeleteRequirement(requirementID)
}

func (instance *RequirementServices) UpdateRequirement(_requirement requirement.Requirement) error {
	return instance.requirementRepository.UpdateRequirement(_requirement)
}

func NewRequirementServices(requirementRepository repository.RequirementLoader) *RequirementServices {
	return &RequirementServices{requirementRepository: requirementRepository}
}
