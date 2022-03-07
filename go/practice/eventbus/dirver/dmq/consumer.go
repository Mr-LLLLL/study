/*
 * @Author: 馬濤
 * @Date: 2021-01-25 15:30:00
 * @LastEditTime: 2021-09-18 02:45:05
 * @LastEditors: 馬濤
 * @Description:
 * @FilePath: \cf\eventbus\dirver\dmq\consumer.go
 * @MT is your father.
 */

package dmq

import (
	dmqConfig "git.dustess.com/mk-base/dmq/config"
	"git.dustess.com/mk-base/dmq/consumer"
	"git.dustess.com/mk-base/dmq/message"
	"git.dustess.com/mk-base/eventbus/dirver/options"
	"git.dustess.com/mk-base/idempotence"
	idempotenceID "git.dustess.com/mk-base/idempotence/id"
	idempotenceStorage "git.dustess.com/mk-base/idempotence/storage"
)

// Consumer dmq
type Consumer struct {
	opts        *options.Server
	consumer    *consumer.PartitionConsumerDelay
	consumeOpts *options.ConsumerOptions
	msg         chan options.Message
}

// NewConsumer NewConsumer
func NewConsumer(addrs []string, topic string, opts ...options.ConsumerOption) *Consumer {
	op := options.NewConsumerOptions()
	for _, opt := range opts {
		opt(op)
	}
	config := dmqConfig.NewConfig(topic, int8(op.DMQDriver), uint8(op.Partition))
	c := &Consumer{
		opts: &options.Server{
			Addrs: addrs,
			Topic: topic,
		},
		msg: make(chan options.Message, 1),
	}
	c.consumeOpts = op
	dmqC := &DMQConsumer{
		opts:        c.opts,
		consumeOpts: c.consumeOpts,
		msg:         c.msg,
	}

	c.consumer = consumer.NewPartitionConsumerDelay(config, dmqC, op.DelayDataLimit)
	c.consumer.SetConsumeInterval(op.DelayTimeInterval)
	return c
}

// Consume 消费
func (c *Consumer) Consume() error {
	c.consumer.Consume()
	return nil
}

// Close 关闭
func (c *Consumer) Close() error {
	c.consumer.Close()
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

// DMQConsumer dmq
type DMQConsumer struct {
	opts        *options.Server
	consumeOpts *options.ConsumerOptions
	msg         chan options.Message
}

// Close 关闭
func (c *DMQConsumer) Close() {

}

// ConsumeClaim ConsumeClaim
func (c *DMQConsumer) ConsumeClaim(message <-chan *message.Message) {
	for msg := range message {
		dmqEsg := options.Message{
			Body:   msg.Value,
			Header: make(map[string]string),
		}
		if eventbusHeader, ok := msg.Header[EventbusHeaderField]; ok {
			ebH, ok := eventbusHeader.(map[string]interface{})
			if ok {
				ebhs := mapInterfaceToString(ebH)
				ideID, ok := ebhs[options.IdempotenceIDField]
				if c.consumeOpts.Idempotence && ok {
					// 检测幂等
					idempotenceService := idempotence.NewIdempotence(idempotenceStorage.Redis)
					if !idempotenceService.SaveIfAbsent(idempotenceID.IdempotenceID(ideID)) {
						continue
					}
				}
				dmqEsg.Header = ebhs
			}
		}
		c.msg <- dmqEsg
	}
}
func mapInterfaceToString(mi map[string]interface{}) map[string]string {
	m := make(map[string]string)
	for i, v := range mi {
		s, ok := v.(string)
		if ok {
			m[i] = s
		}
	}
	return m
}
