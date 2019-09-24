package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleController struct{}

func (a ArticleController) GetArticle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "it works!"})
}
