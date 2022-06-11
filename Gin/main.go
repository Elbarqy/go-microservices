package main

import (
	"log"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required,password"`
}

func verifyPassword(fl validator.FieldLevel) bool {
	var regex = regexp.MustCompile("\\w{8,}")
	password := fl.Field().String()
	return regex.MatchString(password)
}
func main() {
	router := gin.Default()
	address := ":3000"
	v1 := router.Group("/api")
	imageGroup := v1.Group("/images")
	v2 := v1.Group("/user")
	v3 := router.Group("/files") //Static , StaticFS, StaticFile
	v2.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello world")
	})

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("password", verifyPassword)
	}

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

	v3.Static("/assets", "./assets") //localhost:3000/assets/test.jpg
	log.Fatalln(router.Run(address))
}
