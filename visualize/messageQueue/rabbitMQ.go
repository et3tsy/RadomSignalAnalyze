package messageQueue

import (
	"encoding/json"
	"fmt"
	"visual/models"

	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

var (
	conn      *amqp.Connection
	ch        *amqp.Channel
	msgSignal <-chan amqp.Delivery
	msgResult <-chan amqp.Delivery
)

// 初始化
func Init() (err error) {
	// 设置监听地址
	conn, err = amqp.Dial(fmt.Sprintf("amqp://%v:%v@%v:%v",
		viper.GetString("rabbitmq.user"),
		viper.GetString("rabbitmq.pwd"),
		viper.GetString("rabbitmq.ipv4_addr"),
		viper.GetString("rabbitmq.port"),
	))
	if err != nil {
		zap.L().Error("[RabbitMQ]Setting Dial errors.")
		return err
	}

	// 设置抽象套接字
	ch, err = conn.Channel()
	if err != nil {
		zap.L().Error("[RabbitMQ]Setting Channel errors.")
		return err
	}

	// 声明交换器
	err = ch.ExchangeDeclare(
		viper.GetString("rabbitmq.exchanger"), // name
		"fanout",                              // type
		true,                                  // durable
		false,                                 // auto-deleted
		false,                                 // internal
		false,                                 // no-wait
		nil,                                   // arguments
	)
	if err != nil {
		zap.L().Error("[RabbitMQ]Setting ExchangeDeclare errors.")
		return err
	}

	// 声明来自信号产生端的队列
	queueFromCreate, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		zap.L().Error("[RabbitMQ]Setting QueueDeclare errors.")
		return err
	}

	// 声明来自信号分析端的队列
	queueFromAnalyze, err := ch.QueueDeclare(
		viper.GetString("rabbitmq.result_queue"), // name
		false,                                    // durable
		false,                                    // delete when unused
		false,                                    // exclusive
		false,                                    // no-wait
		nil,                                      // arguments
	)
	if err != nil {
		zap.L().Error("[RabbitMQ]Setting QueueDeclare errors.")
		return err
	}

	// 交换器与队列绑定
	err = ch.QueueBind(
		queueFromCreate.Name,                  // queue name
		"",                                    // routing key
		viper.GetString("rabbitmq.exchanger"), // exchange
		false,
		nil,
	)
	if err != nil {
		zap.L().Error("[RabbitMQ]Setting QueueBind errors.")
		return err
	}

	// 设置消费,接收来自信号产生端的消息
	msgSignal, err = ch.Consume(
		queueFromCreate.Name, // queue
		"",                   // consumer
		true,                 // auto-ack
		false,                // exclusive
		false,                // no-local
		false,                // no-wait
		nil,                  // args
	)
	if err != nil {
		zap.L().Error("[RabbitMQ]Setting Consume from Create errors.")
		return err
	}

	// 设置消费,接收来自信号分析端的消息
	msgResult, err = ch.Consume(
		queueFromAnalyze.Name, // queue
		"",                    // consumer
		true,                  // auto-ack
		false,                 // exclusive
		false,                 // no-local
		false,                 // no-wait
		nil,                   // args
	)
	if err != nil {
		zap.L().Error("[RabbitMQ]Setting Consume from Analyze errors.")
		return err
	}

	return nil
}

// 关闭
func Close() {
	conn.Close()
	ch.Close()
}

// 获取信号
func GetSignal() (signal models.Signal, err error) {
	msg := <-msgSignal
	err = json.Unmarshal(msg.Body, &signal)
	if err != nil {
		zap.L().Sugar().Errorf("[RabbitMQ]Unmarshal errors(%v).", err)
	}
	return
}

// 获取结果
func GetResult() (result models.Result, err error) {
	msg := <-msgResult
	err = json.Unmarshal(msg.Body, &result)
	if err != nil {
		zap.L().Sugar().Errorf("[RabbitMQ]Unmarshal errors(%v).", err)
	}
	return
}
