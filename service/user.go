package service

import (
	"errors"
	"github.com/codecodify/go-question/define"
	"github.com/codecodify/go-question/helper"
	"github.com/codecodify/go-question/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
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

// Register 用户注册
// @Summary 用户注册
// @Tags 用户
// @Param name formData string true "用户名"
// @Param password formData string true "密码"
// @Param email formData string true "邮箱"
// @Param phone formData string true "手机"
// @Param code formData string true "验证码"
// @Success 200 {object} object {"status":"success","data":{}}
// @Failure 400 {object} object {"status":"error","error":"错误信息"}
// @Router /user/register [post]    //路由信息，一定要写上
func Register(ctx *gin.Context) {
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")
	email := ctx.PostForm("email")
	code := ctx.PostForm("code")
	phone := ctx.PostForm("phone")
	if len(name) == 0 || len(password) == 0 || len(email) == 0 || len(code) == 0 || len(phone) == 0 {
		helper.HandleError(ctx, errors.New("name or password or email or code or phone is empty"), http.StatusBadRequest)
		return
	}

	// 判断名称是否占用
	if models.GetUserByName(name) {
		helper.HandleError(ctx, errors.New("用户名已被占用"), http.StatusBadRequest)
		return
	}

	// 判断邮箱是否占用
	if models.GetUserByEmail(email) {
		helper.HandleError(ctx, errors.New("邮箱已被占用"), http.StatusBadRequest)
		return
	}

	// 校验验证码
	redisCode, err := helper.GetRedisString(email)
	if err != nil {
		helper.HandleError(ctx, errors.New("验证码已过期"), http.StatusBadRequest)
		return
	}
	if redisCode != code {
		helper.HandleError(ctx, errors.New("验证码错误"), http.StatusBadRequest)
		return
	}

	// 创建用户
	user := models.User{
		Name:     name,
		Password: helper.Md5(password),
		Email:    email,
		Identity: helper.GetUUID(),
		Phone:    phone,
	}
	err = models.CreateUser(&user)
	if err != nil {
		helper.HandleError(ctx, err, http.StatusInternalServerError)
		return
	}
	helper.HandleSuccess(ctx, gin.H{})
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
	code := helper.GetRandomCode()
	err := helper.SendMail(email, code)
	if err != nil {
		helper.HandleError(ctx, err, http.StatusInternalServerError)
		return
	}
	// 将验证码存储到redis
	err = helper.SetRedisString(email, code)
	if err != nil {
		helper.HandleError(ctx, err, http.StatusInternalServerError)
		return
	}
	helper.HandleSuccess(ctx, gin.H{})
}

// GetRankList 用户排行榜
// @Summary 用户排行榜
// @Tags 用户
// @Param page query int false "分页，默认1"
// @Param size query int false "分页大小，默认15"
// @Success 200 {object} object {"status":"success","data":{}}
// @Failure 400 {object} object {"status":"error","error":"错误信息"}
// @Router /user/rank [get]    //路由信息，一定要写上
func GetRankList(ctx *gin.Context) {
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

	// 获取排行榜
	db := models.GetRankList()
	var rankList []models.User
	var count int64
	db.Count(&count).Omit("password").Offset((page - 1) * size).Limit(size).Find(&rankList)
	helper.HandleSuccess(ctx, gin.H{
		"count": count,
		"list":  rankList,
	})

}
