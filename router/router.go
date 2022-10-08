package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/person/CRUD-api/handlers"
	"github.com/person/CRUD-api/services"
	"github.com/person/CRUD-api/utils"
)

func InitializeRouter(services *services.Services) *gin.Engine {

	// Set the default gin router
	r := gin.New()

	r.Use(gin.Recovery())

	// Initialize middlewares
	initializeMiddlewares(r)

	// Initialize routes
	initializeRoutes(r, services)

	return r

}

func initializeRoutes(r *gin.Engine, services *services.Services) {
	untracedGroup := r.Group("/")
	untracedGroup.Use(Ginzap(utils.Logger, time.RFC3339, true, false))

	// health
	untracedGroup.GET("/list", handlers.ListPerson(services))
	untracedGroup.POST("/add", handlers.AddPerson(services))
	untracedGroup.PUT("/update/:index", handlers.UpdatePerson(services))
	untracedGroup.DELETE("/delete/:index", handlers.DeletePerson(services))

	// fallback
	r.NoRoute(handlers.NoRoute)
}
