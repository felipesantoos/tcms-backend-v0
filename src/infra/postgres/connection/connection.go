package connection

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"time"
)

type Manager interface {
	GetConnection() (*sqlx.DB, error)
	CloseConnection(connectionInstance *sqlx.DB)
}

var _ Manager = DatabaseConnectionManager{}

type DatabaseConnectionManager struct{}

func (instance DatabaseConnectionManager) GetConnection() (*sqlx.DB, error) {
	uri := getDatabaseURI()
	db, err := sqlx.Open("postgres", uri)
	if err != nil {
		log.Print("Error while accessing database: " + err.Error())
		return nil, err
	}

	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}

func (instance DatabaseConnectionManager) CloseConnection(connectionInstance *sqlx.DB) {
	err := connectionInstance.Close()
	if err != nil {
		log.Error().Err(err)
	}
}
