/*
 * @Author: 馬濤
 * @Date: 2021-01-23 10:14:49
 * @LastEditTime: 2021-09-24 01:22:27
 * @LastEditors: 馬濤
 * @Description:
 * @FilePath: \cf\eventbus\dirver\server\consumer.go
 * @MT is your father.
 */

package server

import (
	"git.dustess.com/mk-base/eventbus/dirver/dmq"
	"git.dustess.com/mk-base/eventbus/dirver/kafka"
	"git.dustess.com/mk-base/eventbus/dirver/options"
	"git.dustess.com/mk-base/eventbus/dirver/pulsar"
	"git.dustess.com/mk-base/eventbus/dirver/rocketmq"
)

// Consumer 生产者
type Consumer interface {
	Consume() error
	Close() error
	Msg() <-chan options.Message
	Options() *options.ConsumerOptions
	AckID(msgID string)
	NackID(msgID string)
}

// ConsumerDirver ConsumerDirver
type ConsumerDirver interface {
	SubscribeConsumerDirver(dirver ClientDirver, addrs []string, topic string, opts ...options.ConsumerOption)
}

// NewConsumer NewConsumer
func NewConsumer(dirver ClientDirver, addrs []string, topic string, opts ...options.ConsumerOption) Consumer {
	switch dirver {
	case KafkaDirver:
		return kafka.NewConsumer(addrs, topic, opts...)
	case RocketMQDirver:
		return rocketmq.NewConsumer(addrs, topic, opts...)
	case DMQDirver:
		return dmq.NewConsumer(addrs, topic, opts...)
	case PulsarDirver:
		return pulsar.NewConsumer(addrs, topic, opts...)
	}
	return nil
}
