package routes

import (
	"media-sync-tools/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(router *gin.Engine, db *gorm.DB) {
	userGroup := router.Group("/users")
	{
		userGroup.POST("/", func(c *gin.Context) {
			// This is where logic will go!
			controllers.CreateUser(c, db)
		})
		userGroup.GET("/:id", func(c *gin.Context) {
			// This is where logic will go!
			controllers.GetUserByIDOrUsername(c, db)
		})
		userGroup.PUT("/:id", func(c *gin.Context) {
			// This is where logic will go!
			controllers.CreateUser(c, db)
		})
		userGroup.DELETE("/:id", func(c *gin.Context) {
			// This is where logic will go!
			controllers.DeleteUser(c, db)
		})
		userGroup.GET("/", func(c *gin.Context) {
			// This is where logic will go!
			controllers.GetUsers(c, db)
		})

	}
}
