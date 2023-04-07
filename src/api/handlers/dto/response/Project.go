package response

import (
	"github.com/felipesantoos/tcms/src/core/models/project"
	"github.com/google/uuid"
)

type Project struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

func NewProject(projectInstance project.Project) *Project {
	return &Project{
		ID:          projectInstance.ID(),
		Name:        projectInstance.Name(),
		Description: projectInstance.Description(),
	}
}

func NewProjectList(projects []project.Project) []Project {
	projectResponseList := make([]Project, 0)
	for _, projectInstance := range projects {
		projectResponseList = append(projectResponseList, *NewProject(projectInstance))
	}

	return projectResponseList
}
