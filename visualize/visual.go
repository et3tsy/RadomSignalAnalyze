package main

import (
	"log"
	"visual/logger"
	"visual/logic"
	mq "visual/messageQueue"
	"visual/routes"
	"visual/settings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	r *gin.Engine
)

// 产生错误中断程序
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// 进行初始化设置
func setup() {
	// 初始化viper
	if err := settings.Init(); err != nil {
		failOnError(err, "init viper failed")
	}

	// 初始化日志记录
	if err := logger.Init(viper.GetString("app.mod")); err != nil {
		failOnError(err, "init logger failed")
	}

	// 初始化RabbitMQ
	if err := mq.Init(); err != nil {
		failOnError(err, "init RabbitMQ failed")
	}

	// 初始化线段树
	logic.Init()

	// 设置路由
	r = routes.Setup(viper.GetString("app.mod"))
}

// 关闭
func close() {
	mq.Close()
}

func main() {

	// 初始化以及关闭通道
	setup()
	defer close()

	// 接收信号消息并动态更新
	go logic.ReceiveSignalAndUpdate()

	// 接收统计消息并动态更新
	go logic.ReceiveResultAndUpdate()

	// 启动路由
	r.Run(":" + viper.GetString("app.port"))
}
