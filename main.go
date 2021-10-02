package main

import (
	"fmt"
	"github.com/daisuke8000/gin_dev/controller"
	"github.com/daisuke8000/gin_dev/middlewares"
	"github.com/daisuke8000/gin_dev/service"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
	"io"
	"net/http"
	"os"
)




var (
	videoService service.VideoService = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput(){
	f, _ := os.Create("gin.log")
	fmt.Println(f)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main(){
	//server := gin.Default()
	//custom
	setupLogOutput()
	server := gin.New()

	server.Static("/css", "./templates/css")

	server.LoadHTMLGlob("templates/*.html")

	server.Use(gin.Recovery(), middlewares.Logger(),
		middlewares.BasicAuth(), gindump.Dump())
	apiRoutes := server.Group("/api")
	{
		//GET
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})

		//POST
		apiRoutes.POST("/videos", func(ctx *gin.Context){
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}else{
				ctx.JSON(http.StatusOK, gin.H{"message": "Video Input is Valid!!"})
			}
		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	//example
	//server.GET("test", func(ctx *gin.Context) {
	//	ctx.JSON(200, gin.H{
	//		"message": "Ok",
	//	})
	//})
	err := server.Run(":8080")
	if err != nil {
		return
	}
}