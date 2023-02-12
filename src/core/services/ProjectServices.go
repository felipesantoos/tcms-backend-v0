package services

import (
	"github.com/google/uuid"
	"tcms/src/core/domain/project"
	"tcms/src/core/errors"
	"tcms/src/core/errors/logger"
	"tcms/src/core/interfaces/repository"
)

var _ repository.ProjectLoader = ProjectServices{}

type ProjectServices struct {
	projectRepository repository.ProjectLoader
	loggerInstance    logger.Logger
}

func (instance ProjectServices) CreateProject(projectInstance project.Project) (*uuid.UUID, errors.Error) {
	id, err := instance.projectRepository.CreateProject(projectInstance)
	if err != nil {
		instance.loggerInstance.Log(err)
	}

	return id, nil
}
