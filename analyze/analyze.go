package main

import (
	"analyze/calculate"
	"analyze/logger"
	mq "analyze/messageQueue"
	"analyze/models"
	"analyze/settings"
	"encoding/json"
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
	// 初始化以及关闭通道
	setup()
	defer close()

	for {
		// 获取信号
		signal, err := mq.Get()
		if err != nil {
			zap.L().Error("[Analyze]Cannot Fetch signals.")
			break
		}

		// 加入新的信号,进行分析
		calculate.Push(signal.Value)
		result := models.Result{
			Average:    calculate.GetAverage(),
			Variance:   calculate.GetVariance(),
			CreateTime: signal.CreateTime,
		}

		// 序列化
		b, err := json.Marshal(result)
		if err != nil {
			zap.L().Error("[Analyze]Cannot marshal the result.")
			continue
		}

		// 将结果打包发送,传递给可视化微服务
		mq.Publish(b)
		zap.L().Info(fmt.Sprintf("[Analyze]Send result: %v", result))
	}
}
