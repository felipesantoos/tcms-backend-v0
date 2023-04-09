package postgres

import (
	"fmt"
	"github.com/felipesantoos/tcms/src/core/filters"
	"github.com/felipesantoos/tcms/src/core/interfaces/repository"
	"github.com/felipesantoos/tcms/src/core/models/project"
	"github.com/felipesantoos/tcms/src/infra/interfaces/databases"
	"github.com/felipesantoos/tcms/src/infra/repository/postgres/orm"
	"github.com/google/uuid"
)

var _ repository.ProjectLoader = &ProjectPostgresRepository{}

type ProjectPostgresRepository struct {
	databases.DatabaseManager
}

func (instance *ProjectPostgresRepository) GetProjects(projectFilters filters.ProjectFilters) ([]project.Project,
	error) {
	connection, err := instance.GetConnection()
	if err != nil {
		return nil, err
	}

	projectDTOs := make([]orm.Project, 0)
	result := connection.Where("name LIKE ?", fmt.Sprintf("%%%s%%", projectFilters.Name)).
		Find(&projectDTOs)
	if result.Error != nil {
		return nil, result.Error
	}

	projectList := make([]project.Project, 0)
	for _, projectDTO := range projectDTOs {
		projectBuilder := project.NewBuilder()
		projectBuilder.ID(projectDTO.ID).Name(projectDTO.Name).Description(projectDTO.Description)
		_project, err := projectBuilder.Build()
		if err != nil {
			return nil, err
		}

		projectList = append(projectList, *_project)
	}

	return projectList, nil
}

func (instance *ProjectPostgresRepository) GetProject(projectID uuid.UUID) (*project.Project, error) {
	connection, err := instance.GetConnection()
	if err != nil {
		return nil, err
	}

	projectDTO := orm.Project{}
	result := connection.First(&projectDTO, projectID)
	if result.Error != nil {
		return nil, result.Error
	}

	projectBuilder := project.NewBuilder()
	projectBuilder.ID(projectDTO.ID).CreatedAt(projectDTO.CreatedAt).UpdatedAt(projectDTO.UpdatedAt).
		DeletedAt(projectDTO.DeletedAt.Time).Name(projectDTO.Name).Description(projectDTO.Description)
	_project, err := projectBuilder.Build()
	if err != nil {
		return nil, err
	}

	return _project, nil
}

func (instance *ProjectPostgresRepository) CreateProject(_project project.Project) (*uuid.UUID, error) {
	connection, err := instance.GetConnection()
	if err != nil {
		return nil, err
	}

	projectDTO := orm.Project{Name: _project.Name(), Description: _project.Description()}
	result := connection.Create(&projectDTO)
	if result.Error != nil {
		return nil, result.Error
	}

	return &projectDTO.ID, nil
}

func (instance *ProjectPostgresRepository) DeleteProject(projectID uuid.UUID) error {
	connection, err := instance.GetConnection()
	if err != nil {
		return err
	}

	result := connection.Delete(&orm.Project{}, projectID)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (instance *ProjectPostgresRepository) UpdateProject(_project project.Project) error {
	connection, err := instance.GetConnection()
	if err != nil {
		return err
	}

	projectDTO := orm.Project{ID: _project.ID(), Name: _project.Name(), Description: _project.Description()}
	result := connection.Model(&projectDTO).UpdateColumns(projectDTO)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func NewProjectPostgresRepository(connector databases.DatabaseManager) *ProjectPostgresRepository {
	return &ProjectPostgresRepository{connector}
}
