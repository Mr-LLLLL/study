eventbus
======

Eventbus 允许组件之间进行发布-订阅式通信，而无需组件之间显式注册。支持服务内和服务外，服务外目前支持kafka

### 安装

go get git.dustess.com/mk-base/eventbus

### 使用

#### 服务内部

```go

import (
 "git.dustess.com/mk-base/eventbus"
)

func main() {
    bus := eventbus.New()
    // 添加一个异步订阅者
    bus.SubscribeAsync(NewEventUpdateOrderPriceGroup(), false)
    // 往topic：test111111111 发送参数 xxx  xxx的类型个数必须和阅者触发时执行的方法接受参数类型个数一致
    bus.SubscribeRegisterWait()
 ctx:=context.TODO()
    bus.PublishCtx(ctx,"test111111111", "xxx")

}

/**
订阅者组件结构体必须要有event字段
tag含义及作用：
topic：订阅的topic，tag内容就是定义的key
subscribe：订阅者触发时执行的方法，tag内容就是定义检测方法的名字，自定义
quit：eventbus退出时执行的函数，tag内容就是定义的key，自定义
*/


// EventUpdateOrderPriceGroup 修改订单价格 团购
type EventUpdateOrderPriceGroup struct {
 event interface{} `subscribe:"Update" topic:"test111111111" quit:"Quit"`
}

// NewEventUpdateOrderPriceGroup 修改订单价格 团购
func NewEventUpdateOrderPriceGroup() *EventUpdateOrderPriceGroup {
 return new(EventUpdateOrderPriceGroup)
}

// Update  修改订单价格 团购   最少有一个ctx参数
func (event *EventUpdateOrderPriceGroup) Update(ctx context.Context,ccc string) {
 time.Sleep(2 * time.Second)
 fmt.Println(ccc)
}

// Quit 退出
func (event *EventUpdateOrderPriceGroup) Quit() {

}
```

### 自定义Topic和Function

```go
/// 实现 Topic 优先使用函数返回的topic
// Topic Topic
func (event *EventUpdateOrderPriceGroup) Topic() string {
 return "testTopic"
}

// 实现 Function 优先使用函数返回的Function
// Function Function
func (event *EventUpdateOrderPriceGroup) Function() string {
 return "testFunc"
}

```

#### 服务外部

##### kafka

```go

import (
    "git.dustess.com/mk-base/eventbus"
    "git.dustess.com/mk-base/eventbus/dirver/options"
 "git.dustess.com/mk-base/eventbus/dirver/server"
    "git.dustess.com/mk-base/idempotence/storage"
 "context"
)

func main() {
    // 初始化幂等框架
 storage.Init()
 storage.SetRedisPoolConfig(&storage.PoolTimeConfig{
  InitialCap:  5,
  MaxCap:      200,
  MaxIdle:     100,
  IdleTimeout: 30 * time.Second,
 })
    bus := eventbus.New()
    // 添加一个异步订阅者
    bus.SubscribeAsync(NewEventUpdateOrderPriceGroup(), false)
    addr := []string{
  "192.168.0.124:9092",
  "192.168.0.127:9092",
  "192.168.0.128:9092",
    }
    // 添加一个kafka订阅者,理解为kafka的生产者，需要发送
    topic:="test111111111"
    bus.SubscribePublisherDirver(server.KafkaDirver, addr, topic)
    // 添加一个kafka订阅者,目前只支持kafka 理解为kafka的消费者，需要发送
    bus.SubscribeConsumerDirver(server.KafkaDirver, addr, topic,
        // 增加消费者组的配置项，值就是消费者组
  options.WithGroupName("testGroup"),
  // 增加消息幂等
  options.WithIdempotence(true),
    )
    bus.SubscribeRegisterWait()
    // 往topic发送消息，如果是kafka的topic会自动发送到kafka 服务外的格式目前是两个参数 [][]byte和map[string]string 都是必传
    header = make(map[string]string)
 ctx:=context.TODO()
 // 只能用PublishWithCtx
    bus.PublishWithCtx(ctx,"test111111111", [][]byte{[]byte{12}},header)
}

/**
订阅者组件结构体必须要有event字段
tag含义及作用：
topic：订阅的topic，tag内容就是定义的key  订阅kafka的消息topic为  "kafka的topic"+"_"+"消费者组的"
subscribe：订阅者触发时执行的方法，tag内容就是定义检测方法的名字，自定义
quit：eventbus退出时执行的函数，tag内容就是定义的key，自定义
*/


// EventUpdateOrderPriceGroup 修改订单价格 团购
type EventUpdateOrderPriceGroup struct {
 event interface{} `subscribe:"Update" topic:"test111111111_testGroup" quit:"Quit"`
}

// NewEventUpdateOrderPriceGroup 修改订单价格 团购  
func NewEventUpdateOrderPriceGroup() *EventUpdateOrderPriceGroup {
 return new(EventUpdateOrderPriceGroup)
}

// Update  修改订单价格 团购 服务外的格式目前只支持 两个参数并且类型是 []byte 和map[string]string
func (event *EventUpdateOrderPriceGroup) Update(ccc []byte, header map[string]string) {
 time.Sleep(2 * time.Second)
 fmt.Println(ccc)
}

// Quit 退出
func (event *EventUpdateOrderPriceGroup) Quit() {

}

```

