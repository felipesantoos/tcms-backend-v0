package services

import (
	"github.com/felipesantoos/tcms/src/core/interfaces/repository"
	"github.com/felipesantoos/tcms/src/core/interfaces/usecases"
	"github.com/felipesantoos/tcms/src/core/models/project"
	"github.com/google/uuid"
)

var _ usecases.ProjectManager = &ProjectServices{}

type ProjectServices struct {
	projectRepository repository.ProjectLoader
}

func (instance *ProjectServices) GetProjects() ([]project.Project, error) {
	return instance.projectRepository.GetProjects()
}

func (instance *ProjectServices) GetProject(projectID uuid.UUID) (*project.Project, error) {
	return instance.projectRepository.GetProject(projectID)
}

func (instance *ProjectServices) CreateProject(_project project.Project) (*uuid.UUID, error) {
	return instance.projectRepository.CreateProject(_project)
}

func (instance *ProjectServices) DeleteProject(projectID uuid.UUID) error {
	return instance.projectRepository.DeleteProject(projectID)
}

func (instance *ProjectServices) UpdateProject(_project project.Project) error {
	return instance.projectRepository.UpdateProject(_project)
}

func NewProjectServices(projectRepository repository.ProjectLoader) *ProjectServices {
	return &ProjectServices{projectRepository: projectRepository}
}
