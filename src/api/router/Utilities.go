package router

import (
	"fmt"
)

const (
	api           = "/api"
	project       = "/project"
	projectID     = "/:projectID"
	requirement   = "/requirement"
	requirementID = "/:requirementID"
)

func buildRoute(segments ...string) string {
	route := ""
	for _, segment := range segments {
		route = fmt.Sprintf("%s%s", route, segment)
	}

	return route
}
