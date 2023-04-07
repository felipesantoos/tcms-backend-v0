package postgres

import (
	"github.com/felipesantoos/tcms/src/infra/repository/postgres/orm"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

func TestGetProjects_Success(t *testing.T) {
	// Create a test database and get a connection to it
	testDB, err := gorm.Open(postgres.Open("host=localhost user=root password=root dbname=root port=5432"))
	assert.NoError(t, err)
	defer testDB.Exec("DELETE FROM projects")

	// Create a new project
	newProject := &orm.Project{
		ID:          uuid.New(),
		Name:        "Test Project",
		Description: "This is a test project",
	}
	testDB.Create(newProject)

	// Create a new instance of the repository and call the GetProjects method
	repository := NewProjectPostgresRepository(NewConnector())
	projects, err := repository.GetProjects()
	assert.NoError(t, err)

	// Check that the returned project has the expected values
	assert.Len(t, projects, 1)
	assert.Equal(t, newProject.ID, projects[0].ID())
	assert.Equal(t, newProject.Name, projects[0].Name())
	assert.Equal(t, newProject.Description, projects[0].Description())
}

func TestGetProjects_EmptyResult(t *testing.T) {
	// Create a test database and get a connection to it
	testDB, err := gorm.Open(postgres.Open("host=localhost user=root password=root dbname=root port=5432"))
	assert.NoError(t, err)
	defer testDB.Exec("DELETE FROM projects")

	// Create a new instance of the repository and call the GetProjects method
	repository := NewProjectPostgresRepository(NewConnector())
	projects, err := repository.GetProjects()
	assert.NoError(t, err)

	// Check that the returned project slice is empty
	assert.Empty(t, projects)
}

func TestGetProjects_DatabaseError(t *testing.T) {
	// Create a mock repository with a faulty connector
	repository := &ProjectPostgresRepository{ConnectorMock{}}

	// Call the GetProjects method and check that an error is returned
	projects, err := repository.GetProjects()

	assert.Error(t, err, "an error should have been returned")
	assert.Nil(t, projects, "projects should be nil")
}

func TestGetProject_Success(t *testing.T) {
	// Create a test database and get a connection to it
	testDB, err := gorm.Open(postgres.Open("host=localhost user=root password=root dbname=root port=5432"))
	assert.NoError(t, err)
	defer testDB.Exec("DELETE FROM projects")

	// Create a new project
	newProject := &orm.Project{
		ID:          uuid.New(),
		Name:        "Test Project",
		Description: "This is a test project",
	}
	testDB.Create(newProject)

	// Create a new instance of the repository and call the GetProject method
	repository := NewProjectPostgresRepository(NewConnector())
	project, err := repository.GetProject(newProject.ID)
	assert.NoError(t, err)

	// Check that the returned project has the expected values
	assert.Equal(t, newProject.ID, project.ID(), "unexpected project ID")
	assert.Equal(t, newProject.Name, project.Name(), "unexpected project name")
	assert.Equal(t, newProject.Description, project.Description(), "unexpected project description")
}

func TestGetProject_ProjectNotFound(t *testing.T) {
	// Create a test database and get a connection to it
	testDB, err := gorm.Open(postgres.Open("host=localhost user=root password=root dbname=root port=5432"))
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}
	defer testDB.Exec("DELETE FROM projects")

	// Generate a new project ID
	projectID := uuid.New()

	// Create a new instance of the repository and call the GetProject method
	repository := NewProjectPostgresRepository(NewConnector())
	_project, err := repository.GetProject(projectID)

	// Check that the returned error is not nil and that the project is nil
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if _project != nil {
		t.Fatalf("expected nil project, got %v", _project)
	}
}

func TestGetProject_DatabaseError(t *testing.T) {
	// Create a mock repository with a faulty connector
	repository := &ProjectPostgresRepository{ConnectorMock{}}

	// Call the GetProject method and check that an error is returned
	project, err := repository.GetProject(uuid.New())

	assert.Error(t, err, "an error should have been returned")
	assert.Nil(t, project, "project should be nil")
}
