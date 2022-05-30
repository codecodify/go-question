package service

import (
	"github.com/codecodify/go-question/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProblemList(ctx *gin.Context) {
	models.GetProblemList()
	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
