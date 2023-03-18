// Routes/Routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/marceloamoreno87/gomail/publisher/api/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	EmailController := new(controllers.EmailController)

	grp := r.Group("/api")
	{
		grp.POST("/mail", EmailController.Store)
	}

	return r
}
