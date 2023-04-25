package dto

import (
	"github.com/google/uuid"
	"time"
)

type Requirement struct {
	ID          string    `bson:"_id"`
	CreatedAt   time.Time `bson:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at"`
	DeletedAt   time.Time `bson:"deleted_at"`
	Name        string    `bson:"name"`
	Description string    `bson:"description"`
	ProjectID   string    `bson:"project_id"`
}

func NewRequirement(id uuid.UUID, createdAt, updatedAt, deletedAt time.Time, name, description string, projectId uuid.UUID) *Requirement {
	return &Requirement{
		ID:          id.String(),
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		DeletedAt:   deletedAt,
		Name:        name,
		Description: description,
		ProjectID:   projectId.String(),
	}
}
