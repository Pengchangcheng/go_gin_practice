package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	ginDump "github.com/tpkeeper/gin-dump"
	"pcc.com/golangTest/golang-gin-poc/controller"
	"pcc.com/golangTest/golang-gin-poc/middlewares"
	"pcc.com/golangTest/golang-gin-poc/service"
)

var (
	vedioService    service.IVideoService      = service.New()
	vedioController controller.VideoController = controller.New(vedioService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLogOutput()

	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger(),
		middlewares.BasicAuth(), ginDump.Dump())

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, vedioController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		err := vedioController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Vedio input is Valid!!"})
		}
	})

	server.Run(":8080")
}
