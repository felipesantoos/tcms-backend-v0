package postgres

import (
	"github.com/felipesantoos/tcms/src/infra/repository/postgres/orm"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

type Connector struct{}

func NewConnector() *Connector {
	return &Connector{}
}

func (connector Connector) GetConnection() (*gorm.DB, error) {
	info := "host=localhost user=root password=root dbname=root port=5432"
	db, err := gorm.Open(postgres.Open(info))
	if err != nil {
		log.Println("Error while trying to connect to the database...")
		return nil, err
	}

	return db, nil
}

const (
	requerimentTestCases   = "requeriment_test_cases"
	idIsNotNull            = "id IS NOT NULL"
	requirementIDIsNotNull = "requirement_id IS NOT NULL"
)

type Tables struct {
	project      orm.Project
	requirement  orm.Requirement
	testCase     orm.TestCase
	testCaseStep orm.TestCaseStep
}

var tables = Tables{}

func ResetTableObjects() {
	tables = Tables{}
}

func CreateTables() error {
	db, err := NewConnector().GetConnection()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = db.AutoMigrate(&tables.project, &tables.requirement, &tables.testCase, &tables.testCaseStep)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func TruncateTables() error {
	db, err := NewConnector().GetConnection()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	db.Table(requerimentTestCases).Unscoped().Where(requirementIDIsNotNull).Delete(nil)
	db.Unscoped().Where(idIsNotNull).Delete(&tables.testCaseStep)
	db.Unscoped().Where(idIsNotNull).Delete(&tables.testCase)
	db.Unscoped().Where(idIsNotNull).Delete(&tables.requirement)
	db.Unscoped().Where(idIsNotNull).Delete(&tables.project)

	return nil
}

func LoadFakeData() error {
	db, err := NewConnector().GetConnection()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	tables.project = orm.Project{
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   gorm.DeletedAt{},
		Name:        "Neon Nexus",
		Description: "An intuitive platform for managing software projects. Collaborate seamlessly with your team and streamline your workflow. With built-in reporting and analytics, stay on top of your progress and make data-driven decisions.",
		IsActive:    true,
	}

	tables.testCaseStep = orm.TestCaseStep{
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		DeletedAt:      gorm.DeletedAt{},
		TestCaseID:     uuid.UUID{},
		Description:    "Description of example",
		ExpectedResult: "Expected result of example",
	}

	tables.testCase = orm.TestCase{
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		DeletedAt:     gorm.DeletedAt{},
		Code:          1,
		Version:       1,
		Title:         "Title of example",
		Summary:       "Summary of example",
		Importance:    "Importance of example",
		ExecutionType: "Execution type of example",
		Precondition:  "Precondition of example",
		IsActive:      true,
		Steps:         []orm.TestCaseStep{tables.testCaseStep},
	}

	tables.requirement = orm.Requirement{
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   gorm.DeletedAt{},
		Name:        "Requirement of example",
		Description: "Description of example.",
		ProjectID:   tables.project.ID,
		Project:     tables.project,
		TestCases:   []orm.TestCase{tables.testCase},
	}

	result := db.Create(&tables.requirement)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return result.Error
	}

	return nil
}
