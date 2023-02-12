package project

import (
	"github.com/google/uuid"
	"time"
)

type Project struct {
	id          *uuid.UUID
	name        string
	description string
	isActive    bool
	isDeleted   bool
	createdAt   *time.Time
	updatedAt   *time.Time
}

func (instance *Project) ID() *uuid.UUID {
	return instance.id
}

func (instance *Project) Name() string {
	return instance.name
}

func (instance *Project) Description() string {
	return instance.description
}

func (instance *Project) IsActive() bool {
	return instance.isActive
}

func (instance *Project) IsDeleted() bool {
	return instance.isDeleted
}

func (instance *Project) CreatedAt() *time.Time {
	return instance.createdAt
}

func (instance *Project) UpdatedAt() *time.Time {
	return instance.updatedAt
}

func (instance *Project) IsZero() bool {
	return instance == nil || instance == &Project{}
}
