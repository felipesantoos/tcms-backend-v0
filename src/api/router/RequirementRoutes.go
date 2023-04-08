package router

import (
	"github.com/felipesantoos/tcms/src/api/dicontainer"
)

func loadRequirementRoutes(router *Router) {
	requirementHandlers := dicontainer.GetRequirementHandlers()

	router.GET(buildRoute(api, requirement), requirementHandlers.GetRequirements)
	router.GET(buildRoute(api, requirement, requirementID), requirementHandlers.GetRequirement)
	router.POST(buildRoute(api, requirement), requirementHandlers.CreateRequirement)
	router.DELETE(buildRoute(api, requirement, requirementID), requirementHandlers.DeleteRequirement)
	router.PUT(buildRoute(api, requirement, requirementID), requirementHandlers.UpdateRequirement)
}
