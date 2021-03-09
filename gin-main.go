package main

import (
	"github.com/gin-gonic/gin"
	GinCats "study/src/gin-router/cats"
	GinHelloWord "study/src/gin-router/hello-word"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET(GinHelloWord.URL(), GinCats.Get)
	r.GET(GinCats.URL(), GinCats.Get)
	r.POST(GinCats.URL(), GinCats.Post)
	r.Run(":9000")
}
