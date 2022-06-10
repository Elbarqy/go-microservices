package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {
	router := gin.Default()
	address := ":3000"
	v1 := router.Group("/api")
	imageGroup := v1.Group("/images")
	v2 := v1.Group("/user")
	v2.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello world")
	})
	v2.POST("/login", func(ctx *gin.Context) {
		var body Login
		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		if body.User != "manu" || body.Password != "123" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		println("passed")
		ctx.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
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
