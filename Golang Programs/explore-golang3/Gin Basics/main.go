package main

import (
	"ginbasics/controller"
	"ginbasics/middlewares"
	"ginbasics/service"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
}

func main() {
	server := gin.Default()

	setupLogOutput()

	server.Use(gin.Recovery())

	server.Use(middlewares.Logger())

	server.Use(middlewares.BasicAuth())

	// request body and all that.
	server.Use(gindump.Dump())

	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{ // H is shortcut for map string string, converts to json and also sets application type json
			"message": "Hello World!",
		})
	})

	server.GET("/video", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})


	server.POST("/video", func(ctx *gin.Context) {
		err := videoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"status": "success"})
		}
	})

	server.Run(":8080")
}
