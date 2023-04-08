package response

import (
	"github.com/felipesantoos/tcms/src/core/models/project"
	"github.com/google/uuid"
	"time"
)

type Project struct {
	ID          uuid.UUID  `json:"id"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
}

func NewProject(_project project.Project) *Project {
	var createdAt *time.Time
	var updatedAt *time.Time
	var deletedAt *time.Time

	if !_project.IsZero() {
		if !_project.CreatedAt().IsZero() {
			aux := _project.CreatedAt()
			createdAt = &aux
		}

		if !_project.UpdatedAt().IsZero() {
			aux := _project.UpdatedAt()
			updatedAt = &aux
		}

		if !_project.DeletedAt().IsZero() {
			aux := _project.DeletedAt()
			deletedAt = &aux
		}
	}

	projectDTO := &Project{
		ID:          _project.ID(),
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		DeletedAt:   deletedAt,
		Name:        _project.Name(),
		Description: _project.Description(),
	}

	return projectDTO
}

func NewProjectList(projects []project.Project) []Project {
	projectDTOs := make([]Project, 0)
	for _, _project := range projects {
		projectDTOs = append(projectDTOs, *NewProject(_project))
	}

	return projectDTOs
}
