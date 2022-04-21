package controller

import (
	"net/http"
	"strconv"
	"visual/logic"
	"visual/models"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// 传递信号
func GetSignalsHandler(c *gin.Context) {
	// 获取url中参数
	param := c.Query("cols")

	// 进行参数的转换
	cols, err := strconv.Atoi(param)
	if err != nil {
		cols = viper.GetInt("visual.segments")
	}

	// 向前端返回处理的结果集合
	c.JSON(http.StatusOK, gin.H{
		"msg": "200",
		"info": logic.GetSignalStatistics(
			viper.GetInt64("signal.min"),
			viper.GetInt64("signal.max"),
			int64(cols),
		),
	})
}

// 获取结果集
func GetResultHandler(c *gin.Context) {
	res := logic.GetResult()
	c.JSON(http.StatusOK, gin.H{
		"msg": "200",
		"info": models.ParamResult{
			Average:    res.Average,
			Variance:   res.Variance,
			CreateTime: logic.TimeFormatString(res.CreateTime),
		},
	})
}
