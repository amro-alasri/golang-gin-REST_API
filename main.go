package main

import (
	"io"
	"net/http"
	"os"

	"github.com/amro-alasri/golangBasics/controller"
	"github.com/amro-alasri/golangBasics/middlewares"
	"github.com/amro-alasri/golangBasics/service"
	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()
	server := gin.New()

	// add the static
	server.Static("/css", "./templates/css")

	server.LoadHTMLGlob("templates/*.html")

	server.Use(gin.Recovery(), middlewares.Logger())

	apiRoutes := server.Group("/api", middlewares.BasicAuth())
	{
		apiRoutes.GET("/post", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})

		apiRoutes.POST("/post", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "Video Input is Valid!!",
				})
			}
		})
	}
	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	// We can setup this env variable from the EB console
	port := os.Getenv("PORT")
	// Elastic Beanstalk forwards requests to port 5000
	if port == "" {
		port = "5000"
	}
	server.Run(":" + port)
}
