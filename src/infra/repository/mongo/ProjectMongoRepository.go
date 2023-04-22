package mongo

import (
	"context"
	"github.com/felipesantoos/tcms/src/core/filters"
	"github.com/felipesantoos/tcms/src/core/interfaces/repository"
	"github.com/felipesantoos/tcms/src/core/models/project"
	"github.com/felipesantoos/tcms/src/infra/interfaces/databases"
	"github.com/felipesantoos/tcms/src/infra/repository/mongo/dto"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var _ repository.ProjectLoader = &ProjectMongoRepository{}

type ProjectMongoRepository struct {
	databases.MongoDatabaseManager
}

func (instance *ProjectMongoRepository) GetProjects(projectFilters filters.ProjectFilters) ([]project.Project,
	error) {
	connection, err := instance.GetConnection()
	if err != nil {
		return nil, err
	}

	mongoCollection := connection.Database(Database).Collection(ProjectCollection)

	projectFiltersInMongoFormat := bson.M{}
	if projectFilters.Name != "" {
		projectFiltersInMongoFormat["name"] = bson.M{"$regex": primitive.Regex{Pattern: projectFilters.Name, Options: "i"}}
	}

	result, err := mongoCollection.Find(context.TODO(), projectFiltersInMongoFormat)
	if err != nil {
		return nil, err
	}

	var projectDTOs []dto.Project
	err = result.All(context.TODO(), &projectDTOs)
	if err != nil {
		return nil, err
	}

	projectList := make([]project.Project, 0)
	for _, projectDTO := range projectDTOs {
		if projectDTO.IsActive == false {
			continue
		}

		projectUuid, err := uuid.Parse(projectDTO.ID)
		if err != nil {
			return nil, err
		}

		projectBuilder := project.NewBuilder()
		projectBuilder.ID(projectUuid).Name(projectDTO.Name).Description(projectDTO.Description)
		_project, err := projectBuilder.Build()
		if err != nil {
			return nil, err
		}

		projectList = append(projectList, *_project)
	}

	return projectList, nil
}

func (instance *ProjectMongoRepository) GetProject(projectID uuid.UUID) (*project.Project, error) {
	connection, err := instance.GetConnection()
	if err != nil {
		return nil, err
	}

	mongoCollection := connection.Database(Database).Collection(ProjectCollection)

	var projectDTO dto.Project
	err = mongoCollection.FindOne(context.TODO(), bson.M{"_id": projectID.String()}).Decode(&projectDTO)
	if err != nil {
		return nil, err
	}

	projectUuid, err := uuid.Parse(projectDTO.ID)
	if err != nil {
		return nil, err
	}

	projectBuilder := project.NewBuilder()
	projectBuilder.ID(projectUuid).CreatedAt(projectDTO.CreatedAt).UpdatedAt(projectDTO.UpdatedAt).
		DeletedAt(projectDTO.DeletedAt).Name(projectDTO.Name).Description(projectDTO.Description)
	_project, err := projectBuilder.Build()
	if err != nil {
		return nil, err
	}

	return _project, nil
}

func (instance *ProjectMongoRepository) CreateProject(_project project.Project) (*uuid.UUID, error) {
	connection, err := instance.GetConnection()
	if err != nil {
		return nil, err
	}

	mongoCollection := connection.Database(Database).Collection(ProjectCollection)

	productUuid := uuid.New()
	projectDTO := dto.NewProject(productUuid, time.Now(), time.Now(), time.Time{}, _project.Name(), _project.Description(), true)
	_, err = mongoCollection.InsertOne(context.TODO(), projectDTO)
	if err != nil {
		return nil, err
	}

	return &productUuid, nil
}

func (instance *ProjectMongoRepository) DeleteProject(projectID uuid.UUID) error {
	connection, err := instance.GetConnection()
	if err != nil {
		return err
	}

	mongoCollection := connection.Database(Database).Collection(ProjectCollection)
	fieldsToUpdate := bson.D{
		bson.E{Key: "deleted_at", Value: time.Now()},
		bson.E{Key: "is_active", Value: false},
	}
	_, err = mongoCollection.UpdateOne(context.TODO(), bson.M{"_id": projectID.String()}, bson.D{{"$set", fieldsToUpdate}})
	if err != nil {
		return err
	}

	return nil
}

func (instance *ProjectMongoRepository) UpdateProject(_project project.Project) error {
	connection, err := instance.GetConnection()
	if err != nil {
		return err
	}

	mongoCollection := connection.Database(Database).Collection(ProjectCollection)
	fieldsToUpdate := bson.D{
		bson.E{Key: "updated_at", Value: time.Now()},
		bson.E{Key: "name", Value: _project.Name()},
		bson.E{Key: "description", Value: _project.Description()},
	}
	_, err = mongoCollection.UpdateOne(context.TODO(), bson.M{"_id": _project.ID().String()}, bson.D{{"$set", fieldsToUpdate}})
	if err != nil {
		return err
	}

	return nil
}

func NewProjectMongoRepository(connector databases.MongoDatabaseManager) *ProjectMongoRepository {
	return &ProjectMongoRepository{connector}
}
