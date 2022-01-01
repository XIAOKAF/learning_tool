package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"learning-tool/modles"
	"learning-tool/service"
	"learning-tool/tool"
	"strconv"
)

var students []string //定义存学生姓名的切片

func CreateClassroom(ctx *gin.Context) {
	classroomName := ctx.PostForm("classroomName")
	capacity := ctx.PostForm("capacity")

	teacherName, err := GetCookieName(ctx, "name")
	if err != nil {
		tool.ReturnFailure(ctx, "获取姓名失败")
		fmt.Println(err)
		return
	}

	permission, err := service.ReturnPermission(teacherName)
	if err != nil {
		tool.ReturnFailure(ctx, "创建教室失败")
		fmt.Println(err)
		return
	}
	if permission != "administrator" {
		tool.ReturnFailure(ctx, "无权限创建教室")
		return
	}

	c, err := strconv.Atoi(capacity)
	if err != nil {
		tool.ReturnFailure(ctx, "创建教室失败")
		fmt.Println(err)
		return
	}

	classroom := modles.Classroom{
		RoomName: classroomName,
		Capacity: c,
		Teacher:  teacherName,
	}

	err = service.CreateClassroom(classroom)
	if err != nil {
		tool.ReturnFailure(ctx, "创建教室失败")
		fmt.Println(err)
		return
	}

	tool.ReturnSuccess(ctx, "成功创建教室")
}

func JoinClassroom(ctx *gin.Context) {
	classroomId := ctx.PostForm("classroomId")
	studentId, err := GetCookieName(ctx, "studentId")

	roomId, err := strconv.Atoi(classroomId)
	if err != nil {
		fmt.Println(err)
		return
	}

	Id, err := strconv.Atoi(studentId)
	if err != nil {
		fmt.Println(err)
		return
	}

	studentName, err := service.GetStudentName(Id) //通过学号找到学生姓名
	if err != nil {
		fmt.Println(err)
		tool.ReturnFailure(ctx, "获取学生名字失败")
	}
	students = append(students, studentName) //将学生姓名存入切片

	classroom := modles.Classroom{
		RoomId:  roomId,
		Members: students,
	}

	flag, err := service.JoinClassroom(classroom) //加入教室
	if err != nil {
		fmt.Println(err)
		tool.ReturnFailure(ctx, "加入教室失败")
		return
	}
	if flag != 1 {
		tool.ReturnSuccess(ctx, "成功加入教室:"+classroomId)
		return
	}
	tool.ReturnSuccess(ctx, "成功加入教室:"+classroomId+"老师已经发布签到了，不要错过了哦")
}
