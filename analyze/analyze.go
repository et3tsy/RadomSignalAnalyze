package main

import (
	"analyze/calculate"
	"analyze/logger"
	mq "analyze/messageQueue"
	"analyze/models"
	"analyze/settings"
	"encoding/json"
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

	// 初始化结果分析器
	calculate.Init()
}

// 关闭
func close() {
	mq.Close()
}

func main() {

	setup()
	defer close()

	for {
		signal, err := mq.Get()
		if err != nil {
			zap.L().Error("[Analyze]Cannot Fetch signals.")
			break
		}
		calculate.Push(signal.Value)
		result := models.Result{
			Average:    calculate.GetAverage(),
			Variance:   calculate.GetVariance(),
			CreateTime: signal.CreateTime,
		}
		b, err := json.Marshal(result)
		if err != nil {
			zap.L().Error("[Analyze]Cannot Fetch signals.")
			break
		}
		mq.Publish(b)
	}
}
