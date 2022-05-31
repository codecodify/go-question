package service

import (
	"errors"
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
// @Param category_identity query string false "分类标识"
// @Success 200 {object} object {"status":"success","data":{"count": 1, "list": []}}
// @Failure 400 {object} object {"status":"error","error":"错误信息"}
// @Router /problems [get]    //路由信息，一定要写上
func GetProblemList(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		helper.HandleError(ctx, err, http.StatusBadRequest)
		return
	}
	size, err := strconv.Atoi(ctx.DefaultQuery("size", define.DefaultSize))
	if err != nil {
		helper.HandleError(ctx, err, http.StatusBadRequest)
		return
	}
	keyword := ctx.Query("keyword")
	categoryIdentity := ctx.Query("category_identity")

	var count int64
	var problems []*models.Problem
	db := models.GetProblemList(keyword, categoryIdentity)
	db.Count(&count).Omit("content").Offset((page - 1) * size).Limit(size).Find(&problems)
	helper.HandleSuccess(ctx, gin.H{
		"count": count,
		"list":  problems,
	})
}

// FindProblemByIdentity
// @Summary 问题详情
// @Tags 问题
// @Param identity query string true "问题标识"
// @Success 200 {object} object {"status":"success","data":{"id": 1, "identity": "", "title": "", "content": "", "total_num": 0, "problem_categories": []}}
// @Failure 400 {object} object {"status":"error","error":"错误信息"}
// @Router /problem/detail [get]    //路由信息，一定要写上
func FindProblemByIdentity(ctx *gin.Context) {
	identity := ctx.Query("identity")
	if len(identity) == 0 {
		helper.HandleError(ctx, errors.New("identity is empty"), http.StatusBadRequest)
		return
	}
	problem, err := models.FindProblemByIdentity(identity)
	if err != nil {
		helper.HandleError(ctx, err, http.StatusNotFound)
		return
	}
	helper.HandleSuccess(ctx, problem)
}
