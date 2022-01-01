package api

import (
	"github.com/gin-gonic/gin"
)

func InitEngine() {
	engine := gin.Default()

	//学生
	student := engine.Group("student")
	{
		student.POST("register", register)                    //注册
		student.POST("login", login)                          //登录
		student.Use(s)                                        //记住登录状态
		student.POST("/completeInfo", completeInfo)           //完善个人信息
		student.POST("/joinClassroom", JoinClassroom)         //加入教室
		student.GET("downloadAssignment", downloadAssignment) //下载作业/课件
		student.POST("submitAssignment", SubmitAssignment)    //提交作业
		student.POST("/sign", sign)                           //签到
	}

	//教师
	teacher := engine.Group("teacher")
	{
		teacher.POST("teacherRegister", teacherRegister)  //注册
		teacher.POST("login", teacherLogin)               //登录
		teacher.Use(t)                                    //记住登录状态
		teacher.POST("/createClassroom", CreateClassroom) //仅教师有权限创建教室
		teacher.POST("/assignHomework", assignHomework)   //布置作业/发布课件
		teacher.POST("/publishSign", publishSign)         //发布签到
		//仅教师有权限管理学生成绩
		teacher.POST("/createReportCard", createReportCard) //创建成绩单
		teacher.POST("/loginData", loginData)               //录入学生成绩
		teacher.GET("/averageScore", averageScore)          //返回平均分
		teacher.GET("highestScore", highestScore)           //返回最高分
		teacher.GET("specifiedStudent", specifiedStudent)   //查询指定学生的信息
	}

	//其他功能

	engine.POST("/uploadFiles")  //上传文件
	engine.GET("/downloadFiles") //下载文件

	engine.Run(":8090")
}
