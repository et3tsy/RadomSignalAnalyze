package routes

import (
	"net/http"
	"visual/controller"
	logger "visual/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 初始化路由
func Setup(mod string) *gin.Engine {
	if mod == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 获取信号
	r.GET("/data", controller.GetSignalsHandler)

	// 获取结果集
	r.GET("/result", controller.GetResultHandler)

	r.NoRoute(func(c *gin.Context) {
		zap.L().Warn("No route")
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return r
}
