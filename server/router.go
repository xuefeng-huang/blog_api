package server

import (
	"blog_api/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.New()

	article := new(controller.ArticleController)
	router.GET("/articles", article.GetArticle)

	return router

}
