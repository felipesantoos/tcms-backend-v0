package mongo

import (
	"context"
	"github.com/felipesantoos/tcms/src/core/filters"
	"github.com/felipesantoos/tcms/src/core/interfaces/repository"
	"github.com/felipesantoos/tcms/src/core/models/requirement"
	"github.com/felipesantoos/tcms/src/infra/interfaces/databases"
	"github.com/felipesantoos/tcms/src/infra/repository/mongo/dto"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var _ repository.RequirementLoader = &RequirementMongoRepository{}

type RequirementMongoRepository struct {
	databases.MongoDatabaseManager
}

func (instance *RequirementMongoRepository) GetRequirements(requirementFilters filters.RequirementFilters) ([]requirement.Requirement, error) {
	connection, err := instance.GetConnection()
	if err != nil {
		return nil, err
	}

	mongoCollection := connection.Database(Database).Collection(RequirementCollection)

	requirementFiltersInMongoFormat := bson.M{}
	if requirementFilters.Name != "" {
		requirementFiltersInMongoFormat["name"] = bson.M{"$regex": primitive.Regex{Pattern: requirementFilters.Name, Options: "i"}}
	}

	result, err := mongoCollection.Find(context.TODO(), requirementFiltersInMongoFormat)
	if err != nil {
		return nil, err
	}

	var requirementDTOs []dto.Requirement
	err = result.All(context.TODO(), &requirementDTOs)
	if err != nil {
		return nil, err
	}

	requirementList := make([]requirement.Requirement, 0)
	for _, requirementDTO := range requirementDTOs {
		/*if requirementDTO.IsActive == false {
			continue
		}*/

		requirementUuid, err := uuid.Parse(requirementDTO.ID)
		if err != nil {
			return nil, err
		}

		projectUuid, err := uuid.Parse(requirementDTO.ProjectID)
		if err != nil {
			return nil, err
		}

		requirementBuilder := requirement.NewBuilder()
		requirementBuilder.ID(requirementUuid).Name(requirementDTO.Name).Description(requirementDTO.Description).ProjectId(projectUuid)
		_requirement, err := requirementBuilder.Build()
		if err != nil {
			return nil, err
		}

		requirementList = append(requirementList, *_requirement)
	}

	return requirementList, nil
}

func (instance *RequirementMongoRepository) GetRequirement(requirementID uuid.UUID) (*requirement.Requirement, error) {
	//TODO implement me
	panic("implement me")
}

func (instance *RequirementMongoRepository) CreateRequirement(_requirement requirement.Requirement) (*uuid.UUID, error) {
	connection, err := instance.GetConnection()
	if err != nil {
		return nil, err
	}

	mongoCollection := connection.Database(Database).Collection(RequirementCollection)

	requirementUuid := uuid.New()

	requirementDTO := dto.NewRequirement(requirementUuid, time.Now(), time.Now(), time.Time{}, _requirement.Name(), _requirement.Description(), _requirement.ProjectID())
	_, err = mongoCollection.InsertOne(context.TODO(), requirementDTO)
	if err != nil {
		return nil, err
	}

	return &requirementUuid, nil
}

func (instance *RequirementMongoRepository) DeleteRequirement(requirementID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (instance *RequirementMongoRepository) UpdateRequirement(_requirement requirement.Requirement) error {
	//TODO implement me
	panic("implement me")
}

func NewRequirementMongoRepository(connector databases.MongoDatabaseManager) *RequirementMongoRepository {
	return &RequirementMongoRepository{connector}
}
