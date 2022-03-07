/*
 * @Author: 馬濤
 * @Date: 2021-01-25 15:30:00
 * @LastEditTime: 2021-09-18 02:45:13
 * @LastEditors: 馬濤
 * @Description:
 * @FilePath: \cf\eventbus\dirver\rocketmq\consumer.go
 * @MT is your father.
 */

package rocketmq

import (
	"context"

	"git.dustess.com/easyio/go-rocketmq-driver-grpc/rocketmq/consumer"
	rocketmqOptions "git.dustess.com/easyio/go-rocketmq-driver-grpc/rocketmq/options"
	"git.dustess.com/mk-base/eventbus/dirver/options"
	"git.dustess.com/mk-base/eventbus/dirver/utils"
)

// Consumer Kafka
type Consumer struct {
	opts        *options.Server
	consumer    *consumer.Consumer
	consumeOpts *options.ConsumerOptions
	msg         chan options.Message
}

// NewConsumer NewConsumer
func NewConsumer(addrs []string, topic string, opts ...options.ConsumerOption) *Consumer {
	op := options.NewConsumerOptions()
	for _, opt := range opts {
		opt(op)
	}
	if len(addrs) == 0 {
		panic("rocket address error")
	}
	consumer, err := consumer.Connect(context.TODO(), addrs[0])
	if err != nil {
		panic(err)
	}

	c := &Consumer{
		opts: &options.Server{
			Addrs: addrs,
			Topic: topic,
		},
		msg: make(chan options.Message, 1),
	}
	c.consumer = consumer.Cluster(rocketmqOptions.Default).Consumer(c, op.GroupName, topic, op.Tags)
	c.consumeOpts = op
	return c
}

// Consume 消费
func (c *Consumer) Consume() error {
	c.consumer.Subscribe()
	return nil
}

// Close 关闭
func (c *Consumer) Close() error {
	return nil
}

// Msg 消息
func (c *Consumer) Msg() <-chan options.Message {
	return c.msg
}

// Options 配置
func (c *Consumer) Options() *options.ConsumerOptions {
	return c.consumeOpts
}

// AckID AckID
func (c *Consumer) AckID(msgID string) {

}

// NackID NackID
func (c *Consumer) NackID(msgID string) {

}

// ConsumeClaim ConsumeClaim
func (c *Consumer) ConsumeClaim(message <-chan *consumer.Message) {
	for i := range message {
		msg := options.Message{
			Body:   utils.StringToBytes(i.Body),
			Header: make(map[string]string),
		}
		c.msg <- msg
	}
}