##### pulsar

###### 常规发送

```go

// 常规发送
import (
    "git.dustess.com/mk-base/eventbus"
    "git.dustess.com/mk-base/eventbus/dirver/options"
 "git.dustess.com/mk-base/eventbus/dirver/server"
    "git.dustess.com/mk-base/idempotence/storage"
 "context"
)

func main() {
    // 初始化幂等框架
 storage.Init()
 storage.SetRedisPoolConfig(&storage.PoolTimeConfig{
  InitialCap:  5,
  MaxCap:      200,
  MaxIdle:     100,
  IdleTimeout: 30 * time.Second,
 })
    bus := eventbus.New()
    // 添加一个异步订阅者
    bus.SubscribeAsync(NewEventUpdateOrderPriceGroup(), false)
addr := []string{
  "http://pulsar-vgp2awkmr4ne.tdmq-pulsar.ap-sh.public.tencenttdmq.com:8080",
 }
 token := "eyJrZXlJZCI6InB1bHNhci12Z3AyYXdrbXI0bmUiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJwdWxzYXItdmdwMmF3a21yNG5lX2FsaS10ZXN0In0.yWu-zTsJtgx_bSBzYrOcYgX-FmCQeKyuiYo_HrcmwF8"
    // 添加一个pulsar订阅者,理解为pulsar的生产者，需要发送
    topic:="pulsar-vgp2awkmr4ne/ali-test-mk/mall-test"
 bus.SubscribePublisherDirver(server.PulsarDirver, addr, topic,
  options.WithProducerToken(token),// token
 )                                 
    // 添加一个kafka订阅者,目前只支持kafka 理解为kafka的消费者，需要发送
    bus.SubscribeConsumerDirver(server.PulsarDirver, addr, topic,
  // 增加消费者组的配置项，值就是消费者组
  options.WithGroupName("testGroup"),
        // token                        
  options.WithConsumerToken(token),
 )
    bus.SubscribeRegisterWait()
    // 往topic发送消息，格式目前是两个参数 [][]byte和map[string]string 都是必传
    header = make(map[string]string)
 ctx:=context.TODO()
 // 只能用PublishWithCtx
    bus.PublishWithCtx(ctx,topic, [][]byte{[]byte{12}},header)
}

/**
订阅者组件结构体必须要有event字段
tag含义及作用：
topic：订阅的topic，tag内容就是定义的key  订阅kafka的消息topic为  "kafka的topic"+"_"+"消费者组的"
subscribe：订阅者触发时执行的方法，tag内容就是定义检测方法的名字，自定义
quit：eventbus退出时执行的函数，tag内容就是定义的key，自定义
*/


// EventUpdateOrderPriceGroup 修改订单价格 团购
type EventUpdateOrderPriceGroup struct {
 event interface{} `subscribe:"Update" topic:"pulsar-vgp2awkmr4ne/ali-test-mk/mall-test_testGroup" quit:"Quit"`
}

// NewEventUpdateOrderPriceGroup 修改订单价格 团购  
func NewEventUpdateOrderPriceGroup() *EventUpdateOrderPriceGroup {
 return new(EventUpdateOrderPriceGroup)
}

// Update  修改订单价格 团购 服务外的格式目前只支持 两个参数并且类型是 []byte 和map[string]string
func (event *EventUpdateOrderPriceGroup) Update(ccc []byte, header map[string]string) {
 time.Sleep(2 * time.Second)
 fmt.Println(ccc)
}

// Quit 退出
func (event *EventUpdateOrderPriceGroup) Quit() {

}

```

