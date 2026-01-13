package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/MVMmaksM/demo-go-gin/models"
)

func ShowIndexPage(ctx *gin.Context) {
	articles := models.GetAllArticles()

	ctx.HTML(http.StatusOK, "index.html", gin.H{"title": "Home page", "payload": articles})
}
