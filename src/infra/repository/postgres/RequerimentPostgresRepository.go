package postgres

import (
	"fmt"
	"github.com/felipesantoos/tcms/src/core/filters"
	"github.com/felipesantoos/tcms/src/core/interfaces/repository"
	"github.com/felipesantoos/tcms/src/core/models/requirement"
	"github.com/felipesantoos/tcms/src/infra/interfaces/databases"
	"github.com/felipesantoos/tcms/src/infra/repository/postgres/orm"
	"github.com/google/uuid"
)

var _ repository.RequirementLoader = &RequirementPostgresRepository{}

type RequirementPostgresRepository struct {
	databases.DatabaseManager
}

func (instance *RequirementPostgresRepository) GetRequirements(requirementFilters filters.RequirementFilters) (
	[]requirement.Requirement, error) {
	connection, err := instance.GetConnection()
	if err != nil {
		return nil, err
	}

	requirementDTOs := make([]orm.Requirement, 0)
	result := connection.Where("name LIKE ?", fmt.Sprintf("%%%s%%", requirementFilters.Name)).
		Find(&requirementDTOs)
	if result.Error != nil {
		return nil, result.Error
	}

	requirementList := make([]requirement.Requirement, 0)
	for _, requirementDTO := range requirementDTOs {
		requirementList = append(requirementList, *requirement.NewForShortView(requirementDTO.ID, requirementDTO.Name,
			requirementDTO.Description, requirementDTO.ProjectID))
	}

	return requirementList, nil
}

func (instance *RequirementPostgresRepository) GetRequirement(requirementID uuid.UUID) (*requirement.Requirement,
	error) {
	connection, err := instance.GetConnection()
	if err != nil {
		return nil, err
	}

	requirementDTO := orm.Requirement{}
	result := connection.First(&requirementDTO, requirementID)
	if result.Error != nil {
		return nil, result.Error
	}

	_requirement := requirement.NewForDetailedView(requirementDTO.ID, requirementDTO.CreatedAt,
		requirementDTO.UpdatedAt, requirementDTO.DeletedAt.Time, requirementDTO.Name, requirementDTO.Description,
		requirementDTO.ProjectID)

	return _requirement, nil
}

func (instance *RequirementPostgresRepository) CreateRequirement(_requirement requirement.Requirement) (*uuid.UUID,
	error) {
	connection, err := instance.GetConnection()
	if err != nil {
		return nil, err
	}

	requirementDTO := orm.Requirement{Name: _requirement.Name(), Description: _requirement.Description(),
		ProjectID: _requirement.ProjectID()}
	result := connection.Create(&requirementDTO)
	if result.Error != nil {
		return nil, result.Error
	}

	return &requirementDTO.ID, nil
}

func (instance *RequirementPostgresRepository) DeleteRequirement(requirementID uuid.UUID) error {
	connection, err := instance.GetConnection()
	if err != nil {
		return err
	}

	result := connection.Delete(&orm.Requirement{}, requirementID)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (instance *RequirementPostgresRepository) UpdateRequirement(_requirement requirement.Requirement) error {
	connection, err := instance.GetConnection()
	if err != nil {
		return err
	}

	requirementDTO := orm.Requirement{ID: _requirement.ID(), Name: _requirement.Name(),
		Description: _requirement.Description(), ProjectID: _requirement.ProjectID()}
	result := connection.Model(&requirementDTO).UpdateColumns(requirementDTO)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func NewRequirementPostgresRepository(connector databases.DatabaseManager) *RequirementPostgresRepository {
	return &RequirementPostgresRepository{connector}
}