###### 故障转移

需要主动确认

```go

import (
    "git.dustess.com/mk-base/eventbus"
    "git.dustess.com/mk-base/eventbus/dirver/options"
 "git.dustess.com/mk-base/eventbus/dirver/server"
    "git.dustess.com/mk-base/idempotence/storage"
 "context"
)

func main() {
    // 初始化幂等框架
 storage.Init()
 storage.SetRedisPoolConfig(&storage.PoolTimeConfig{
  InitialCap:  5,
  MaxCap:      200,
  MaxIdle:     100,
  IdleTimeout: 30 * time.Second,
 })
bus := eventbus.New()
// 添加一个异步订阅者
bus.SubscribeAsync(NewEventUpdateOrderPriceGroup(), false)
addr := []string{
  "http://pulsar-vgp2awkmr4ne.tdmq-pulsar.ap-sh.public.tencenttdmq.com:8080",
 }
 token := "eyJrZXlJZCI6InB1bHNhci12Z3AyYXdrbXI0bmUiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJwdWxzYXItdmdwMmF3a21yNG5lX2FsaS10ZXN0In0.yWu-zTsJtgx_bSBzYrOcYgX-FmCQeKyuiYo_HrcmwF8"
// 添加一个pulsar订阅者,理解为pulsar的生产者，需要发送
topic:="pulsar-vgp2awkmr4ne/ali-test-mk/mall-test"
bus.SubscribePublisherDirver(server.PulsarDirver, addr, topic,
  options.WithProducerToken(token),// token
 )                                 
// 添加一个kafka订阅者,目前只支持kafka 理解为kafka的消费者，需要发送
bus.SubscribeConsumerDirver(server.PulsarDirver, addr, topic,
// 增加消费者组的配置项，值就是消费者组
options.WithGroupName("testGroup"),
// token                        
options.WithConsumerToken(token),
// pulsar订阅类型 0 独占 1 共享订阅模式(可以用做延迟，其他都不行) 2故障转移 3 KeyShared
options.WithSubscriptionType(2),
// RetryEnable 是否重试
options.WithRetryEnable(true),
 )
bus.SubscribeRegisterWait()
// 往topic发送消息，格式目前是两个参数 [][]byte和map[string]string 都是必传
header = make(map[string]string)
 ctx:=context.TODO()
 // 只能用PublishWithCtx
bus.PublishWithCtx(ctx,topic, [][]byte{[]byte{12}},header)
}

/**
订阅者组件结构体必须要有event字段
tag含义及作用：
topic：订阅的topic，tag内容就是定义的key  订阅kafka的消息topic为  "kafka的topic"+"_"+"消费者组的"
subscribe：订阅者触发时执行的方法，tag内容就是定义检测方法的名字，自定义
quit：eventbus退出时执行的函数，tag内容就是定义的key，自定义
*/


// EventUpdateOrderPriceGroup 修改订单价格 团购
type EventUpdateOrderPriceGroup struct {
 event interface{} `subscribe:"Update" topic:"pulsar-vgp2awkmr4ne/ali-test-mk/mall-test_testGroup" quit:"Quit"`
}

// NewEventUpdateOrderPriceGroup 修改订单价格 团购  
func NewEventUpdateOrderPriceGroup() *EventUpdateOrderPriceGroup {
 return new(EventUpdateOrderPriceGroup)
}

// Update  修改订单价格 团购 服务外的格式目前只支持 两个参数并且类型是 []byte 和map[string]string
func (event *EventUpdateOrderPriceGroup) Update(ccc []byte, header map[string]string) {
id, ok := header[options.ValueIDField]
if ok {
    // 主动确认，用id
    // 参数传入 驱动 topic，消费者组，id
 defer eventbus.EventBusHandler.Ack(server.PulsarDirver, "pulsar-vgp2awkmr4ne/ali-test-mk/mall-test", "testGroup", id)
}
 time.Sleep(2 * time.Second)
 fmt.Println(ccc)
}

// Quit 退出
func (event *EventUpdateOrderPriceGroup) Quit() {

}

```

