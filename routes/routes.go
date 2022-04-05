package routes

import (
	"jabar-nearby-places/middlewares"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine = gin.Default()

func SetupRouter() *gin.Engine {
	r.Use(middlewares.ErrorMiddleware())

	r.Static("/docs/", "dist")
	return r
}
