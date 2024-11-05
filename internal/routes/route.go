package routes

import (
	"github.com/elonghi/encurtago/internal/controller"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine, linkController *controller.LinkController) {
	router.GET("/", func(c *gin.Context) {
		linkController.Index(c)
	})

	api := router.Group("/api")
	api.GET("/encurtar/", func(c *gin.Context) {
		linkController.Shorten(c)
	})
	api.GET("/desencurtar/", func(c *gin.Context) {
		linkController.Unshorten(c)
	})
	api.POST("/encurtar/", func(c *gin.Context) {
		linkController.Shorten(c)
	})

	web := router.Group("/")
	web.GET("/:shortURL", func(c *gin.Context) {
		linkController.Redirect(c)
	})

}
