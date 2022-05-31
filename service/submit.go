package service

import (
	"github.com/codecodify/go-question/define"
	"github.com/codecodify/go-question/helper"
	"github.com/codecodify/go-question/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetSubmitList
// @Summary 提交列表
// @Tags 提交
// @Param problem_identity query string false "问题标识"
// @Param user_identity query string false "用户标识"
// @Param status query int false "状态"
// @Param page query int false "分页，默认1"
// @Param size query int false "分页大小，默认15"
// @Success 200 {object} object {"status":"success","data":{"count": 1, "list": []}}
// @Failure 400 {object} object {"status":"error","error":"错误信息"}
// @Router /submits [get]    //路由信息，一定要写上
func GetSubmitList(ctx *gin.Context) {
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
	status, err := strconv.Atoi(ctx.DefaultQuery("status", define.DefaultSubmitStatus))
	if err != nil {
		helper.HandleError(ctx, err, http.StatusBadRequest)
		return
	}
	problemIdentity := ctx.Query("problem_identity")
	userIdentity := ctx.Query("user_identity")

	db := models.GetSubmitList(problemIdentity, userIdentity, status)

	var count int64
	var submits []*models.Submit
	db.Count(&count).Offset((page - 1) * size).Limit(size).Find(&submits)
	helper.HandleSuccess(ctx, gin.H{
		"count": count,
		"list":  submits,
	})

}
