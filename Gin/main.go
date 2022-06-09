package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	address := ":3000"
	v1 := router.Group("/api")
	imageGroup := v1.Group("/images")
	v2 := v1.Group("/user")
	v2.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello world")
	})
	v2.GET("/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		query := ctx.Query("query")
		ctx.String(http.StatusOK, id+" "+query)
	})
	imageGroup.POST("/upload", func(ctx *gin.Context) {
		text := ctx.PostForm("text")
		ctx.String(http.StatusOK, text)
	})
	log.Fatalln(router.Run(address))
}
