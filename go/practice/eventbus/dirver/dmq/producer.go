/*
 * @Author: 馬濤
 * @Date: 2021-01-23 10:15:52
 * @LastEditTime: 2021-01-28 11:07:18
 * @LastEditors: 馬濤
 * @Description:
 * @FilePath: /cf/eventbus/dirver/kafka/producer.go
 * @MT is your father.
 */

package dmq

import (
	"context"

	dmqConfig "git.dustess.com/mk-base/dmq/config"
	"git.dustess.com/mk-base/dmq/message"
	"git.dustess.com/mk-base/dmq/producer"
	"git.dustess.com/mk-base/eventbus/dirver/options"
	"git.dustess.com/mk-base/log"
)

// Producer Kafka
type Producer struct {
	opts     *options.Server
	producer *producer.SyncProducer
}

const EventbusHeaderField = "EventbusHeader"

// NewProducer NewProducer
func NewProducer(addrs []string, topic string, opts ...options.ProducerOption) *Producer {
	op := options.NewProducerOptions()
	for _, opt := range opts {
		opt(op)
	}
	config := dmqConfig.NewConfig(topic, int8(op.DMQDriver), uint8(op.Partition))
	dmq := producer.NewSyncProducer(config)
	p := &Producer{
		opts: &options.Server{
			Addrs: addrs,
			Topic: topic,
		},
		producer: dmq,
	}
	return p
}

// Send 发送
func (client *Producer) Send(ctx context.Context, msg [][]byte, header map[string]string, opts ...options.ProducerOption) error {
	if len(msg) == 0 {
		return nil
	}

	op := options.NewProducerOptions()
	for _, opt := range opts {
		opt(op)
	}

	h := make(message.Header)
	h[message.HeaderXDelay] = op.DelayTime

	headers := options.DefaulHeader()
	headers = options.HeaderWithContext(ctx, headers)
	headers = options.HeaderMerge(headers, header)

	messages := make([]*message.ManyMessage, 0, len(msg))
	for i := range msg[:] {
		// 设置header
		msgHeader := options.DefaulHeader()
		msgHeader = options.HeaderWithID(msgHeader)
		msgHeader = options.HeaderMerge(msgHeader, headers)

		h := make(message.Header)
		h[message.HeaderXDelay] = op.DelayTime
		h[EventbusHeaderField] = msgHeader
		m := &message.ManyMessage{
			Value:  msg[i],
			Header: h,
		}
		messages = append(messages, m)
	}
	dmqMsg := message.NewManyMessage(messages)
	if len(dmqMsg) == 1 {
		return client.sendOne(ctx, dmqMsg[0], op)
	}
	return client.sendMany(ctx, dmqMsg, op)
}

// Close 关闭
func (client *Producer) Close() error {
	return nil
}

// send 发送
func (client *Producer) sendOne(ctx context.Context, msg *message.Message, opt *options.ProducerOptions) error {
	_, _, err := client.producer.SendMessage(ctx, msg)
	if err != nil {
		log.Infof("eventbus dmq sendOne msg[%+v] err[%v]", msg,err)
	}
	return err
}

// send 发送
func (client *Producer) sendMany(ctx context.Context, msg []*message.Message, opt *options.ProducerOptions) error {
	err := client.producer.SendMessages(ctx, msg)
	if err != nil {
		log.Infof("eventbus dmq sendMany msg[%+v] err[%v]", msg,err)
	}
	return err
}
