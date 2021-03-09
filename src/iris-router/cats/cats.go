package cats

import (
	"github.com/kataras/iris/v12"
)

type PostDto struct {
	id       int
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func URL() string {
	return "/cats/:name"
}

func Get(ctx iris.Context) {
	//fmt.Println(ctx.Params[0].Value)
	ctx.StatusCode(iris.StatusOK)
	ctx.WriteString("hello word")
}

//
//func Post(ctx *gin.Context) {
//	var form PostDto
//	// This will infer what binder to use depending on the content-type header.
//	if err := ctx.ShouldBind(&form); err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	ctx.ShouldBind(&form)
//	fmt.Println(form)
//	var form2 PostDto = form
//	ctx.ShouldBindJSON(&form2)
//	fmt.Println(form2)
//	var form3 PostDto = form
//	ctx.ShouldBindQuery(&form3)
//	fmt.Println(form3)
//	var form4 PostDto = form
//	ctx.ShouldBindHeader(&form4)
//	fmt.Println(form4)
//	var form5 PostDto = form
//	ctx.ShouldBindUri(&form5)
//	fmt.Println(form5)
//
//	fmt.Println(ctx.PostForm("User"))
//	if form.User != "manu" || form.Password != "123" {
//		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
//		return
//	}
//	ctx.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
//}
