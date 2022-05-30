package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Ping
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Success 200 {string} json "{"message":"hello world"}"
// @Router / [get]
func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}
