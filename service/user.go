package service

import (
	"errors"
	"github.com/codecodify/go-question/helper"
	"github.com/codecodify/go-question/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

// Login
// @Summary 用户登陆
// @Tags 用户
// @Param name formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {object} object {"status":"success","data":{}}
// @Failure 400 {object} object {"status":"error","error":"错误信息"}
// @Router /user/login [post]    //路由信息，一定要写上
func Login(ctx *gin.Context) {
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")
	if len(name) == 0 || len(password) == 0 {
		helper.HandleError(ctx, errors.New("name or password is empty"), http.StatusBadRequest)
		return
	}
	user, err := models.Login(name, helper.Md5(password))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			helper.HandleError(ctx, errors.New("用户名或密码错误"), http.StatusBadRequest)
			return
		}
		helper.HandleError(ctx, err, http.StatusInternalServerError)
	}

	// 生成密钥
	token, err := helper.GetUserToken(user.Identity, user.Name)
	if err != nil {
		helper.HandleError(ctx, err, http.StatusBadRequest)
		return
	}
	helper.HandleSuccess(ctx, gin.H{
		"token": token,
	})
}

// SendMail 发送邮件
// @Summary 发送邮件
// @Tags 用户
// @Param email query string true "接收邮件地址"
// @Success 200 {object} object {"status":"success","data":{}}
// @Failure 400 {object} object {"status":"error","error":"错误信息"}
// @Router /user/sendmail [get]    //路由信息，一定要写上
func SendMail(ctx *gin.Context) {
	email := ctx.Query("email")
	if len(email) == 0 {
		helper.HandleError(ctx, errors.New("email is empty"), http.StatusBadRequest)
		return
	}

	// 发送邮件
	code := "123456"
	err := helper.SendMail(email, code)
	if err != nil {
		helper.HandleError(ctx, err, http.StatusInternalServerError)
		return
	}
	helper.HandleSuccess(ctx, gin.H{})
}