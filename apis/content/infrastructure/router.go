package infrastructure

import (
	"github.com/gin-contrib/cors"
	gin "github.com/gin-gonic/gin"
	"github.com/wakatakeru/hiroba/apis/content/interfaces/controllers"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowHeaders = []string{
		"Content-Type",
		"Authorization",
	}
	router.Use(cors.New(config))

	contentController := controllers.NewContentController(NewSqlHandler(), NewJWTHandler())

	// GET Endpoints
	router.GET("/contents/:id", func(c *gin.Context) { contentController.Show(c) })
	router.GET("/contents/:site_id", func(c *gin.Context) { contentController.SiteIndex(c) })
	router.GET("/contents/:user_id", func(c *gin.Context) { contentController.UserIndex(c) })

	// POST Endpoints
	router.POST("/contents", func(c *gin.Context) { contentController.Create(c) })

	Router = router
}
