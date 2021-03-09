package hello_word

import (
	"github.com/kataras/iris/v12"
)

func URL() string {
	return "/hello"
}

func Get(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	ctx.WriteString("hello word")
}
func Post(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	ctx.WriteString("hello word")
}
