package mq

import (
	"errors"
	"fmt"
	"github.com/njylll/thirdparty_auxiliary_tool_go/setting"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type RabbitmqClient struct {
	ConnStr string
	Topic   string
	Conn    *amqp.Connection
	Ch      *amqp.Channel
}

var RbClient *RabbitmqClient

func Setup() {
	var err error
	logrus.Info("开始初始化rabbitmq")

	RbClient, err = newRabbitmqClient(setting.SendCmdSetting.ConnStr, setting.SendCmdSetting.Topic)
	if err != nil {
		logrus.Fatalf("初始化rabbitmq失败, err: %v", err)
	}

	logrus.Info("初始化rabbitmq成功")
}

func newRabbitmqClient(connStr, topic string) (*RabbitmqClient, error) {
	client := &RabbitmqClient{
		ConnStr: connStr,
		Topic:   topic,
	}
	//连接mq
	var err error
	client.Conn, err = amqp.Dial("amqp://" + setting.SendCmdSetting.ConnStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ, err: %v", err)
	}

	client.Ch, err = client.Conn.Channel()
	//声明交换机
	err = client.Ch.ExchangeDeclare(
		setting.SendCmdSetting.Topic, // name
		"topic",                      // type
		true,                         // durable
		false,                        // auto-deleted
		false,                        // internal
		false,                        // no-wait
		nil,                          // arguments
	)
	return client, nil
}

func (r *RabbitmqClient) SendMsgToRabbitMq(topic string, data []byte) error {
	err := r.Ch.Publish(
		topic, // exchange
		"",    // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		})
	if err != nil {
		return fmt.Errorf("发送消息至mq失败, err: %v", err)
	}
	return nil
}

// todo: 何时关闭
func (r *RabbitmqClient) Close() error {
	return errors.Join(
		r.Conn.Close(),
		r.Ch.Close(),
	)
}
