package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"learning-tool/modles"
	"learning-tool/service"
	"learning-tool/tool"
)

func teacherRegister(ctx *gin.Context) {
	name := ctx.PostForm("name")
	password := ctx.PostForm("pwd")

	teacher := modles.Teachers{
		Name:        name,
		Password:    password,
		AccessLevel: "administrator", //教师为管理员身份
	}

	err := service.TeacherRegister(teacher)
	if err != nil {
		fmt.Println(err)
		tool.ReturnFailure(ctx, "注册失败")
		return
	}

	tool.ReturnSuccess(ctx, "注册成功")
}

func teacherLogin(ctx *gin.Context) {
	name := ctx.PostForm("name")
	password := ctx.PostForm("pwd")

	ctx.SetCookie("name", name, 600, "/", "localhost", false, true)

	teacher := modles.Teachers{
		Name:        name,
		Password:    password,
		AccessLevel: "administrator",
	}

	//判断密码是否正确
	truePassword, err := service.TeacherLogin(teacher)

	if err != nil {
		fmt.Println(err)
		return
	}

	if truePassword != password {
		tool.ReturnFailure(ctx, "密码错误")
		return
	}

	tool.ReturnSuccess(ctx, "登录成功")
}
