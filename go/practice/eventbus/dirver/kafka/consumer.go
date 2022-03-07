/*
 * @Author: 馬濤
 * @Date: 2021-01-25 15:30:00
 * @LastEditTime: 2021-09-18 02:44:55
 * @LastEditors: 馬濤
 * @Description:
 * @FilePath: \cf\eventbus\dirver\kafka\consumer.go
 * @MT is your father.
 */

package kafka

import (
	"fmt"

	"git.dustess.com/mk-base/eventbus/dirver/options"
	"git.dustess.com/mk-base/eventbus/dirver/utils"
	"git.dustess.com/mk-base/idempotence"
	idempotenceID "git.dustess.com/mk-base/idempotence/id"
	idempotenceStorage "git.dustess.com/mk-base/idempotence/storage"
	"git.dustess.com/mk-base/kafka-driver/consumer"
	"git.dustess.com/mk-base/log"
	"github.com/Shopify/sarama"
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
	c := &Consumer{
		opts: &options.Server{
			Addrs: addrs,
			Topic: topic,
		},
		msg: make(chan options.Message, 1),
	}
	c.consumer = consumer.NewConsumer(addrs, topic, op.GroupName, c)
	c.consumeOpts = op
	return c
}

// Consume 消费
func (c *Consumer) Consume() error {
	return c.consumer.Consume()
}

// Close 关闭
func (c *Consumer) Close() error {
	return c.consumer.GracefulShutdownWithError()
}

// Msg 消息
func (c *Consumer) Msg() <-chan options.Message {
	return c.msg
}

// Options 配置
func (c *Consumer) Options() *options.ConsumerOptions {
	return c.consumeOpts
}

// Setup is run at the beginning of a new session, before ConsumeClaim.
func (c *Consumer) Setup(_ sarama.ConsumerGroupSession) error {
	fmt.Println("Setup")
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
// but before the offsets are committed for the very last time.
func (c *Consumer) Cleanup(_ sarama.ConsumerGroupSession) error {
	fmt.Println("Cleanup")
	return nil
}

// AckID AckID
func (c *Consumer) AckID(msgID string) {

}

// NackID NackID
func (c *Consumer) NackID(msgID string) {

}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
// Once the Messages() channel is closed, the Handler must finish its processing
// loop and exit.
func (c *Consumer) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for payload := range claim.Messages() {
		log.Infof("eventbus kafka ConsumeClaim topic[%v] group[%v] msg[%+v]", c.opts.Topic, c.consumeOpts.GroupName, utils.BytesToString(payload.Value))
		header := make(map[string]string)
		for i := range payload.Headers[:] {
			header[utils.BytesToString(payload.Headers[i].Key)] = utils.BytesToString(payload.Headers[i].Value)
		}
		msg := options.Message{
			Body:   payload.Value,
			Header: header,
		}
		ideID, ok := header[options.IdempotenceIDField]
		if c.consumeOpts.Idempotence && ok {
			// 检测幂等
			idempotenceService := idempotence.NewIdempotence(idempotenceStorage.Redis)
			if !idempotenceService.SaveIfAbsent(idempotenceID.IdempotenceID(ideID)) {
				sess.MarkMessage(payload, "done")
				continue
			}
		}
		c.msg <- msg
		// 确认消息
		sess.MarkMessage(payload, "done")
	}
	return nil
}
