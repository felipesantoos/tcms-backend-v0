package dto

import (
	"github.com/felipesantoos/tcms/src/core/models/tc"
	"github.com/google/uuid"
	"time"
)

type Requirement struct {
	ID          string        `bson:"_id"`
	CreatedAt   time.Time     `bson:"created_at"`
	UpdatedAt   time.Time     `bson:"updated_at"`
	DeletedAt   time.Time     `bson:"deleted_at"`
	Name        string        `bson:"name"`
	Description string        `bson:"description"`
	ProjectID   uuid.UUID     `bson:"project_id"`
	TestCases   []tc.TestCase `bson:"test_cases"`
}
