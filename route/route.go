package route

import (
	"net/http"

	"example.com/m/controller"
	"example.com/m/models"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		// models.DB.AutoMigrate(&models.Movies{})
		v1.GET("/users", func(ctx *gin.Context) {
			var f *[]models.User = controller.AllUsers()
			ctx.JSON(http.StatusOK, gin.H{"data": f})
		})

		v1.GET("/user/:id", func(ctx *gin.Context) {
			var f *models.User = controller.FindUser()
			ctx.JSON(http.StatusOK, gin.H{"data": f})
		})

		v1.POST("/users", controller.NewUser)
		v1.DELETE("/user/:id", controller.DeleteUser) // new
	}
	return r
}
