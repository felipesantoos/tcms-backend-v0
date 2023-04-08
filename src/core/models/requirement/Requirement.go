package requirement

import (
	"github.com/google/uuid"
	"reflect"
	"time"
)

type Requirement struct {
	id          uuid.UUID
	createdAt   time.Time
	updatedAt   time.Time
	deletedAt   time.Time
	name        string
	description string
	projectID   uuid.UUID
}

func (instance *Requirement) ID() uuid.UUID {
	return instance.id
}

func (instance *Requirement) CreatedAt() time.Time {
	return instance.createdAt
}

func (instance *Requirement) UpdatedAt() time.Time {
	return instance.updatedAt
}

func (instance *Requirement) DeletedAt() time.Time {
	return instance.deletedAt
}

func (instance *Requirement) Name() string {
	return instance.name
}

func (instance *Requirement) Description() string {
	return instance.description
}

func (instance *Requirement) ProjectID() uuid.UUID {
	return instance.projectID
}

func (instance *Requirement) IsZero() bool {
	return reflect.DeepEqual(instance, &Requirement{})
}

func NewForShortView(id uuid.UUID, name, description string, projectID uuid.UUID) *Requirement {
	return &Requirement{
		id:          id,
		name:        name,
		description: description,
		projectID:   projectID,
	}
}

func NewForDetailedView(id uuid.UUID, createdAt, updatedAt, deletedAt time.Time, name, description string,
	projectID uuid.UUID) *Requirement {
	return &Requirement{
		id:          id,
		name:        name,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
		deletedAt:   deletedAt,
		description: description,
		projectID:   projectID,
	}
}
