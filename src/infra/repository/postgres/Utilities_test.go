package postgres

import (
	"errors"
	"gorm.io/gorm"
)

type ConnectorMock struct{}

func (connector ConnectorMock) GetConnection() (*gorm.DB, error) {
	return nil, errors.New("connection error")
}

func SetUpTables() {
	CreateTables()
	TruncateTables()
	ResetTableObjects()
}