###### 延迟队列

```go

// 常规发送
import (
    "git.dustess.com/mk-base/eventbus"
    "git.dustess.com/mk-base/eventbus/dirver/options"
 "git.dustess.com/mk-base/eventbus/dirver/server"
    "git.dustess.com/mk-base/idempotence/storage"
 "context"
)

func main() {
    // 初始化幂等框架
 storage.Init()
 storage.SetRedisPoolConfig(&storage.PoolTimeConfig{
  InitialCap:  5,
  MaxCap:      200,
  MaxIdle:     100,
  IdleTimeout: 30 * time.Second,
 })
    bus := eventbus.New()
    // 添加一个异步订阅者
    bus.SubscribeAsync(NewEventUpdateOrderPriceGroup(), false)
addr := []string{
  "http://pulsar-vgp2awkmr4ne.tdmq-pulsar.ap-sh.public.tencenttdmq.com:8080",
 }
 token := "eyJrZXlJZCI6InB1bHNhci12Z3AyYXdrbXI0bmUiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJwdWxzYXItdmdwMmF3a21yNG5lX2FsaS10ZXN0In0.yWu-zTsJtgx_bSBzYrOcYgX-FmCQeKyuiYo_HrcmwF8"
    // 添加一个pulsar订阅者,理解为pulsar的生产者，需要发送
    topic:="pulsar-vgp2awkmr4ne/ali-test-mk/mall-test"
 bus.SubscribePublisherDirver(server.PulsarDirver, addr, topic,
  options.WithProducerToken(token),// token
 )                                 
    // 添加一个kafka订阅者,目前只支持kafka 理解为kafka的消费者，需要发送
    bus.SubscribeConsumerDirver(server.PulsarDirver, addr, topic,
  // 增加消费者组的配置项，值就是消费者组
  options.WithGroupName("testGroup"),
        // token                        
  options.WithConsumerToken(token),
  // pulsar订阅类型 0 独占 1 共享订阅模式(可以用做延迟，其他都不行) 2故障转移 3 KeyShared
  options.WithSubscriptionType(1),
 )
    bus.SubscribeRegisterWait()
    // 往topic发送消息，格式目前是两个参数 [][]byte和map[string]string 都是必传
    header = make(map[string]string)
 ctx:=context.TODO()
 // 只能用PublishWithCtx
    bus.PublishWithCtx(ctx,topic, [][]byte{[]byte{12}},header, 
 options.WithDelayTime(10*time.Second),// 消息延迟时间 单位 Duration
 )
}

/**
订阅者组件结构体必须要有event字段
tag含义及作用：
topic：订阅的topic，tag内容就是定义的key  订阅kafka的消息topic为  "kafka的topic"+"_"+"消费者组的"
subscribe：订阅者触发时执行的方法，tag内容就是定义检测方法的名字，自定义
quit：eventbus退出时执行的函数，tag内容就是定义的key，自定义
*/


// EventUpdateOrderPriceGroup 修改订单价格 团购
type EventUpdateOrderPriceGroup struct {
 event interface{} `subscribe:"Update" topic:"pulsar-vgp2awkmr4ne/ali-test-mk/mall-test_testGroup" quit:"Quit"`
}

// NewEventUpdateOrderPriceGroup 修改订单价格 团购  
func NewEventUpdateOrderPriceGroup() *EventUpdateOrderPriceGroup {
 return new(EventUpdateOrderPriceGroup)
}

// Update  修改订单价格 团购 服务外的格式目前只支持 两个参数并且类型是 []byte 和map[string]string
func (event *EventUpdateOrderPriceGroup) Update(ccc []byte, header map[string]string) {
 time.Sleep(2 * time.Second)
 fmt.Println(ccc)
}

// Quit 退出
func (event *EventUpdateOrderPriceGroup) Quit() {

}

```

