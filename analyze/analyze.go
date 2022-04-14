package main

import (
	"analyze/logger"
	mq "analyze/messageQueue"
	"analyze/settings"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// 产生错误中断程序
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
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

	for {
		signal, err := mq.Get()
		if err != nil {
			zap.L().Error("[Analyze]Cannot Fetch signals.")
			break
		}
		fmt.Printf("%v\n", signal.Value)
	}
}
