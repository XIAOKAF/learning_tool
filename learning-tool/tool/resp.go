package tool

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ReturnFailure(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": 200,
		"info":    data,
	})
}

func ReturnSuccess(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": 200,
		"info":    data,
	})
}
