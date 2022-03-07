package eventbus

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"sync"

	"git.dustess.com/mk-base/eventbus/dirver/options"
	dirverServer "git.dustess.com/mk-base/eventbus/dirver/server"
	"git.dustess.com/mk-base/log"
)

//BusSubscriber 订阅
type BusSubscriber interface {
	Subscribe(observer interface{})
	SubscribeAsync(observer interface{}, transactional bool)
	SubscribeOnce(observer interface{})
	SubscribeOnceAsync(observer interface{})
	Unsubscribe(observer interface{})
}

//BusPublisher 发布
type BusPublisher interface {
	Publish(topic string, args ...interface{})
	PublishWithCtx(ctx context.Context, topic string, args ...interface{})
	Ack(dirver dirverServer.ClientDirver, topic, groupName, msgID string)
}

//BusController 检查
type BusController interface {
	SetMaxTaskNum(num int64)
	Preview() map[string]map[string]struct{}
	HasCallback(topic string) bool
	WaitAsync()
	Quit()
	SubscribeRegisterWait()
	AllTask() map[string][]*EventHandler
}

//Bus 总线
type Bus interface {
	BusController
	BusSubscriber
	BusPublisher
	dirverServer.PublisherDirver
	dirverServer.ConsumerDirver
}

// EventBus 事件总线
type EventBus struct {
	handlers         *sync.Map
	wg               *sync.WaitGroup
	lock             *sync.Mutex
	pool             *Pool
	comsumerPool     map[string]dirverServer.Consumer
	comsumerPoolLock *sync.Mutex
	subscribeWg      *sync.WaitGroup
}

// EventHandler EventHandler
type EventHandler struct {
	observerType  reflect.Type
	observer      interface{}
	callBack      reflect.Value
	quitCallBack  reflect.Value
	flagOnce      bool
	async         bool
	transactional bool
	lock          *sync.Mutex
	topic         string
	ext           string //扩展信息
}

// Ext Ext
func (event *EventHandler) Ext() string {
	return event.ext
}

// Topic Topic
type Topic interface {
	Topic() string
}

// Topic Topic
type Function interface {
	Function() string
}

// New new
func New() Bus {
	pool, err := NewPool(200)
	if err != nil {
		panic(err)
	}
	b := &EventBus{
		new(sync.Map),
		new(sync.WaitGroup),
		new(sync.Mutex),
		pool,
		make(map[string]dirverServer.Consumer, 0),
		new(sync.Mutex),
		new(sync.WaitGroup),
	}
	return Bus(b)
}

// doSubscribe 处理订阅逻辑
func (bus *EventBus) doSubscribe(topic string, handler *EventHandler) error {
	bus.lock.Lock()
	defer bus.lock.Unlock()
	handlerInterface, _ := bus.handlers.LoadOrStore(topic, make([]*EventHandler, 0))
	handlers := handlerInterface.([]*EventHandler)
	for i := range handlers[:] {
		if handlers[i].observerType.Elem() == handler.observerType.Elem() {
			return nil
		}
	}
	bus.handlers.Store(topic, append(handlers, handler))
	return nil
}

func (bus *EventBus) checkObserver(observer interface{}) ([]string, reflect.Type, string, string, string) {
	var t reflect.Type
	var fn string
	var quit string
	var ok bool
	var ext string

	t = reflect.TypeOf(observer)
	if t.Elem().Kind() != reflect.Struct {
		panic(fmt.Sprintf("%s is not of type reflect.Struct", t.Kind()))
	}
	fnField, _ := t.Elem().FieldByName("event")
	if fnField.Tag == "" {
		panic(fmt.Sprintf("%v has no field or no fn field", fnField))
	}

	fn, _ = fnField.Tag.Lookup("subscribe")
	obsFunc, ok := observer.(Function)
	if ok {
		fn = obsFunc.Function()
	}
	if fn == "" {
		panic("subscribe tag doesn't exist or empty")
	}
	quit, ok = fnField.Tag.Lookup("quit")
	if !ok || quit == "" {
		panic("quit tag doesn't exist or empty")
	}

	topics, _ := fnField.Tag.Lookup("topic")
	obsTopic, ok := observer.(Topic)
	if ok {
		topics = obsTopic.Topic()
	}
	if topics == "" {
		panic("topic tag doesn't exist or empty")
	}

	topic := strings.Split(topics, ",")
	ext, _ = fnField.Tag.Lookup("ext")
	return topic, t, fn, quit, ext
}

