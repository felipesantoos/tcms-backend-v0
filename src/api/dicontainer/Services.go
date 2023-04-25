package dicontainer

import (
	"github.com/felipesantoos/tcms/src/core/interfaces/usecases"
	"github.com/felipesantoos/tcms/src/core/services"
)

func GetProjectServices() usecases.ProjectManager {
	return services.NewProjectServices(GetProjectMongoRepository())
}

func GetRequirementServices() usecases.RequirementManager {
	return services.NewRequirementServices(GetRequirementMongoRepository())
}
