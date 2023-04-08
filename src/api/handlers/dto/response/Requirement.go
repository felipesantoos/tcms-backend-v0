package response

import (
	"github.com/felipesantoos/tcms/src/core/models/requirement"
	"github.com/google/uuid"
	"time"
)

type Requirement struct {
	ID          uuid.UUID  `json:"id"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	ProjectID   uuid.UUID  `json:"project_id"`
}

func NewRequirement(_requirement requirement.Requirement) *Requirement {
	var createdAt *time.Time
	var updatedAt *time.Time
	var deletedAt *time.Time

	if !_requirement.IsZero() {
		if !_requirement.CreatedAt().IsZero() {
			aux := _requirement.CreatedAt()
			createdAt = &aux
		}

		if !_requirement.UpdatedAt().IsZero() {
			aux := _requirement.UpdatedAt()
			updatedAt = &aux
		}

		if !_requirement.DeletedAt().IsZero() {
			aux := _requirement.DeletedAt()
			deletedAt = &aux
		}
	}

	requirementDTO := &Requirement{
		ID:          _requirement.ID(),
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		DeletedAt:   deletedAt,
		Name:        _requirement.Name(),
		Description: _requirement.Description(),
		ProjectID:   _requirement.ProjectID(),
	}

	return requirementDTO
}

func NewRequirementList(requirements []requirement.Requirement) []Requirement {
	requirementDTOs := make([]Requirement, 0)
	for _, _requirement := range requirements {
		requirementDTOs = append(requirementDTOs, *NewRequirement(_requirement))
	}

	return requirementDTOs
}
