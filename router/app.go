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
	r.GET("/problems", service.GetProblemList)

	// 配置swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
