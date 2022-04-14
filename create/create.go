package main

import (
	"create/logger"
	mq "create/messageQueue"
	"create/models"
	"create/random"
	"create/settings"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// 产生错误中断程序
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

const (
	duration        time.Duration = time.Millisecond * 1000 // 产生信号的间隔时间
	randomSignalMin               = 100                     // 信号的最小值
	randomSignalMax               = 10000000                // 信号的最大值
)

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
		value := random.GetGaussRandomNum(randomSignalMin, randomSignalMax) // 产生随机信号

		// 序列化, 以json格式传递
		b, err := json.Marshal(models.Signal{
			Value:      value,
			CreateTime: time.Now(),
		})
		if err != nil {
			zap.L().Error(fmt.Sprintf("%v", err))
			break
		}

		zap.L().Info(fmt.Sprintf("Succeed to create value -- %v", value))

		// 推送
		mq.Publish(b)

		// 睡眠
		time.Sleep(duration)
	}

}
