package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"learning-tool/modles"
	"learning-tool/service"
	"learning-tool/tool"
	"os"
	"strconv"
	"time"
)

func assignHomework(ctx *gin.Context) {
	roomId := ctx.PostForm("roomId")
	//通过cookie获取老师的姓名
	teacherName, err := GetCookieName(ctx, "name")
	if err != nil {
		tool.ReturnFailure(ctx, "可惜没有作业啦")
		fmt.Println(err)
		return
	}
	//获取作业，文档或图片
	file, err := ctx.FormFile("assignment")
	if err != nil {
		fmt.Println(err)
		tool.ReturnFailure(ctx, "哦豁")
		return
	}
	//命名并保存文件
	fileName := "./uploadfile/" + strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	err = ctx.SaveUploadedFile(file, fileName)
	if err != nil {
		fmt.Println(err)
		tool.ReturnFailure(ctx, "作业保存失败")
	}

	//定义集合作业的切片
	id, err := strconv.Atoi(roomId)
	if err != nil {
		tool.ReturnFailure(ctx, "没有作业啦啦啦啦")
		fmt.Println(err)
		return
	}
	//实例化作业结构体
	assignment := modles.Assignment{
		RoomId:   id,
		Teacher:  teacherName,
		Homework: fileName[1:],
	}

	err = service.AssignHomework(assignment)
	if err != nil {
		fmt.Println(err)
		tool.ReturnFailure(ctx, "失败")
		return
	}
	tool.ReturnSuccess(ctx, "成功布置作业")
}

func downloadAssignment(ctx *gin.Context) {
	/* 1、错误
	filePath := ctx.Param("/uploadfile/zuoye")
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		tool.ReturnFailure(ctx, "下载作业失败")
		return
	}

	defer file.Close()
	ctx.Writer.Header().Add("content-type", "application/octet-stream")
	_, err = io.Copy(ctx.Writer, file)
	if err != nil {
		fmt.Println(err)
		tool.ReturnFailure(ctx, "下载作业失败")
		return
	}
	tool.ReturnSuccess(ctx, "作业来啦!")
	*/
	/* 2、错误
	filename := ctx.DefaultQuery("filename","")
	ctx.Writer.Header().Add("Content-Disposition",fmt.Sprintf("attachment;filename=%s", filename))
	ctx.Writer.Header().Add("Content-Type","application/octet-stream")
	ctx.File("./uploadfile/zuoye.docx")
	*/

	file, header, err := ctx.Request.FormFile("text/plain")
	if err != nil {
		fmt.Println(err)
		tool.ReturnFailure(ctx, "失败")
		return
	}
	filename := header.Filename

	out, err := os.Create("/downloadfile" + filename)
	if err != nil {
		fmt.Println(err)
		tool.ReturnFailure(ctx, "失败")
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		fmt.Println(err)
		tool.ReturnFailure(ctx, "失败")
		return
	}
	tool.ReturnSuccess(ctx, "成功")
}

func SubmitAssignment(ctx *gin.Context) {

}
