package messageQueue

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

var (
	conn *amqp.Connection
	ch   *amqp.Channel
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
		return err
	}

	// 设置抽象套接字
	ch, err = conn.Channel()
	if err != nil {
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
		return err
	}

	return nil
}

// 关闭
func Close() {
	conn.Close()
	ch.Close()
}

// 推送json内容
func Publish(body []byte) error {
	return ch.Publish(
		"signal", // exchange
		"",       // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: viper.GetString("rabbitmq.content-type"),
			Body:        body,
		})
}
