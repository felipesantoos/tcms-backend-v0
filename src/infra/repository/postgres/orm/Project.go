package orm

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Project struct {
	ID          uuid.UUID      `gorm:"primaryKey;type:uuid;sql:type:uuid;default:gen_random_uuid()"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string         `binding:"required"`
	Description string
}

func (projectDTO *Project) BeforeCreate(tx *gorm.DB) error {
	projectDTO.ID = uuid.New()
	return nil
}
