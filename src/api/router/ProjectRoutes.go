package router

import (
	"github.com/felipesantoos/tcms/src/api/dicontainer"
)

func loadProjectRoutes(router *Router) {
	projectHandlers := dicontainer.GetProjectHandlers()

	router.GET(buildRoute(api, project), projectHandlers.GetProjects)
	router.GET(buildRoute(api, project, projectID), projectHandlers.GetProject)
	router.POST(buildRoute(api, project), projectHandlers.CreateProject)
	router.DELETE(buildRoute(api, project, projectID), projectHandlers.DeleteProject)
	router.PUT(buildRoute(api, project, projectID), projectHandlers.UpdateProject)
}
