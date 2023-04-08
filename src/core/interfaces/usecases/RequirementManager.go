package usecases

import (
	"github.com/felipesantoos/tcms/src/core/filters"
	"github.com/felipesantoos/tcms/src/core/models/requirement"
	"github.com/google/uuid"
)

type RequirementManager interface {
	GetRequirements(requirementFilters filters.RequirementFilters) ([]requirement.Requirement, error)
	GetRequirement(requirementID uuid.UUID) (*requirement.Requirement, error)
	CreateRequirement(_requirement requirement.Requirement) (*uuid.UUID, error)
	DeleteRequirement(requirementID uuid.UUID) error
	UpdateRequirement(_requirement requirement.Requirement) error
}
