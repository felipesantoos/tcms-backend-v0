package response

import (
	"github.com/google/uuid"
)

type EntityCreated struct {
	ID *uuid.UUID `json:"id"`
}
