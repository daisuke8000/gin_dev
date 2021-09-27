package main

import (
	"github.com/daisuke8000/gin_dev/controller"
	"github.com/daisuke8000/gin_dev/service"
	"github.com/gin-gonic/gin"
)


var (
	videoService service.VideoService = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main(){
	server := gin.Default()

	//GET
	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	//POST
	server.POST("/videos", func(ctx *gin.Context){
		ctx.JSON(200, videoController.Save(ctx))
	})

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