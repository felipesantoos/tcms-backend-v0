package project

import (
	"tcms/src/core/errors"
	"tcms/src/core/messages"
)

type builder struct {
	project       *Project
	invalidFields []errors.InvalidField
}

func NewBuilder() *builder {
	return &builder{project: &Project{}}
}

func (instance *builder) Build() (*Project, errors.Error) {
	if len(instance.invalidFields) > 0 {
		return nil, errors.NewValidationError(messages.InvalidProjectErrorMessage, instance.invalidFields...)
	}
	return instance.project, nil
}
