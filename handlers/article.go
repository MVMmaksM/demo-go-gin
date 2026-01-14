package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/MVMmaksM/demo-go-gin/models"
)

func ShowIndexPage(ctx *gin.Context) {
	articles := models.GetAllArticles()

	render(ctx, gin.H{"title": "Home page", "payload": articles}, "index.html")
}

func GetDetailArticle(ctx *gin.Context) {
	articleId, err := strconv.Atoi(ctx.Param("article_id"))
	//если передан кривой id
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	article, err := models.GetArticleById(articleId)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	ctx.HTML(http.StatusOK, "article.html", gin.H{"payload": article})
}
