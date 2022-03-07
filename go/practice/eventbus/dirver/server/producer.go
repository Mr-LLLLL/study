/*
 * @Author: 馬濤
 * @Date: 2021-01-23 10:14:49
 * @LastEditTime: 2021-09-16 01:56:41
 * @LastEditors: 馬濤
 * @Description:
 * @FilePath: \cf\eventbus\dirver\server\producer.go
 * @MT is your father.
 */

package server

import (
	"context"

	"git.dustess.com/mk-base/eventbus/dirver/dmq"
	"git.dustess.com/mk-base/eventbus/dirver/kafka"
	"git.dustess.com/mk-base/eventbus/dirver/options"
	"git.dustess.com/mk-base/eventbus/dirver/pulsar"
	"git.dustess.com/mk-base/eventbus/dirver/rocketmq"
)

// ClientDirver ClientDirver
type ClientDirver int8

const (
	_ ClientDirver = iota
	// KafkaDirver KafkaDirver
	KafkaDirver
	// RocketMQDirver RocketMQDirver
	RocketMQDirver
	// DMQDirver 延迟队列
	DMQDirver
	// PulsarDirver 延迟队列
	PulsarDirver
)

// PublisherDirver PublisherDirver
type PublisherDirver interface {
	SubscribePublisherDirver(dirver ClientDirver, addrs []string, topic string, opts ...options.ProducerOption)
}

// Producer 生产者
type Producer interface {
	Send(ctx context.Context, msg [][]byte, header map[string]string, opts ...options.ProducerOption) error
	Close() error
}

// NewProducer NewProducer
func NewProducer(dirver ClientDirver, addrs []string, topic string, opts ...options.ProducerOption) Producer {
	switch dirver {
	case KafkaDirver:
		return kafka.NewProducer(addrs, topic, opts...)
	case RocketMQDirver:
		return rocketmq.NewProducer(addrs, topic, opts...)
	case DMQDirver:
		return dmq.NewProducer(addrs, topic, opts...)
	case PulsarDirver:
		return pulsar.NewProducer(addrs, topic, opts...)
	}
	return nil
}
