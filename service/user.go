package service

import (
	"errors"
	"github.com/codecodify/go-question/helper"
	"github.com/codecodify/go-question/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// FindUserByIdentity
// @Summary 用户详情
// @Tags 用户
// @Param identity query string true "用户标识"
// @Success 200 {object} object {"status":"success","data":{}}
// @Failure 400 {object} object {"status":"error","error":"错误信息"}
// @Router /user/detail [get]    //路由信息，一定要写上
func FindUserByIdentity(ctx *gin.Context) {
	identity := ctx.Query("identity")
	if len(identity) == 0 {
		helper.HandleError(ctx, errors.New("identity is empty"), http.StatusBadRequest)
		return
	}
	user, err := models.FindUserByIdentity(identity)
	if err != nil {
		helper.HandleError(ctx, errors.New("user not found"), http.StatusNotFound)
		return
	}
	helper.HandleSuccess(ctx, user)
}
