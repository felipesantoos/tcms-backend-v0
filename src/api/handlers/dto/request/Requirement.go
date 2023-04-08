package request

import (
	"github.com/felipesantoos/tcms/src/core/models/requirement"
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

func (instance *Requirement) ConvertToModel() *requirement.Requirement {
	return requirement.NewForDetailedView(instance.ID, instance.CreatedAt, instance.UpdatedAt, instance.DeletedAt,
		instance.Name, instance.Description, instance.ProjectID)
}
