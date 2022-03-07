/*
 * @Author: 馬濤
 * @Date: 2021-01-23 10:57:48
 * @LastEditTime: 2021-09-18 02:26:18
 * @LastEditors: 馬濤
 * @Description:
 * @FilePath: \cf\eventbus\dirver\options\consumer.go
 * @MT is your father.
 */

package options

import (
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

// ConsumerOptions ConsumerOptions
type ConsumerOptions struct {
	// 消费者组名 pulsar对应的SubscriptionName
	GroupName string
	// 标签
	Tags []string
	// 幂等 pulsar无效
	Idempotence bool
	// Partition 分区
	Partition int64
	// DMQDriver 延迟队列驱动设 1 redis 2mongo
	DMQDriver int64
	// DelayTimeInterval 延迟时间间隔 Duration
	DelayTimeInterval time.Duration
	// DelayDataLimit 延迟数据拉取数据条数
	DelayDataLimit int64

	// SubscriptionType pulsar订阅类型 0 独占 1 共享订阅模式 2故障转移 3KeyShared
	/**
	Exclusive
	独占，同一主题上只能有一个具有相同订阅名的使用者
	Shared
	共享订阅模式，多个使用者将能够使用相同的订阅名称,消息将根据连接的消费者之间的循环循环
	Failover
	故障转移订阅模式下，多个使用者将能够使用相同的订阅名称,但是只有一个消费者会接收消息。如果该消费者断开连接，其他已连接的消费者中的一个将开始接收消息。
	KeyShared
	KeyShared订阅模式，多个消费者将能够使用相同,具有相同密钥的订阅和所有消息将仅分发给一个使用者
	*/
	SubscriptionType pulsar.SubscriptionType

	// Token Token
	Token string
	// RetryEnable 是否重试
	RetryEnable bool
	// MaxDeliveries 发送到死信号队列前重制的最大次数，和RetryLetterTopic和RetryEnable配合使用
	MaxDeliveries uint32
	// DeadLetterTopic 死信topic，消息失败之后，投递的topic
	DeadLetterTopic string
	// RetryLetterTopic 重试topic，RetryEnable==true必须设置，否则系统自动创建
	RetryLetterTopic string
}

// NewConsumerOptions NewConsumerOptions
func NewConsumerOptions() *ConsumerOptions {
	opt := &ConsumerOptions{
		DelayDataLimit:    1,
		DelayTimeInterval: 1 * time.Second,
		DMQDriver:         2,
		Partition:         8,
		SubscriptionType:  1,
	}
	return opt
}

// ConsumerOption 消费者配置
type ConsumerOption func(*ConsumerOptions)

// WithGroupName 消费者组
/**
 * @description:  消费者组
 * @version:
 * @author: 馬濤
 * @param {string} groupName  消费者组
 * @return {*}
 */
func WithGroupName(groupName string) ConsumerOption {
	return func(opts *ConsumerOptions) {
		opts.GroupName = groupName
	}
}

// WithTags 消费者标签
/**
 * @description: 消费者标签
 * @version:
 * @author: 馬濤
 * @param {[]string} tags
 * @return {*}
 */
func WithTags(tags []string) ConsumerOption {
	return func(opts *ConsumerOptions) {
		opts.Tags = tags
	}
}

// WithIdempotence 是否幂等校验
/**
 * @description: 是否幂等校验
 * @version:
 * @author: 馬濤
 * @param {[]string} idempotence bool
 * @return {*}
 */
func WithIdempotence(idempotence bool) ConsumerOption {
	return func(opts *ConsumerOptions) {
		opts.Idempotence = idempotence
	}
}

// WithConsumePartition 设置分区数  partition must 2^n
/**
 * @description: 设置分区数
 * @version:
 * @author: 馬濤
 * @param {string} tag
 * @return {*}
 */
func WithConsumePartition(partition int64) ConsumerOption {
	return func(opts *ConsumerOptions) {
		opts.Partition = partition
	}
}

// WithConsumeDMQDriver 设置延迟队列驱动
/**
 * @description: 设置延迟队列驱动 1 redis 2mongo
 * @version:
 * @author: 馬濤
 * @param {string} tag
 * @return {*}
 */
func WithConsumeDMQDriver(dmqDriver int64) ConsumerOption {
	return func(opts *ConsumerOptions) {
		opts.DMQDriver = dmqDriver
	}
}

// WithDelayTimeInterval 延迟时间间隔 Duration
/**
 * @description: 延迟时间间隔 Duration
 * @version:
 * @author: 馬濤
 * @param {delayTimeInterval} time.Duration
 * @return {*}
 */
func WithDelayTimeInterval(delayTimeInterval time.Duration) ConsumerOption {
	return func(opts *ConsumerOptions) {
		opts.DelayTimeInterval = delayTimeInterval
	}
}

// WithDelayDataLimit 延迟数据拉取数据条数
/**
 * @description: 延迟数据拉取数据条数
 * @version:
 * @author: 馬濤
 * @param {int64} delayDataLimit
 * @return {*}
 */
func WithDelayDataLimit(delayDataLimit int64) ConsumerOption {
	return func(opts *ConsumerOptions) {
		opts.DelayDataLimit = delayDataLimit
	}
}

// WithSubscriptionType pulsar订阅类型 0 独占 1 共享订阅模式 2故障转移 3 KeyShared
/**
 * @description: pulsar订阅类型 0 独占 1 共享订阅模式 2故障转移 3 KeyShared
 * @version:
 * @author: 馬濤
 * @param {int} subscriptionType
 * @return {*}
 */
func WithSubscriptionType(subscriptionType int) ConsumerOption {
	return func(opts *ConsumerOptions) {
		opts.SubscriptionType = pulsar.SubscriptionType(subscriptionType)
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
func WithConsumerToken(token string) ConsumerOption {
	return func(opts *ConsumerOptions) {
		opts.Token = token
	}
}

// WithRetryEnable 设置重试（pulsar）
func WithRetryEnable(isEnable bool) ConsumerOption {
	return func(opts *ConsumerOptions) {
		opts.RetryEnable = isEnable
	}
}

// WithMaxDeliveries 发送到死信号队列前重制的最大次数，和RetryLetterTopic和RetryEnable配合使用（pulsar）
func WithMaxDeliveries(maxDeliveries uint32) ConsumerOption {
	return func(opts *ConsumerOptions) {
		opts.MaxDeliveries = maxDeliveries
	}
}

// WithDeadLetterTopic 死信topic，消息失败之后，投递的topic（pulsar）
func WithDeadLetterTopic(deadLetterTopic string) ConsumerOption {
	return func(opts *ConsumerOptions) {
		opts.DeadLetterTopic = deadLetterTopic
	}
}

// WithRetryLetterTopic 重试topic，RetryEnable==true必须设置，否则系统自动创建（pulsar）
func WithRetryLetterTopic(retryLetterTopic string) ConsumerOption {
	return func(opts *ConsumerOptions) {
		opts.RetryLetterTopic = retryLetterTopic
	}
}
