package main

import (
	 controller "GoginRestApi/controller"
	_"GoginRestApi/docs"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Go Gin Swagger Rest API service demo 
// @version 1.0
// @description This is sample server.
// @host localhost:8080
// @BasePath /api/v1


func main() {
	r:= setupRouter()
	_= r.Run(":8080")
	log.Println("router")
}

func setupRouter() *gin.Engine {

	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Welcome To Go Gin Swagger Rest API Sample"})
	})
	
	
	v1 := r.Group("/api/v1")
	{
		articles := v1.Group("/article")
		{
			articles.POST("/create", controller.CreateArticle)
			articles.GET("/detail/:id", controller.GetDetail)
			articles.GET("/list", controller.ListArticles)

		}
	}
	

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r

}
