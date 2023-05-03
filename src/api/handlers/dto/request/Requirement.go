package request

import (
	"github.com/google/uuid"
	"time"
)

type Requirement struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ProjectID   uuid.UUID `json:"project_id"`
}
