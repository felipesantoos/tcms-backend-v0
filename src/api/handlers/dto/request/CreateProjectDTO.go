package request

import (
	"tcms/src/core/domain/project"
	"tcms/src/core/errors"
)

type CreateProjectDTO struct {
	Name        string
	Description string
}

func (instance CreateProjectDTO) ConvertToDomain() (*project.Project, errors.Error) {
	projectBuilder := project.NewBuilder()
	projectBuilder.WithName(instance.Name).WithDescription(instance.Description)

	projectInstance, validationError := projectBuilder.Build()
	if validationError != nil {
		return nil, validationError
	}

	return projectInstance, nil
}
