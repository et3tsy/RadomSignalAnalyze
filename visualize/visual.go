package main

import (
	"log"
	"visual/logger"
	mq "visual/messagequeue"
	"visual/settings"

	"github.com/spf13/viper"
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

}

// 关闭
func close() {
	mq.Close()
}

func main() {
	// 初始化以及关闭通道
	setup()
	defer close()

	for {
		// // 获取信号
		// signal, err := mq.Get()
		// if err != nil {
		// 	zap.L().Error("[Analyze]Cannot Fetch signals.")
		// 	break
		// }

	}
}
