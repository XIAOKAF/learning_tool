package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"learning-tool/tool"
)

func s(ctx *gin.Context) {
	//获取Name为studentId的cookie
	studentId, err := ctx.Cookie("studentId")
	if err != nil {
		fmt.Println(err)
		tool.ReturnFailure(ctx, "请先登录")
		ctx.Abort()
	}
	ctx.Set("studentId", studentId)
	ctx.Next()
}

func t(ctx *gin.Context) {
	name, err := ctx.Cookie("name")
	if err != nil {
		fmt.Println(err)
		tool.ReturnFailure(ctx, "请先登录")
		ctx.Abort()
	}
	ctx.Set("name", name)
	ctx.Next()
}

func GetCookieName(ctx *gin.Context, name string) (string, error) {
	n, err := ctx.Cookie(name)
	if err != nil {
		return n, err
	}
	return n, nil
}
