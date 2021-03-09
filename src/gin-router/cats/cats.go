package cats

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type PostDto struct {
	id       int
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func URL() string {
	return "/cats/:name"
}

func Get(ctx *gin.Context) {
	name := ctx.Param("name")
	id := ctx.Query("id")
	//fmt.Println(ctx.Params[0].Value)
	ctx.String(http.StatusOK, "cats param=%s id=%s", name, id)
}

func Post(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
	var m map[string]interface{}
	err = json.Unmarshal(body, &m)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(m)
	var form PostDto
	// This will infer what binder to use depending on the content-type header.
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.ShouldBind(&form)
	fmt.Println(form)
	var form2 PostDto = form
	ctx.ShouldBindJSON(&form2)
	fmt.Println(form2)
	var form3 PostDto = form
	ctx.ShouldBindQuery(&form3)
	fmt.Println(form3)
	var form4 PostDto = form
	ctx.ShouldBindHeader(&form4)
	fmt.Println(form4)
	var form5 PostDto = form
	ctx.ShouldBindUri(&form5)
	fmt.Println(form5)

	fmt.Println(ctx.PostForm("User"))
	if form.User != "manu" || form.Password != "123" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}
