package request

import (
	"github.com/felipesantoos/tcms/src/core/models/project"
	"github.com/google/uuid"
)

type Project struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

func (instance *Project) ConvertToModel() *project.Project {
	return project.New(instance.ID, instance.Name, instance.Description)
}
