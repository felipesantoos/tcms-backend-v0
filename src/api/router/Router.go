package router

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func New() *Router {
	return &Router{gin.Default()}
}

func (router *Router) LoadRoutes() {
	loadProjectRoutes(router)
}
