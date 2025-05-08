package server

import (
	"context"

	"github.com/pkg/errors"

	"sunflower-blog-svc/app/blog/internal/conf"
	"sunflower-blog-svc/app/blog/internal/service/mq"
	pkgmq "sunflower-blog-svc/pkg/mq"

	"github.com/go-kratos/kratos/v2/log"
)

type MQServer struct {
	mqConf *conf.MQ

	logger *log.Helper

	MQReceiver       map[string]pkgmq.Receiver
	consumerHandlers *mq.HandlerRegistry

	senderFactory *mq.SenderFactory
}

func (m *MQServer) Start(ctx context.Context) error {
	// todo: 启动所有的 receiver
	for name, receiver := range m.MQReceiver {
		handler, ok := m.consumerHandlers.Get(name)
		if !ok {
			return errors.New("unknown mq consumer: " + name)
		}
		receiver.Receive(ctx, handler)
	}

	m.logger.Info("MQ server started")
	return nil
}

func (m *MQServer) Stop(ctx context.Context) error {
	// 关闭所有的 sender
	err := m.senderFactory.Close()
	if err != nil {
		log.Error("failed to close kafka sender: ", err)
		return errors.Wrap(err, "failed to close kafka sender: ")
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
	handlers *mq.HandlerRegistry,
	senderFactory *mq.SenderFactory,
) *MQServer {
	mqServer := &MQServer{
		mqConf:           bc.Mq,
		logger:           log.NewHelper(log.With(logger, "server", "MQ")),
		MQReceiver:       make(map[string]pkgmq.Receiver),
		consumerHandlers: handlers,
		senderFactory:    senderFactory,
	}

	kafkaConf := bc.Mq.GetKafka()

	senders := make(map[string]pkgmq.Sender)
	// 处理生产者配置
	for name, producer := range kafkaConf.Producers.Items {
		sender, err := pkgmq.NewKafkaSender(
			kafkaConf.Producers.Endpoint,
			producer.Topic,
		)
		if err != nil {
			log.Fatalf("failed to create kafka sender: %v", err)
		}
		senders[name] = sender
	}
	senderFactory.AddSenders(senders)

	// 处理消费者配置
	for name, consumer := range mqServer.mqConf.Kafka.Consumers.Items {
		// 由于 NewKafkaReceiver 只接受单个 topic，我们为每个 topic 创建一个接收器
		for _, topic := range consumer.Topics {
			receiver, err := pkgmq.NewKafkaReceiver(
				mqServer.mqConf.Kafka.Consumers.Endpoint,
				topic,
			)
			if err != nil {
				log.Fatalf("failed to create kafka receiver: %v", err)
			}
			// 使用 name_topic 作为键
			receiverKey := name + "_" + topic
			mqServer.MQReceiver[receiverKey] = receiver
		}
	}

	return mqServer
}
