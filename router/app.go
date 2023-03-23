package router

import (
	_ "Gin_OJ/docs"
	"Gin_OJ/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	//swagger配置
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//路由规则
	r.GET("/ping", service.Ping)

	//问题
	r.GET("/problem_list", service.GetProblemList)
	r.GET("/problem_detail", service.GetProblemDetail)

	//用户
	r.POST("/login", service.Login)
	r.GET("/user-detail", service.GetUserDetail)
	r.POST("/send-code", service.SendCode)
	r.POST("/register", service.Register)
	// 排行榜
	r.GET("/rank-list", service.GetRankList)
	//提交记录
	//r.GET("/submit-list", service.GetSubmitList)

	return r
}
