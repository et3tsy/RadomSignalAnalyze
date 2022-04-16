package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 传递信号
func GetSignalsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "200",
	})
}

// 获取结果集
func GetResultHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "200",
	})
}
