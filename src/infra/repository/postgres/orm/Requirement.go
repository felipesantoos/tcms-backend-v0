package orm

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Requirement struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;sql:type:uuid;default:gen_random_uuid()"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string
	Description string
	ProjectID   uuid.UUID
	Project     Project
	TestCases   []TestCase `gorm:"many2many:requeriment_test_cases"`
}

func (requirementDTO *Requirement) BeforeCreate(tx *gorm.DB) error {
	requirementDTO.ID = uuid.New()
	return nil
}
