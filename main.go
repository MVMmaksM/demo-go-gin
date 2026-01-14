package main

import (
	"log"

	"github.com/MVMmaksM/demo-go-gin/handlers"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", handlers.ShowIndexPage)
	router.GET("/article/view/:article_id", handlers.GetDetailArticle)

	err := router.Run(":9000")
	if err != nil {
		log.Fatal(err)
	}
}
