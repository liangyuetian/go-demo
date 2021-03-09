package main

import (
	"github.com/kataras/iris/v12"
	HelloWord "study/src/iris-router/hello-word"
	PreOrder "study/src/iris-router/pre-order"
	Mysql "study/src/mysql"
)

func main() {
	Mysql.InitDB()
	app := iris.New()

	helloAPI := app.Party(HelloWord.URL())

	{
		helloAPI.Use(iris.Compression)

		// GET: http://localhost:8080/books
		helloAPI.Get("/", HelloWord.Get)
		// POST: http://localhost:8080/books
		helloAPI.Post("/", HelloWord.Post)
	}

	preOrderApi := app.Party(PreOrder.URL())
	preOrderApi.Get("/", PreOrder.Get)
	preOrderApi.Post("/", PreOrder.Post)
	//preOrderMac := mvc.New(preOrderApi)
	//preOrderMac.Handle(new(PreOrder.Controller))
	app.Listen(":9000")
}

//import (
//	"github.com/gin-gonic/gin"
//	GinCats "study/src/gin-router/cats"
//	GinHelloWord "study/src/gin-router/hello-word"
//)
//
//func main() {
//	r := gin.Default()
//	r.GET("/ping", func(c *gin.Context) {
//		c.JSON(200, gin.H{
//			"message": "pong",
//		})
//	})
//	r.GET(GinHelloWord.URL(), GinCats.Get)
//	r.GET(GinCats.URL(), GinCats.Get)
//	r.POST(GinCats.URL(), GinCats.Post)
//	r.Run(":9000")
//}
