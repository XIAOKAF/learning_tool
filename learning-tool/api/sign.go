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

var SignAmount int

func publishSign(ctx *gin.Context) {
	roomId := ctx.PostForm("roomId")               //获取发布签到教室
	currentTime := time.Now()                      //获取发布签到的时间
	holdingTime := ctx.PostForm("holdingTime")     //设置签到持续的时间,单位为分
	teacherName, err := GetCookieName(ctx, "name") //获取教师名字
	if err != nil {
		tool.ReturnFailure(ctx, "发布签到失败")
		fmt.Println(err)
		return
	}

	//设置结束签到的时间
	minute, err := time.ParseDuration(holdingTime + "m")
	if err != nil {
		tool.ReturnFailure(ctx, "设置签到时间失败")
		fmt.Println(err)
		return
	}
	endingTime := currentTime.Add(minute)

	id, err := strconv.Atoi(roomId)
	if err != nil {
		tool.ReturnFailure(ctx, "发布签到失败")
		fmt.Println(err)
		return
	}

	sign := modles.Sign{
		RoomId:      id,
		Teacher:     teacherName,
		PublishTime: currentTime,
		OverTime:    endingTime,
	}

	err = service.PublishSign(sign)
	if err != nil {
		tool.ReturnFailure(ctx, "发布签到失败")
		fmt.Println(err)
		return
	}
	tool.ReturnSuccess(ctx, "发布签到成功")
}

func sign(ctx *gin.Context)  {
	SignAmount++
	classroomId := ctx.PostForm("roomId")
	roomID, err := strconv.Atoi(classroomId)
	if err != nil {
		fmt.Println(err)
		tool.ReturnFailure(ctx,"签到失败")
		return
	}

	studentName,err := GetCookieName(ctx,"studentId")
	if err != nil {
		fmt.Println(err)
		tool.ReturnFailure(ctx,"签到失败")
		return
	}
	student := make([]string,40)
	student = append(student,studentName)

	sign := modles.Sign{
		RoomId: roomID,
		SignAmount: SignAmount,
		Student: student,
	}
	err = service.Sign(sign)
	if err != nil {
		fmt.Println(err)
		tool.ReturnFailure(ctx,"签到失败")
		return
	}
}