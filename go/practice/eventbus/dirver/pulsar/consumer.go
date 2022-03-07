/*
 * @Author: 馬濤
 * @Date: 2021-01-25 15:30:00
 * @LastEditTime: 2021-09-27 00:52:47
 * @LastEditors: 馬濤
 * @Description:
 * @FilePath: \cf\eventbus\dirver\pulsar\consumer.go
 * @MT is your father.
 */

package pulsar

import (
	"sync"

	"git.dustess.com/mk-base/eventbus/dirver/options"
	"git.dustess.com/mk-base/eventbus/dirver/utils"
	"git.dustess.com/mk-base/log"
	"git.dustess.com/shared-golib/pulsar-driver/consumer"
	"github.com/apache/pulsar-client-go/pulsar"
)

// Consumer Kafka
type Consumer struct {
	opts          *options.Server
	consumer      *consumer.Consumer
	consumeOpts   *options.ConsumerOptions
	msg           chan options.Message
	pulsarHandler *PulsarConsumer
}

// NewConsumer NewConsumer
func NewConsumer(addrs []string, topic string, opts ...options.ConsumerOption) *Consumer {
	op := options.NewConsumerOptions()
	for _, opt := range opts {
		opt(op)
	}
	c := &Consumer{
		opts: &options.Server{
			Addrs: addrs,
			Topic: topic,
		},
		msg:           make(chan options.Message, 1),
		pulsarHandler: new(PulsarConsumer),
	}
	c.consumeOpts = op
	// 配置
	conf := consumer.Config{
		Addrs:            addrs,
		Topic:            topic,
		Name:             "",
		SubscriptionName: op.GroupName,
		Token:            op.Token,
		SubscriptionType: op.SubscriptionType,
		RetryEnable:      op.RetryEnable,
		MaxDeliveries:    op.MaxDeliveries,
		DeadLetterTopic:  op.DeadLetterTopic,
		RetryLetterTopic: op.RetryLetterTopic,
	}

	var err error
	pulsarHandler := &PulsarConsumer{
		opts:        c.opts,
		consumeOpts: c.consumeOpts,
		msg:         c.msg,
		pool:        new(sync.Map),
	}
	c.consumer, err = consumer.NewConsumer(conf, pulsarHandler)
	if err != nil {
		panic(err)
	}
	c.pulsarHandler = pulsarHandler
	c.consumeOpts = op
	return c
}

// Consume 消费
func (c *Consumer) Consume() error {
	return nil
}

// Close 关闭
func (c *Consumer) Close() error {
	c.consumer.GracefulShutdown()
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
	id, ok := c.pulsarHandler.pool.Load(msgID)
	if !ok {
		log.Warnf("eventbus pulsar AckID msgID[%+v] not found", msgID)
		return
	}
	messageID, ok := id.(pulsar.MessageID)
	c.pulsarHandler.pool.Delete(msgID)
	if !ok {
		log.Warnf("eventbus pulsar AckID msgID[%+v] not type", msgID)
		return
	}
	c.consumer.AckID(messageID)
}

// NackID NackID
func (c *Consumer) NackID(msgID string) {
	id, ok := c.pulsarHandler.pool.Load(msgID)
	if !ok {
		log.Warnf("eventbus pulsar NackID msgID[%+v] not found", msgID)
		return
	}
	c.pulsarHandler.pool.Delete(msgID)
	messageID, ok := id.(pulsar.MessageID)
	if !ok {
		log.Warnf("eventbus pulsar NackID msgID[%+v] not type", msgID)
		return
	}
	c.consumer.NackID(messageID)
}

// PulsarConsumer Pilsar
type PulsarConsumer struct {
	opts        *options.Server
	consumeOpts *options.ConsumerOptions
	msg         chan options.Message
	pool        *sync.Map
}

// Close 关闭
func (c *PulsarConsumer) Close() {

}

// ConsumeClaim ConsumeClaim
func (c *PulsarConsumer) ConsumeClaim(messages chan consumer.GroupSession) error {
	for v := range messages {
		message := v.GetMessage()
		header := message.Properties()
		log.Infof("eventbus Pulsar ConsumeClaim topic[%v] group[%v] header[%+v] msg[%+v]", c.opts.Topic, c.consumeOpts.GroupName, header, utils.BytesToString(message.Payload()))

		// 填入id
		header[options.ValueIDField] = utils.BytesToString(message.ID().Serialize())
		msg := options.Message{
			Body:   message.Payload(),
			Header: header,
		}
		c.msg <- msg
		// 不重试才自动确认
		if !c.consumeOpts.RetryEnable {
			v.Ack()
			continue
		}
		// 加入池
		c.pool.Store(header[options.ValueIDField], message.ID())
	}
	return nil
}
