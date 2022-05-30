package service

import (
	"github.com/codecodify/go-question/define"
	"github.com/codecodify/go-question/helper"
	"github.com/codecodify/go-question/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetProblemList
// @Summary 问题列表
// @Tags 问题
// @Param page query int false "分页，默认1"
// @Param size query int false "分页大小，默认15"
// @Param keyword query string false "关键字"
// @Success 200 {object} object {"status":"success","data":{"count": 1, "list": []}}
// @Failure 400 {object} object {"status":"error","error":"错误信息"}
// @Router /problems [get]    //路由信息，一定要写上
func GetProblemList(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		helper.HandleError(ctx, err, http.StatusBadRequest)
	}
	size, err := strconv.Atoi(ctx.DefaultQuery("size", define.DefaultSize))
	if err != nil {
		helper.HandleError(ctx, err, http.StatusBadRequest)
	}
	keyword := ctx.Query("keyword")

	var count int64
	var problems []*models.Problem
	db := models.GetProblemList(keyword)
	db.Count(&count).Omit("content").Offset((page - 1) * size).Limit(size).Find(&problems)
	helper.HandleSuccess(ctx, gin.H{
		"count": count,
		"list":  problems,
	})
}
