package postgres

import (
	"github.com/felipesantoos/tcms/src/infra/repository/postgres/orm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Connector struct{}

func (connector Connector) GetConnection() (*gorm.DB, error) {
	info := "host=localhost user=root password=root dbname=root port=5432"
	db, err := gorm.Open(postgres.Open(info))
	if err != nil {
		log.Println("Error while trying to connect to the database...")
		return nil, err
	}

	return db, nil
}

func (connector Connector) CreateTables() error {
	db, err := connector.GetConnection()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = db.AutoMigrate(&orm.Project{})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func NewConnector() *Connector {
	return &Connector{}
}
