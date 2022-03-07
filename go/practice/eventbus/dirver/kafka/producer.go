/*
 * @Author: 馬濤
 * @Date: 2021-01-23 10:15:52
 * @LastEditTime: 2021-09-18 02:15:22
 * @LastEditors: 馬濤
 * @Description:
 * @FilePath: \cf\eventbus\dirver\kafka\producer.go
 * @MT is your father.
 */

package kafka

import (
	"context"

	"git.dustess.com/mk-base/eventbus/dirver/options"
	"git.dustess.com/mk-base/eventbus/dirver/utils"
	kafkaProducer "git.dustess.com/mk-base/kafka-driver/producer"
	"git.dustess.com/mk-base/log"
)

// Producer Kafka
type Producer struct {
	opts     *options.Server
	producer *kafkaProducer.SyncProducer
}

// NewProducer NewProducer
func NewProducer(addrs []string, topic string, opts ...options.ProducerOption) *Producer {
	tracep, err := kafkaProducer.NewSyncProducer(addrs, topic)
	if err != nil {
		panic(err)
	}
	p := &Producer{
		opts: &options.Server{
			Addrs: addrs,
			Topic: topic,
		},
		producer: tracep,
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
	messages := make([]*kafkaProducer.Message, 0, len(msg))
	for i := range msg[:] {
		// 设置header
		msgHeader := options.DefaulHeader()
		msgHeader = options.HeaderWithID(msgHeader)
		msgHeader = options.HeaderMerge(msgHeader, headers)

		messag := utils.BytesToString(msg[i])
		message := client.getMessage(messag, msgHeader, op)
		messages = append(messages, message)
	}
	if len(messages) == 1 {
		return client.sendOne(messages[0], op)
	}
	return client.sendMany(messages, op)
}

// Close 关闭
func (client *Producer) Close() error {
	return nil
}

// send 发送
func (client *Producer) sendOne(msg *kafkaProducer.Message, opt *options.ProducerOptions) error {
	log.Infof("eventbus kafka sendOne opt[%v] msg[%+v]", opt, msg)
	_, _, err := client.producer.ProduceOneMessage(msg)
	if err != nil {
		log.Errorf("eventbus kafka sendOne opt[%v] msg[%+v] err[%v]", opt, msg, err)
	}
	return err
}

// sendMany 发送
func (client *Producer) sendMany(msg []*kafkaProducer.Message, opt *options.ProducerOptions) error {
	log.Infof("eventbus kafka sendMany opt[%v] msg[%+v]", opt, msg)
	if err := client.producer.ProduceManyMessage(msg); err != nil {
		log.Errorf("eventbus kafka sendMany opt[%v] msg[%+v] err[%v]", opt, msg, err)
		return err
	}
	return nil
}

// getMessage 获取message
func (client *Producer) getMessage(msg string, headers map[string]string, opt *options.ProducerOptions) *kafkaProducer.Message {
	if opt.Order {
		key, ok := options.HeaderOrderField(headers)
		if ok {
			return kafkaProducer.NewMessageWithKey(headers, msg, utils.StringToBytes(key))
		}
	}
	return kafkaProducer.NewMessage(headers, msg)
}
