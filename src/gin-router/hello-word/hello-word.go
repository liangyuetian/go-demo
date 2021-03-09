package hello_word

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func URL() string {
	return "/hello"
}

func Get(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hello Word")
}

func Post(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hello Word")
}
