package server

import (
	"context"
	"github.com/pkg/errors"

	"sunflower-blog-svc/app/blog/internal/conf"
	"sunflower-blog-svc/pkg/mq"

	"github.com/go-kratos/kratos/v2/log"
)

type mqConf struct {
	EndPoint []string
	Topic    string
}

type MQServer struct {
	mqConf *conf.MQ

	logger     *log.Helper
	MQSender   map[string]mq.Sender
	MQReceiver map[string]mq.Receiver
}

func (m *MQServer) Start(ctx context.Context) error {
	// 处理生产者配置
	for name, producer := range m.mqConf.Kafka.Producers.Items {
		sender, err := mq.NewKafkaSender(
			m.mqConf.Kafka.Producers.Endpoint,
			producer.Topic,
		)
		if err != nil {
			return errors.Wrap(err, "failed to create kafka sender: ")
		}
		m.MQSender[name] = sender
	}

	// 处理消费者配置
	for name, consumer := range m.mqConf.Kafka.Consumers.Items {
		// 由于 NewKafkaReceiver 只接受单个 topic，我们为每个 topic 创建一个接收器
		for _, topic := range consumer.Topics {
			receiver, err := mq.NewKafkaReceiver(
				m.mqConf.Kafka.Consumers.Endpoint,
				topic,
			)
			if err != nil {
				return errors.Wrap(err, "failed to create kafka receiver: ")
			}
			// 使用 name_topic 作为键
			receiverKey := name + "_" + topic
			m.MQReceiver[receiverKey] = receiver
		}
	}

	m.logger.Info("MQ server started")
	return nil
}

func (m *MQServer) Stop(ctx context.Context) error {
	// 关闭所有的 sender
	for name, sender := range m.MQSender {
		if err := sender.Close(); err != nil {
			log.Error("failed to close kafka sender: ", name, ", error: ", err)
			return errors.Wrap(err, "failed to close kafka sender: ")
		}
	}

	// 关闭所有的 receiver
	for name, receiver := range m.MQReceiver {
		if err := receiver.Close(); err != nil {
			log.Error("failed to close kafka receiver: ", name, ", error: ", err)
			return errors.Wrap(err, "failed to close kafka receiver: ")
		}
	}

	m.logger.Info("MQ server stopped")
	return nil
}

func NewMQServer(
	bc *conf.Bootstrap,
	logger log.Logger,
) *MQServer {
	mqServer := &MQServer{
		mqConf:     bc.Mq,
		logger:     log.NewHelper(log.With(logger, "server", "MQ")),
		MQSender:   make(map[string]mq.Sender),
		MQReceiver: make(map[string]mq.Receiver),
	}

	return mqServer
}
