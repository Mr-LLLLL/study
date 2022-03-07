/*
 * @Author: 馬濤
 * @Date: 2021-01-23 10:15:52
 * @LastEditTime: 2021-02-03 11:49:04
 * @LastEditors: 馬濤
 * @Description:
 * @FilePath: /cf/eventbus/dirver/rocketmq/producer.go
 * @MT is your father.
 */

package rocketmq

import (
	"context"

	rocketmqOptions "git.dustess.com/easyio/go-rocketmq-driver-grpc/rocketmq/options"
	"git.dustess.com/easyio/go-rocketmq-driver-grpc/rocketmq/producer"
	"git.dustess.com/mk-base/eventbus/dirver/options"
	"git.dustess.com/mk-base/eventbus/dirver/utils"
)

// Producer Kafka
type Producer struct {
	opts     *options.Server
	producer *producer.Producer
}

// NewProducer NewProducer
func NewProducer(addrs []string, topic string, opts ...options.ProducerOption) *Producer {
	if len(addrs) == 0 {
		panic("rocket address error")
	}
	client, err := producer.Connect(context.TODO(), addrs[0])
	if err != nil {
		panic(err)
	}
	proc := client.Cluster(rocketmqOptions.Default).Producer()

	p := &Producer{
		opts: &options.Server{
			Addrs: addrs,
			Topic: topic,
		},
		producer: proc,
	}
	return p
}

// Send 发送
func (client *Producer) Send(ctx context.Context, msg [][]byte, header map[string]string, opts ...options.ProducerOption) error {
	if len(msg) == 0 {
		return nil
	}
	msgs := make([]*producer.Message, 0, len(msg))
	op := options.NewProducerOptions()
	for _, opt := range opts {
		opt(op)
	}
	for i := range msg[:] {
		m := &producer.Message{
			Topic: client.opts.Topic,
			Body:  utils.BytesToString(msg[i]),
			Tag:   op.Tag,
		}
		msgs = append(msgs, m)
	}
	if len(msgs) == 1 {
		return client.sendOne(msgs[0], op)
	}
	return client.sendMany(msgs, op)
}

// Close 关闭
func (client *Producer) Close() error {
	return nil
}

// send 发送
func (client *Producer) sendOne(msg *producer.Message, opt *options.ProducerOptions) error {
	if opt.IsAsync {
		return client.producer.SendAsyncOne(context.TODO(), msg.Topic, msg.Body, msg.Tag)
	}
	return client.producer.SendSyncOne(context.TODO(), msg.Topic, msg.Body, msg.Tag)
}

// sendMany 发送
func (client *Producer) sendMany(msg []*producer.Message, opt *options.ProducerOptions) error {
	if opt.IsAsync {
		return client.producer.SendAsyncBatch(context.TODO(), msg)
	}
	return client.producer.SendSyncBatch(context.TODO(), msg)
}
