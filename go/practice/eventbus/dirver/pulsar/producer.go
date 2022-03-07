package pulsar

import (
	"context"
	"time"

	"git.dustess.com/mk-base/eventbus/dirver/options"
	"git.dustess.com/mk-base/log"
	pulsarProducer "git.dustess.com/shared-golib/pulsar-driver/producer"
)

// 最大延迟时间
const maxDeliverDuration = 863000 * time.Second

// Producer Kafka
type Producer struct {
	opts     *options.Server
	producer *pulsarProducer.Producer
}

// NewProducer NewProducer
func NewProducer(addrs []string, topic string, opts ...options.ProducerOption) *Producer {
	op := options.NewProducerOptions()
	for _, opt := range opts {
		opt(op)
	}
	config := pulsarProducer.Config{
		Addrs: addrs,
		Topic: topic,
		Name:  "",
		Token: op.Token,
	}
	pulasr, err := pulsarProducer.NewProducer(config)
	if err != nil {
		panic(err)
	}

	p := &Producer{
		opts: &options.Server{
			Addrs: addrs,
			Topic: topic,
		},
		producer: pulasr,
	}
	return p
}

// Send 发送
func (client *Producer) Send(ctx context.Context, msg [][]byte, header map[string]string, opts ...options.ProducerOption) error {
	if len(msg) == 0 {
		return nil
	}
	headers := options.DefaulHeader()
	headers = options.HeaderWithContext(ctx, headers)
	headers = options.HeaderMerge(headers, header)
	op := options.NewProducerOptions()
	for _, opt := range opts {
		opt(op)
	}
	messages := make([]*pulsarProducer.Message, 0, len(msg))
	for i := range msg[:] {
		// 设置header
		msgHeader := options.DefaulHeader()
		msgHeader = options.HeaderWithID(msgHeader)
		msgHeader = options.HeaderMerge(msgHeader, headers)

		message := client.getMessage(msg[i], msgHeader, op)
		messages = append(messages, message)
	}
	if len(messages) == 1 {
		return client.sendOne(ctx, messages[0], op)
	}
	return client.sendMany(ctx, messages, op)
}

// Close 关闭
func (client *Producer) Close() error {
	return nil
}

// send 发送
func (client *Producer) sendOne(ctx context.Context, msg *pulsarProducer.Message, opt *options.ProducerOptions) error {
	log.Infof("eventbus pulsar sendOne ctx[%v] opt[%v] msg[%+v]", ctx, opt, msg)
	_, err := client.producer.SendOneWithMessage(ctx, msg)
	if err != nil {
		log.Errorf("eventbus pulsar sendOne ctx[%v] opt[%v] msg[%+v]", ctx, opt, msg)
	}
	return err
}

// sendMany 发送
func (client *Producer) sendMany(ctx context.Context, msg []*pulsarProducer.Message, opt *options.ProducerOptions) error {
	for i := range msg[:] {
		log.Infof("eventbus pulsar sendOne ctx[%v] opt[%v] msg[%+v]", ctx, opt, msg[i])
		if _, err := client.producer.SendOneWithMessage(ctx, msg[i]); err != nil {
			log.Errorf("eventbus pulsar sendOne ctx[%v] opt[%v] msg[%+v]", ctx, opt, msg[i])
		}
	}

	return nil
}

// getMessage 获取message
func (client *Producer) getMessage(msg []byte, headers map[string]string, opt *options.ProducerOptions) *pulsarProducer.Message {
	message := &pulsarProducer.Message{
		Properties: headers,
		Payload:    msg,
	}
	if opt.Order {
		key, ok := options.HeaderOrderField(headers)
		if ok {
			message.Key = key
			return message
		}
	}
	if opt.DelayTime > 0 {
		// 大于最大时间，取最大时间为延迟时间并且把剩余的放到header里
		if opt.DelayTime > maxDeliverDuration {
			remainDuration := opt.DelayTime - maxDeliverDuration
			message.DeliverAfter = maxDeliverDuration
			message.Properties = options.HeaderWithPulsarDeliverRemainDuration(message.Properties, remainDuration)
		} else {
			message.DeliverAfter = opt.DelayTime
		}
	}
	return message
}
