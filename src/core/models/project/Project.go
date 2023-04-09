package project

import (
	"github.com/google/uuid"
	"reflect"
	"time"
)

type Project struct {
	id          uuid.UUID
	createdAt   time.Time
	updatedAt   time.Time
	deletedAt   time.Time
	name        string
	description string
}

func (instance *Project) ID() uuid.UUID {
	return instance.id
}

func (instance *Project) CreatedAt() time.Time {
	return instance.createdAt
}

func (instance *Project) UpdatedAt() time.Time {
	return instance.updatedAt
}

func (instance *Project) DeletedAt() time.Time {
	return instance.deletedAt
}

func (instance *Project) Name() string {
	return instance.name
}

func (instance *Project) Description() string {
	return instance.description
}

func (instance *Project) IsZero() bool {
	return reflect.DeepEqual(instance, &Project{})
}
