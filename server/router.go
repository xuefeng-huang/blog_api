package server

import (
	"blog_api/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.New()

	articleController := new(controller.ArticleController)
	router.GET("/articles/:id", articleController.GetArticleById)
	router.GET("/articles", articleController.GetAllArticles)
	router.POST("/articles", articleController.CreateArticle)

	return router

}
