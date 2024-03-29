package dicontainer

import (
	"github.com/felipesantoos/tcms/src/api/handlers"
)

func GetProjectHandlers() *handlers.ProjectHandlers {
	return handlers.NewProjectHandlers(GetProjectServices())
}

func GetRequirementHandlers() *handlers.RequirementHandlers {
	return handlers.NewRequirementHandlers(GetRequirementServices())
}
