package router

import "github.com/gin-gonic/gin"

// InitRouter 初始化路由信息
func InitRouter(engine *gin.Engine) {
	api := engine.Group("/api")
	api.GET("/username", UsernameRouter)
	api.POST("/json", JsonRouter)
}
