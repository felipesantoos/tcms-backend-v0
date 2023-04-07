package databases

import (
	"gorm.io/gorm"
)

type DatabaseManager interface {
	GetConnection() (*gorm.DB, error)
}
