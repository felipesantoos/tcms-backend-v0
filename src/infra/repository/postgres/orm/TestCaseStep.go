package orm

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type TestCaseStep struct {
	ID             uuid.UUID `gorm:"primaryKey;type:uuid;sql:type:uuid;default:gen_random_uuid()"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	TestCaseID     uuid.UUID
	Description    string
	ExpectedResult string
}