func (bus *EventBus) register(observer interface{}, flagOnce, async, transactional bool) {
	topic, t, fn, quit, ext := bus.checkObserver(observer)
	for i := range topic[:] {
		function, ok := t.MethodByName(fn)
		if !ok {
			continue
		}
		quit, ok := t.MethodByName(quit)
		if ok {
			_ = bus.doSubscribe(topic[i], &EventHandler{
				t, observer, function.Func, quit.Func, flagOnce, async, transactional, new(sync.Mutex), topic[i], ext,
			})
		}
	}
}
func (bus *EventBus) registerProducer(observer dirverServer.Producer, topic string, async, transactional bool) {
	t := reflect.TypeOf(observer)
	function, ok := t.MethodByName("Send")
	if !ok {
		return
	}
	quit, ok := t.MethodByName("Close")
	if ok {
		_ = bus.doSubscribe(topic, &EventHandler{
			t, observer, function.Func, quit.Func, false, true, transactional, new(sync.Mutex), topic, "",
		})
	}
}
func (bus *EventBus) registerConsumer(dir dirverServer.ClientDirver, observer dirverServer.Consumer, topic string) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("eventbus, topic: %s catch err:%s", topic, err)
		}
	}()
	go observer.Consume()
	k := fmt.Sprintf("%v-%v-%v", dir, topic, observer.Options().GroupName)
	bus.comsumerPoolLock.Lock()
	bus.comsumerPool[k] = observer
	bus.comsumerPoolLock.Unlock()
	t := topic
	if len(observer.Options().GroupName) > 0 {
		t = t + "_" + observer.Options().GroupName
	}
	for msg := range observer.Msg() {
		bus.Publish(t, msg.Body, msg.Header)
	}
}

// Subscribe 订阅-同步
func (bus *EventBus) Subscribe(observer interface{}) {
	bus.subscribeWg.Add(1)
	go func() {
		defer bus.subscribeWg.Done()
		bus.register(observer, false, false, false)
	}()
}

// SubscribeAsync  订阅-异步
func (bus *EventBus) SubscribeAsync(observer interface{}, transactional bool) {
	bus.subscribeWg.Add(1)
	go func() {
		defer bus.subscribeWg.Done()
		bus.register(observer, false, true, transactional)
	}()

}

// SubscribeOnce 订阅-只执行一次-同步
func (bus *EventBus) SubscribeOnce(observer interface{}) {
	bus.subscribeWg.Add(1)
	go func() {
		defer bus.subscribeWg.Done()
		bus.register(observer, true, false, false)
	}()
}

// SubscribeOnceAsync 订阅-只执行一次-异步
func (bus *EventBus) SubscribeOnceAsync(observer interface{}) {
	bus.subscribeWg.Add(1)
	go func() {
		defer bus.subscribeWg.Done()
		bus.register(observer, true, true, false)
	}()
}

// HasCallback 查看事件订阅的函数
func (bus *EventBus) HasCallback(topic string) bool {
	handlersInterface, ok := bus.handlers.Load(topic)
	if ok {
		handlers := handlersInterface.([]*EventHandler)
		return len(handlers) > 0
	}
	return false
}

// Unsubscribe 删除订阅
func (bus *EventBus) Unsubscribe(observer interface{}) {
	topic, t, fn, _, _ := bus.checkObserver(observer)
	for i := range topic[:] {
		function, ok := t.MethodByName(fn)
		if !ok {
			continue
		}
		bus.removeHandler(topic[i], bus.findHandlerIdx(topic[i], function.Func))
	}

}

func (bus *EventBus) removeHandler(topic string, idx int) {
	handlerInterface, ok := bus.handlers.Load(topic)
	if !ok {
		return
	}
	handlers := handlerInterface.([]*EventHandler)
	l := len(handlers)

	if !(0 <= idx && idx < l) {
		return
	}
	handlers = append(handlers[:idx], handlers[idx+1:]...)
	if len(handlers) > 0 {
		bus.handlers.Store(topic, handlers)
	} else {
		bus.handlers.Delete(topic)
	}
}

// Publish 推送
// Deprecated: 废弃 使用PublishWithCtx
func (bus *EventBus) Publish(topic string, args ...interface{}) {
	if handlerInterface, ok := bus.handlers.Load(topic); ok {
		handlers := handlerInterface.([]*EventHandler)
		if len(handlers) == 0 {
			return
		}
		for i, handler := range handlers {
			if handler.flagOnce {
				bus.removeHandler(topic, i)
			}
			if !handler.async {
				bus.doPublish(handler, args...)
			} else {
				bus.wg.Add(1)
				if handler.transactional {
					handler.lock.Lock()
				}
				task := &task{
					f:            bus.doPublishAsync,
					eventHandler: handlers[i],
					args:         args,
				}
				if err := bus.pool.Submit(task); err != nil {
					log.Errorf("eventbus pool submit topic [%s] callBack [%s] err[%v] ", topic, task.eventHandler.callBack.String(), err)
				}
			}
		}
	}
}

// PublishWithCtx 推送
func (bus *EventBus) PublishWithCtx(ctx context.Context, topic string, args ...interface{}) {
	eventCtx := options.DetachCtx(ctx)
	params := make([]interface{}, 0, len(args)+1)
	params = append(append(params, eventCtx), args...)
	bus.Publish(topic, params...)
}

