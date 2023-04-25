package dicontainer

import (
	"github.com/felipesantoos/tcms/src/core/interfaces/repository"
	"github.com/felipesantoos/tcms/src/infra/repository/mongo"
	"github.com/felipesantoos/tcms/src/infra/repository/postgres"
)

func GetProjectPostgresRepository() repository.ProjectLoader {
	return postgres.NewProjectPostgresRepository(GetPostgresDatabaseManager())
}

func GetRequirementPostgresRepository() repository.RequirementLoader {
	return postgres.NewRequirementPostgresRepository(GetPostgresDatabaseManager())
}

func GetPostgresDatabaseManager() *postgres.Connector {
	return postgres.NewConnector()
}

func GetProjectMongoRepository() repository.ProjectLoader {
	return mongo.NewProjectMongoRepository(GetMongoDatabaseManager())
}

func GetRequirementMongoRepository() repository.RequirementLoader {
	return mongo.NewRequirementMongoRepository(GetMongoDatabaseManager())
}

func GetMongoDatabaseManager() *mongo.Connector {
	return mongo.NewConnector()
}
