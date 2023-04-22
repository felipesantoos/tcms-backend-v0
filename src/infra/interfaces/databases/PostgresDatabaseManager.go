package databases

import (
	"gorm.io/gorm"
)

type PostgresDatabaseManager interface {
	GetConnection() (*gorm.DB, error)
}
