package repository

import (
	"github.com/felipesantoos/tcms/src/core/models/project"
	"github.com/google/uuid"
)

type ProjectLoader interface {
	GetProjects() ([]project.Project, error)
	GetProject(projectID uuid.UUID) (*project.Project, error)
	CreateProject(_project project.Project) (*uuid.UUID, error)
	DeleteProject(projectID uuid.UUID) error
	UpdateProject(_project project.Project) error
}
