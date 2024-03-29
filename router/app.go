package router

import (
	_ "github.com/codecodify/go-question/docs"
	"github.com/codecodify/go-question/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()

	// 配置路由规则
	r.GET("/", service.Ping)
	// 问题列表
	r.GET("/problems", service.GetProblemList)
	// 问题详情
	r.GET("/problem/detail", service.FindProblemByIdentity)

	// 用户详情
	r.GET("/user/detail", service.FindUserByIdentity)
	// 用户登陆
	r.POST("/user/login", service.Login)
	// 发送邮箱
	r.GET("/user/sendmail", service.SendMail)
	// 用户注册
	r.POST("/user/register", service.Register)
	// 排行榜
	r.GET("/user/rank", service.GetRankList)

	// 提交列表
	r.GET("/submits", service.GetSubmitList)

	// 配置swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
