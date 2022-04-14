package messageQueue

import (
	"analyze/models"
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

var (
	conn *amqp.Connection
	ch   *amqp.Channel
	msgs <-chan amqp.Delivery
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

	// 声明队列
	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		zap.L().Error("[RabbitMQ]Setting QueueDeclare errors.")
		return err
	}

	// 交换器与队列绑定
	err = ch.QueueBind(
		q.Name,                                // queue name
		"",                                    // routing key
		viper.GetString("rabbitmq.exchanger"), // exchange
		false,
		nil,
	)
	if err != nil {
		zap.L().Error("[RabbitMQ]Setting QueueBind errors.")
		return err
	}

	// 设置消费
	msgs, err = ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		zap.L().Error("[RabbitMQ]Setting Consume errors.")
		return err
	}

	return nil
}

// 关闭
func Close() {
	conn.Close()
	ch.Close()
}

// 获取json内容
func Get() (signal models.Signal, err error) {
	msg := <-msgs
	err = json.Unmarshal(msg.Body, &signal)
	if err != nil {
		zap.L().Error("[RabbitMQ]Unmarshal errors.")
	}
	return
}

// 推送json内容
func Publish(body []byte) error {
	return ch.Publish(
		"analyze_to_visualize", // exchange
		"",                     // routing key
		false,                  // mandatory
		false,                  // immediate
		amqp.Publishing{
			ContentType: viper.GetString("rabbitmq.content-type"),
			Body:        body,
		})
}