func (bus *EventBus) doPublish(handler *EventHandler, args ...interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("eventbus, topic:%s callBack: %s catch err:%s", handler.topic, handler.callBack.String(), err)
		}
	}()
	passedArguments := bus.setUpPublish(handler, args...)
	handler.callBack.Call(passedArguments)
}

func (bus *EventBus) doPublishAsync(handler *EventHandler, args ...interface{}) {
	defer bus.wg.Done()
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("eventbus, topic:%s callBack: %s catch err:%s", handler.topic, handler.callBack.String(), err)
		}
	}()
	if handler.transactional {
		defer handler.lock.Unlock()
	}

	bus.doPublish(handler, args...)

}

func (bus *EventBus) findHandlerIdx(topic string, callback reflect.Value) int {
	if handlerInterface, ok := bus.handlers.Load(topic); ok {
		handlers := handlerInterface.([]*EventHandler)
		for i := range handlers[:] {
			if handlers[i].callBack.Type() == callback.Type() &&
				handlers[i].callBack.Pointer() == callback.Pointer() {
				return i
			}
		}
	}
	return -1
}

func (bus *EventBus) setUpPublish(callback *EventHandler, args ...interface{}) []reflect.Value {
	funcType := callback.callBack.Type()
	passedArguments := make([]reflect.Value, 0, len(args)+1)
	passedArguments = append(passedArguments, reflect.ValueOf(callback.observer))
	for i := range args[:] {
		if args[i] == nil {
			passedArguments = append(passedArguments, reflect.New(funcType.In(i)).Elem())
		} else {
			passedArguments = append(passedArguments, reflect.ValueOf(args[i]))
		}
	}

	return passedArguments
}

// SetMaxTaskNum 设置最大执行任务数
func (bus *EventBus) SetMaxTaskNum(num int64) {
	bus.pool.Tune(int(num))
}

// WaitAsync 等待
func (bus *EventBus) WaitAsync() {
	bus.wg.Wait()
}

// Preview 预览任务
func (bus *EventBus) Preview() map[string]map[string]struct{} {
	var s strings.Builder
	res := make(map[string]map[string]struct{})
	bus.handlers.Range(func(topic, value interface{}) bool {
		handlers := value.([]*EventHandler)
		topics := topic.(string)
		s.WriteString(fmt.Sprintf("\n-------------------------\n%s:\n", topic))
		res[topics] = make(map[string]struct{})
		for i := range handlers[:] {
			res[topics][handlers[i].callBack.String()] = struct{}{}
			s.WriteString(fmt.Sprintf("%s\n", handlers[i].callBack.String()))
		}
		return true
	})
	fmt.Println(s.String())
	return res
}

// SubscribePublisherDirver 订阅发布驱动
func (bus *EventBus) SubscribePublisherDirver(dir dirverServer.ClientDirver, addrs []string, topic string, opts ...options.ProducerOption) {
	bus.subscribeWg.Add(1)
	go func() {
		defer bus.subscribeWg.Done()
		client := dirverServer.NewProducer(dir, addrs, topic, opts...)
		bus.registerProducer(client, topic, true, false)
	}()
}

// SubscribeConsumerDirver 订阅消费者驱动
func (bus *EventBus) SubscribeConsumerDirver(dir dirverServer.ClientDirver, addrs []string, topic string, opts ...options.ConsumerOption) {
	bus.subscribeWg.Add(1)
	go func() {
		defer bus.subscribeWg.Done()
		client := dirverServer.NewConsumer(dir, addrs, topic, opts...)
		go bus.registerConsumer(dir, client, topic)
	}()
}

// Quit 退出
func (bus *EventBus) Quit() {
	bus.handlers.Range(func(topic, value interface{}) bool {
		handlers := value.([]*EventHandler)
		for i := range handlers[:] {
			handlers[i].quitCallBack.Call([]reflect.Value{reflect.ValueOf(handlers[i].observer)})
			// _ = bus.Unsubscribe(handlers[i].observer)
		}

		return true
	})
	for _, v := range bus.comsumerPool {
		v.Close()
	}
}

// SubscribeRegisterWait 订阅注册等待
func (bus *EventBus) SubscribeRegisterWait() {
	bus.subscribeWg.Wait()
}

// AllTask 所有任务
func (bus *EventBus) AllTask() map[string][]*EventHandler {
	event := make(map[string][]*EventHandler)
	bus.handlers.Range(func(topic, value interface{}) bool {
		handlers := value.([]*EventHandler)
		topics := topic.(string)
		event[topics] = append(event[topics], handlers...)
		return true
	})
	return event
}

// Ack 确认消息
func (bus *EventBus) Ack(dirver dirverServer.ClientDirver, topic, groupName, msgID string) {
	k := fmt.Sprintf("%v-%v-%v", dirver, topic, groupName)
	bus.comsumerPoolLock.Lock()
	comsumer, ok := bus.comsumerPool[k]
	bus.comsumerPoolLock.Unlock()
	if ok {
		comsumer.AckID(msgID)
	}
}
