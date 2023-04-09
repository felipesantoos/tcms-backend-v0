package project

import (
	"errors"
	"github.com/google/uuid"
	"strings"
	"time"
)

type builder struct {
	project       *Project
	invalidFields []string
}

func NewBuilder() *builder {
	return &builder{project: &Project{}}
}

func (instance *builder) ID(id uuid.UUID) *builder {
	if id.ID() == 0 {
		instance.invalidFields = append(instance.invalidFields, "O ID do projeto é inválido!")
		return instance
	}
	instance.project.id = id
	return instance
}

func (instance *builder) CreatedAt(createdAt time.Time) *builder {
	if createdAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "A data de criação do projeto é inválida!")
		return instance
	}
	instance.project.createdAt = createdAt
	return instance
}

func (instance *builder) UpdatedAt(updatedAt time.Time) *builder {
	if updatedAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "A data de atualização do projeto é inválida!")
		return instance
	}
	instance.project.updatedAt = updatedAt
	return instance
}

func (instance *builder) DeletedAt(deletedAt time.Time) *builder {
	instance.project.deletedAt = deletedAt
	return instance
}

func (instance *builder) Name(name string) *builder {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		instance.invalidFields = append(instance.invalidFields, "O nome do projeto é inválido!")
		return instance
	}
	instance.project.name = name
	return instance
}

func (instance *builder) Description(description string) *builder {
	instance.project.description = description
	return instance
}

func (instance *builder) Build() (*Project, error) {
	if len(instance.invalidFields) > 0 {
		return nil, errors.New(strings.Join(instance.invalidFields, ";"))
	}

	return instance.project, nil
}
