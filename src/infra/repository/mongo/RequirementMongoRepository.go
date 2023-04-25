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
	"time"
)

var _ repository.RequirementLoader = &RequirementMongoRepository{}

type RequirementMongoRepository struct {
	databases.MongoDatabaseManager
}

func (r RequirementMongoRepository) GetRequirements(requirementFilters filters.RequirementFilters) ([]requirement.Requirement, error) {
	//TODO implement me
	panic("implement me")
}

func (instance RequirementMongoRepository) CreateRequirement(_requirement requirement.Requirement) (*uuid.UUID, error) {
	//TODO implement me
	panic("implement me")
}

func (instance *RequirementMongoRepository) GetRequirement(requirementID uuid.UUID) (*requirement.Requirement, error) {
	connection, err := instance.GetConnection()
	if err != nil {
		return nil, err
	}

	mongoCollection := connection.Database(Database).Collection(RequirementCollection)

	var requirementDTO dto.Requirement
	err = mongoCollection.FindOne(context.TODO(), bson.M{"_id": requirementID.String()}).Decode(&requirementDTO)
	if err != nil {
		return nil, err
	}

	_requirement := requirement.NewForDetailedView(requirementID, requirementDTO.CreatedAt,
		requirementDTO.UpdatedAt, requirementDTO.DeletedAt, requirementDTO.Name, requirementDTO.Description,
		requirementDTO.ProjectID)

	return _requirement, nil
}

func (instance *RequirementMongoRepository) DeleteRequirement(requirementID uuid.UUID) error {
	connection, err := instance.GetConnection()
	if err != nil {
		return err
	}

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

	mongoCollection := connection.Database(Database).Collection(RequirementCollection)
	fieldsToUpdate := bson.D{
		bson.E{Key: "updated_at", Value: time.Now()},
		bson.E{Key: "name", Value: _requirement.Name()},
		bson.E{Key: "description", Value: _requirement.Description()},
		bson.E{Key: "project_id", Value: _requirement.ProjectID()},
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
