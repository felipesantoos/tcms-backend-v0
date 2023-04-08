package orm

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type TestCase struct {
	ID            uuid.UUID `gorm:"primaryKey;type:uuid;sql:type:uuid;default:gen_random_uuid()"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	Code          uint64         `gorm:"unique"`
	Version       uint64
	Title         string
	Summary       string
	Importance    string
	ExecutionType string
	Precondition  string
	IsActive      bool `gorm:"default:true"`
	Steps         []TestCaseStep
}
