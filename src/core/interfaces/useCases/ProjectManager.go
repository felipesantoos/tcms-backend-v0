package useCases

import (
	"github.com/google/uuid"
	"tcms/src/core/domain/project"
	"tcms/src/core/errors"
)

type ProjectManager interface {
	CreateProject(projectInstance project.Project) (*uuid.UUID, errors.Error)
}
