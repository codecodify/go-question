package helper

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleError(ctx *gin.Context, err error, code int) {
	ctx.JSON(code, gin.H{
		"status": "error",
		"error":  err.Error(),
	})
}

func HandleSuccess(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   data,
	})
}
