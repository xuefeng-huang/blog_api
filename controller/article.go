package controller

import (
	"blog_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleController struct{}

var articleModel = new(models.Article)

func (a *ArticleController) GetArticleById(c *gin.Context) {
	if c.Param("id") != "" {
		article, err := articleModel.GetArticleByID(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "error getting article", "data": nil})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Success", "data": article})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "bad request", "data": nil})
	c.Abort()
}

func (a *ArticleController) CreateArticle(c *gin.Context) {
	if err := c.ShouldBindJSON(articleModel); err == nil {
		id, err := articleModel.CreateArticle()
		if err == nil {
			c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Success", "data": gin.H{"id": id}})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "create article failed", "data": nil})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err.Error(), "data": nil})
	}
}

func (a *ArticleController) GetAllArticles(c *gin.Context) {
	articleList, err := articleModel.GetAllArticles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "get all articles failed",
			"data":    nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Success",
			"data":    articleList,
		})
	}
}
