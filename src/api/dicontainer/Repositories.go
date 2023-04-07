package dicontainer

import (
	"github.com/felipesantoos/tcms/src/core/interfaces/repository"
	"github.com/felipesantoos/tcms/src/infra/repository/postgres"
)

func GetProjectPostgresRepository() repository.ProjectLoader {
	return postgres.NewProjectPostgresRepository(GetDatabaseManager())
}

func GetDatabaseManager() *postgres.Connector {
	return postgres.NewConnector()
}