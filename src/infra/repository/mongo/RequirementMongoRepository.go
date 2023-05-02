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
	defer connection.Disconnect(context.TODO())

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
	connection, err := instance.GetConnection()
	if err != nil {
		return nil, err
	}
	defer connection.Disconnect(context.TODO())

	mongoCollection := connection.Database(Database).Collection(RequirementCollection)

	var requirementDTO dto.Requirement
	err = mongoCollection.FindOne(context.TODO(), bson.M{"_id": requirementID.String()}).Decode(&requirementDTO)
	if err != nil {
		return nil, err
	}

	projectUuid, err := uuid.Parse(requirementDTO.ProjectID)
	if err != nil {
		return nil, err
	}

	_requirement := requirement.NewForDetailedView(requirementID, requirementDTO.CreatedAt,
		requirementDTO.UpdatedAt, requirementDTO.DeletedAt, requirementDTO.Name, requirementDTO.Description,
		projectUuid)

	return _requirement, nil
}

func (instance *RequirementMongoRepository) CreateRequirement(_requirement requirement.Requirement) (*uuid.UUID, error) {
	connection, err := instance.GetConnection()
	if err != nil {
		return nil, err
	}
	defer connection.Disconnect(context.TODO())

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
	connection, err := instance.GetConnection()
	if err != nil {
		return err
	}
	defer connection.Disconnect(context.TODO())

	mongoCollection := connection.Database(Database).Collection(RequirementCollection)

	_, err = mongoCollection.DeleteOne(context.TODO(), bson.M{"_id": requirementID.String()})
	if err != nil {
		return err
	}
	return nil
}

func (instance *RequirementMongoRepository) UpdateRequirement(_requirement requirement.Requirement) error {
	connection, err := instance.GetConnection()
	if err != nil {
		return err
	}
	defer connection.Disconnect(context.TODO())

	mongoCollection := connection.Database(Database).Collection(RequirementCollection)
	fieldsToUpdate := bson.D{
		bson.E{Key: "updated_at", Value: time.Now()},
		bson.E{Key: "name", Value: _requirement.Name()},
		bson.E{Key: "description", Value: _requirement.Description()},
		bson.E{Key: "project_id", Value: _requirement.ProjectID().String()},
	}
	_, err = mongoCollection.UpdateOne(context.TODO(), bson.M{"_id": _requirement.ID().String()}, bson.D{{"$set", fieldsToUpdate}})
	if err != nil {
		return err
	}

	return nil
}

func NewRequirementMongoRepository(connector databases.MongoDatabaseManager) *RequirementMongoRepository {
	return &RequirementMongoRepository{connector}
}
