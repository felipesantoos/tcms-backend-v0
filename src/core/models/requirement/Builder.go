package requirement

import (
	"errors"
	"github.com/google/uuid"
	"strings"
	"time"
)

type builder struct {
	requirement   *Requirement
	invalidFields []string
}

func NewBuilder() *builder {
	return &builder{requirement: &Requirement{}}
}

func (instance *builder) ID(id uuid.UUID) *builder {
	if id.ID() == 0 {
		instance.invalidFields = append(instance.invalidFields, "O ID do requerimento é inválido!")
		return instance
	}
	instance.requirement.id = id
	return instance
}

func (instance *builder) CreatedAt(createdAt time.Time) *builder {
	if createdAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "A data de criação do requerimento é inválida!")
		return instance
	}
	instance.requirement.createdAt = createdAt
	return instance
}

func (instance *builder) UpdatedAt(updatedAt time.Time) *builder {
	if updatedAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "A data de atualização do requerimento é inválida!")
		return instance
	}
	instance.requirement.updatedAt = updatedAt
	return instance
}

func (instance *builder) DeletedAt(deletedAt time.Time) *builder {
	instance.requirement.deletedAt = deletedAt
	return instance
}

func (instance *builder) Name(name string) *builder {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		instance.invalidFields = append(instance.invalidFields, "O nome do requerimento é inválido!")
		return instance
	}
	instance.requirement.name = name
	return instance
}

func (instance *builder) Description(description string) *builder {
	instance.requirement.description = description
	return instance
}

func (instance *builder) ProjectId(projectId uuid.UUID) *builder {
	if projectId.ID() == 0 {
		instance.invalidFields = append(instance.invalidFields, "O ID do projeto é inválido!")
		return instance
	}
	instance.requirement.projectID = projectId
	return instance
}

func (instance *builder) Build() (*Requirement, error) {
	if len(instance.invalidFields) > 0 {
		return nil, errors.New(strings.Join(instance.invalidFields, ";"))
	}

	return instance.requirement, nil
}
