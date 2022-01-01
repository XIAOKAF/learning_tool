package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"learning-tool/modles"
	"learning-tool/service"
	"learning-tool/tool"
	"strconv"
)

func createReportCard(ctx *gin.Context) {
	subject := ctx.PostForm("subject")
	teacher, err := GetCookieName(ctx, "name")
	if err != nil {
		fmt.Println("获取教师名字失败:", err)
		tool.ReturnFailure(ctx, "创建成绩单失败")
		return
	}
	grades := modles.Grades{
		Subject: subject,
		Teacher: teacher,
	}
	err = service.CreateReportCard(grades)
	if err != nil {
		fmt.Println("创建成绩单失败", err)
		tool.ReturnFailure(ctx, "创建成绩单失败")
		return
	}
	tool.ReturnSuccess(ctx, "成功创建成绩单")
}

func loginData(ctx *gin.Context) {
	teacher, err := GetCookieName(ctx, "name")
	if err != nil {
		fmt.Println("获取教师姓名失败:", err)
		tool.ReturnFailure(ctx, "录入信息失败")
		return
	}
	subject := ctx.PostForm("subject")
	studentId := ctx.PostForm("studentId")
	studentName := ctx.PostForm("studentName")
	scores := ctx.PostForm("scores")

	id, err := strconv.Atoi(studentId)
	if err != nil {
		fmt.Println("学号转换失败:", err)
		tool.ReturnFailure(ctx, "录入信息失败")
		return
	}
	s, err := strconv.ParseFloat(scores, 32)
	if err != nil {
		fmt.Println("姓名转换失败:", err)
		tool.ReturnFailure(ctx, "录入信息失败")
		return
	}

	grades := modles.Grades{
		Subject:     subject,
		Teacher:     teacher,
		StudentId:   id,
		StudentName: studentName,
		Scores:      float32(s),
	}

	err = service.InsertGrades(grades)
	if err != nil {
		fmt.Println("录入成绩失败:", err)
		tool.ReturnFailure(ctx, "录入信息失败")
		return
	}
	tool.ReturnSuccess(ctx, "成功录入成绩")
}

func averageScore(ctx *gin.Context) {
	teacher, err := GetCookieName(ctx, "name")
	if err != nil {
		fmt.Println("获取教师姓名失败:", err)
		tool.ReturnFailure(ctx, "返回平均分失败")
		return
	}
	subject := ctx.PostForm("subject")

	grades := modles.Grades{
		Subject: subject,
		Teacher: teacher,
	}

	averageScore, err := service.AverageScores(grades)
	if err != nil {
		fmt.Println("计算平均分失败:", err)
		tool.ReturnFailure(ctx, "返回平均分失败")
		return
	}
	avg := strconv.FormatFloat(float64(averageScore), 'E', 2, 64)
	tool.ReturnSuccess(ctx, "平均分是:"+avg)
}

func highestScore(ctx *gin.Context) {
	teacher, err := GetCookieName(ctx, "name")
	if err != nil {
		fmt.Println("获取教师姓名失败:", err)
		tool.ReturnFailure(ctx, "返回最高分失败")
		return
	}
	subject := ctx.PostForm("subject")

	grades := modles.Grades{
		Subject: subject,
		Teacher: teacher,
	}

	maxScore, student, err := service.HighestScore(grades)
	if err != nil {
		fmt.Println(err)
		tool.ReturnFailure(ctx, "获取"+subject+"科目中的最高分失败")
		return
	}
	max := strconv.FormatFloat(float64(maxScore), 'E', 2, 64)
	tool.ReturnSuccess(ctx, student+"取得"+subject+"的最高分:"+max)
}

func specifiedStudent(ctx *gin.Context) {
	subject := ctx.PostForm("subject") //查询科目
	studentId := ctx.PostForm("studentId")
	studentName := ctx.PostForm("studentName")
	id, err := strconv.Atoi(studentId)
	if err != nil {
		fmt.Println("学号转化错误:", err)
		tool.ReturnFailure(ctx, "查询失败")
		return
	}

	grades := modles.Grades{
		Subject:     subject,
		StudentId:   id,
		StudentName: studentName,
		Scores:      0,
	}
	grades, err = service.SpecifiedStudent(grades)
	if err != nil {
		fmt.Println("查询失败", err)
		tool.ReturnFailure(ctx, "查询成绩失败")
		return
	}
	s := strconv.FormatFloat(float64(grades.Scores), 'E', 2, 64)
	tool.ReturnSuccess(ctx, studentName+"的"+subject+"成绩是"+s)
}
