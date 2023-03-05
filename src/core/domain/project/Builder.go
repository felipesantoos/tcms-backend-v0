package project

import (
	"github.com/google/uuid"
	"tcms/src/core/errors"
	"tcms/src/core/messages"
	"time"
)

type builder struct {
	project       *Project
	invalidFields []errors.InvalidField
}

func NewBuilder() *builder {
	return &builder{project: &Project{}}
}

func (instance *builder) WithID(id uuid.UUID) *builder {
	instance.project.id = &id
	return instance
}

func (instance *builder) WithName(name string) *builder {
	instance.project.name = name
	return instance
}

func (instance *builder) WithDescription(description string) *builder {
	instance.project.description = description
	return instance
}

func (instance *builder) WithIsActive(isActive bool) *builder {
	instance.project.isActive = isActive
	return instance
}

func (instance *builder) WithIsDeleted(isDeleted bool) *builder {
	instance.project.isDeleted = isDeleted
	return instance
}

func (instance *builder) WithCreatedAt(createdAt time.Time) *builder {
	instance.project.createdAt = &createdAt
	return instance
}

func (instance *builder) WithUpdatedAt(updatedAt time.Time) *builder {
	instance.project.updatedAt = &updatedAt
	return instance
}

func (instance *builder) Build() (*Project, errors.Error) {
	if len(instance.invalidFields) > 0 {
		return nil, errors.NewValidationError(messages.InvalidProjectErrorMessage, instance.invalidFields...)
	}
	return instance.project, nil
}