#### DMQ延迟队列

```go
//延迟队列
import (
    "git.dustess.com/mk-base/eventbus"
    "git.dustess.com/mk-base/eventbus/dirver/options"
 "git.dustess.com/mk-base/eventbus/dirver/server"
    "git.dustess.com/mk-base/idempotence/storage"
 "context"
)

func main() {
    // 初始化幂等框架
 storage.Init()
 storage.SetRedisPoolConfig(&storage.PoolTimeConfig{
  InitialCap:  5,
  MaxCap:      200,
  MaxIdle:     100,
  IdleTimeout: 30 * time.Second,
 })
    bus := eventbus.New()
    // 添加一个异步订阅者
    bus.SubscribeAsync(NewEventUpdateOrderPriceGroup(), false)
    addr := []string{}
 // 添加一个kafka订阅者,目前只支持kafka 理解为kafka的生产者，需要发送
 topic := "dmq_test"
 eventbus.EventBusHandler.SubscribePublisherDirver(server.DMQDirver, addr, topic,
  options.WithPartition(8), // 设置分区数 必须是 2^n （partition must 2^n）
  options.WithDMQDriver(2), // 1 redis 2mongo 设置驱动（推荐mongo）
 )
 // 添加一个延迟队列订阅者
 bus.SubscribeConsumerDirver(server.DMQDirver, addr, topic,
  // 增加消息幂等
  options.WithIdempotence(true),
   // 设置分区数 必须是 2^n （partition must 2^n）
  options.WithConsumePartition(8),
  // 1 redis 2mongo 设置驱动（推荐mongo）
  options.WithConsumeDMQDriver(2),
  // 延迟时间间隔 单位 Duration
  options.WithDelayTimeInterval(1*time.Second),
  // 设置消费者组（必须）
  options.WithGroupName("ts"),
  // 消费者每个分区取的条数 目前只要·redis支持
     options.WithDelayDataLimit(100),
 )
 
 bus.SubscribeRegisterWait()
 header := make(map[string]string)
 ctx := context.TODO()
 bus.Preview()
 // 只能用PublishCtx
 fmt.Println(time.Now())
 // 消息体类型是  [][]byte 可以批量发送
 bus.PublishWithCtx(ctx, topic, [][]byte{[]byte{12}}, header, 
 // 消息延迟时间 单位 Duration
 options.WithDelayTime(10*time.Second)，
 )
}

/**
订阅者组件结构体必须要有event字段
tag含义及作用：
topic：订阅的topic，tag内容就是定义的key  订阅消息topic为  "topic"+"_"+"消费者组的"
subscribe：订阅者触发时执行的方法，tag内容就是定义检测方法的名字，自定义
quit：eventbus退出时执行的函数，tag内容就是定义的key，自定义
*/

// EventUpdateOrderPriceGroup 修改订单价格 团购
type EventUpdateOrderPriceGroup struct {
 event interface{} `subscribe:"Update" topic:"test111111111_testGroup" quit:"Quit"`
}

// NewEventUpdateOrderPriceGroup 修改订单价格 团购  
func NewEventUpdateOrderPriceGroup() *EventUpdateOrderPriceGroup {
 return new(EventUpdateOrderPriceGroup)
}

// Update  修改订单价格 团购 服务外的格式目前只支持 两个参数并且类型是 []byte 和map[string]string
func (event *EventUpdateOrderPriceGroup) Update(ccc []byte, header map[string]string) {
 time.Sleep(2 * time.Second)
 fmt.Println(ccc)
}

// Quit 退出
func (event *EventUpdateOrderPriceGroup) Quit() {

}

```
