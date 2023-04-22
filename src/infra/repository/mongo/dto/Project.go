package dto

import (
	"github.com/google/uuid"
	"time"
)

type Project struct {
	ID          string    `bson:"_id"`
	CreatedAt   time.Time `bson:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at"`
	DeletedAt   time.Time `bson:"deleted_at"`
	Name        string    `bson:"name"`
	Description string    `bson:"description"`
	IsActive    bool      `bson:"is_active"`
}

func NewProject(id uuid.UUID, createdAt, updatedAt, deletedAt time.Time, name, description string, isActive bool) *Project {
	return &Project{
		ID:          id.String(),
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		DeletedAt:   deletedAt,
		Name:        name,
		Description: description,
		IsActive:    isActive,
	}
}
