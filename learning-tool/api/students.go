package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"learning-tool/modles"
	"learning-tool/service"
	"learning-tool/tool"
	"strconv"
	"time"
)

//注册
func register(ctx *gin.Context) {
	studentId := ctx.PostForm("studentId")
	realname := ctx.PostForm("realname")
	password := ctx.PostForm("pwd")

	id, err := strconv.Atoi(studentId)
	if err != nil {
		fmt.Println("学号转换错误:")
		fmt.Println(err)
		tool.ReturnFailure(ctx, "转化错误")
		return
	}

	student := modles.Students{
		StudentId:   id,
		RealName:    realname,
		Password:    password,
		AccessLevel: "Normal User",
	}

	//学号是否存在
	flag, err := service.IsStudentExist(student)

	if flag != true {
		if err == nil {
			fmt.Println(err)
			tool.ReturnFailure(ctx, "The student id does not exist.")
			return
		}
		fmt.Println(err)
		return
	}

	//学号与姓名是否匹配
	name, err := service.IsIdAndNameMatch(student)
	if err != nil {
		fmt.Println(err)
		return
	}
	if name != realname {
		tool.ReturnFailure(ctx, "学号与真实姓名不匹配")
		return
	}

	//学号存在并且真实姓名正确就进行注册
	err = service.Register(student)
	if err != nil {
		fmt.Println(err)
		return
	}
	tool.ReturnSuccess(ctx, "Register successfully!")
}

//登录
func login(ctx *gin.Context) {
	studentId := ctx.PostForm("studentId")
	password := ctx.PostForm("pwd")
	ctx.SetCookie("studentId", studentId, 600, "/", "localhost", false, true)

	id, err := strconv.Atoi(studentId)
	if err != nil {
		fmt.Println(err)
		return
	}

	students := modles.Students{
		StudentId: id,
		Password:  password,
	}

	//判断密码是否正确
	truePassword, err := service.IsPasswordCorrect(students)

	if err != nil {
		fmt.Println(err)
		return
	}

	if truePassword != password {
		tool.ReturnFailure(ctx, "Wrong password!")
		return
	}

	tool.ReturnSuccess(ctx, "Login successfully!")
}

//完善个人信息
func completeInfo(ctx *gin.Context) {
	//获取cookie的值，作为后续索引值
	studentId, err := GetCookieName(ctx, "studentId")
	if err != nil {
		fmt.Println(err)
		return
	}

	id, err := strconv.Atoi(studentId)
	if err != nil {
		tool.ReturnFailure(ctx,"更新个人信息失败")
		fmt.Println("学号格式转换失败:")
		fmt.Println(err)
		return
	}

	nickname := ctx.PostForm("nickname")
	mobile := ctx.PostForm("mobile")

	phone, err := strconv.Atoi(mobile)
	if err != nil {
		tool.ReturnFailure(ctx,"更新个人信息失败")
		fmt.Println("电话号码格式转换失败:")
		fmt.Println(err)
		return
	}

	if len(mobile) != 11 {
		tool.ReturnFailure(ctx, "电话号码格式错误")
		return
	}

	//头像上传
	//1、解析上传的参数:file
	//2、file保存到本地
	//3、更新用户表中的头像字段（文件的路径）
	//4、返回结果
	file, err := ctx.FormFile("avatar")
	if err != nil {
		fmt.Println(err)
		tool.ReturnFailure(ctx, "头像更新失败")
		return
	}

	//存放文件的目录+文件的命名（避免文件名重复）
	fileName := "./uploadfile/" + strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	//保存文件
	err = ctx.SaveUploadedFile(file, fileName)
	if err != nil {
		tool.ReturnFailure(ctx, "头像更新失败")
		return
	}

	//http://localhost:8080/static/.../Lee.png

	//实例化student
	student := modles.Students{
		StudentId: id,
		Mobile:    phone,
		NickName:  nickname,
		Avatar:    fileName[1:],
	}

	err = service.UpdateInfo(student)
	if err != nil {
		tool.ReturnFailure(ctx, "个人信息更新失败")
		fmt.Println(err)
		return
	}
	tool.ReturnSuccess(ctx, "更新成功")
}
