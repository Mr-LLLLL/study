/*
 * @Author: 馬濤
 * @Date: 2021-01-23 10:58:09
 * @LastEditTime: 2021-09-14 15:32:45
 * @LastEditors: 馬濤
 * @Description:
 * @FilePath: \cf\eventbus\dirver\options\producer.go
 * @MT is your father.
 */

package options

import (
	"time"
)

// ProducerOptions ProducerOptions
type ProducerOptions struct {
	// 是否异步发送消息
	IsAsync bool
	// 发送的tag rocketmq
	Tag string
	// 是否顺序消费，需要同步
	Order bool
	// Partition 分区 （初始化生效）
	Partition int64
	// DMQDriver 延迟队列驱动 （初始化生效）
	DMQDriver int64
	// DelayTime 延迟时间 单位 Duration
	DelayTime time.Duration
	// Token Token
	Token string
}

// NewProducerOptions NewProducerOptions
func NewProducerOptions() *ProducerOptions {
	opt := &ProducerOptions{
		IsAsync: false,
	}
	return opt
}

// ProducerOption 生产者配置
type ProducerOption func(*ProducerOptions)

// WithIsAsync 是否异步发送
/**
 * @description:  是否异步发送
 * @version:
 * @author: 馬濤
 * @param {bool} isAsync
 * @return {*}
 */
func WithIsAsync(isAsync bool) ProducerOption {
	return func(opts *ProducerOptions) {
		opts.IsAsync = isAsync
	}
}

// WithTag 发送的tag
/**
 * @description:  发送的tag
 * @version:
 * @author: 馬濤
 * @param {string} tag
 * @return {*}
 */
func WithTag(tag string) ProducerOption {
	return func(opts *ProducerOptions) {
		opts.Tag = tag
	}
}

// WithOrder  是否顺序消费，顺序消费，需要注册同步消费者
/**
 * @description: 是否顺序消费
 * @version:
 * @author: 馬濤
 * @param {string} tag
 * @return {*}
 */
func WithOrder(order bool) ProducerOption {
	return func(opts *ProducerOptions) {
		opts.Order = order
	}
}

// WithPartition 设置分区数
/**
 * @description: 设置分区数  partition must 2^n
 * @version:
 * @author: 馬濤
 * @param {string} tag
 * @return {*}
 */
func WithPartition(partition int64) ProducerOption {
	return func(opts *ProducerOptions) {
		opts.Partition = partition
	}
}

// WithDMQDriver 设置延迟队列驱动
/**
 * @description: 设置延迟队列驱动 1 redis 2mongo
 * @version:
 * @author: 馬濤
 * @param {string} tag
 * @return {*}
 */
func WithDMQDriver(dmqDriver int64) ProducerOption {
	return func(opts *ProducerOptions) {
		opts.DMQDriver = dmqDriver
	}
}

// WithDelayTime 设置延迟时间 Duration
/**
 * @description: 设置延迟时间 Duration
 * @version:
 * @author: 馬濤
 * @param {string} tag
 * @return {*}
 */
func WithDelayTime(delayTime time.Duration) ProducerOption {
	return func(opts *ProducerOptions) {
		opts.DelayTime = delayTime
	}
}

// WithDelayTime 设置Token
/**
 * @description: 设置Token
 * @version:
 * @author: 馬濤
 * @param {string} tag
 * @return {*}
 */
func WithProducerToken(token string) ProducerOption {
	return func(opts *ProducerOptions) {
		opts.Token = token
	}
}
