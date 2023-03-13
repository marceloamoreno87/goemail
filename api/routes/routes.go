// Routes/Routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/marceloamoreno87/gomail/api/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp := r.Group("/api")
	{
		grp.POST("/send-email", controllers.SendEmail)
	}

	return r
}
