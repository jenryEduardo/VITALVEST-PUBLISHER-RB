package routes

import (
	"publisher/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(r *gin.Engine) {
	routes := r.Group("/AMQP")

	routes.POST("/",controllers.GetItems)

}