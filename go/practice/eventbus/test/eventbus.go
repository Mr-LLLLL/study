package main

import (
	"context"
	"fmt"
	"time"

	"git.dustess.com/mk-base/eventbus"
	"git.dustess.com/mk-base/log"
)

var EventBusHandler eventbus.Bus

// InitEventBus 初始化
func InitEventBus() {
	log.Infof("eventbus start.")
	EventBusHandler = eventbus.New()
	EventBusHandler.SetMaxTaskNum(3)

}

// Wait Wait
func Wait() {
	EventBusHandler.SubscribeRegisterWait()
	EventBusHandler.Preview()
}

// Quit Quit
func Quit() {
	log.Infof("eventbus exit.")
	EventBusHandler.Quit()
}

func InitEvent() {
	EventBusHandler.Subscribe(NewEvent())
}

type Event struct {
	event interface{} `subscribe:"Handle" topic:"test" quit:"Quit"`
}

// NewEvent
func NewEvent() *Event {
	return new(Event)
}

// PushOne
func (e *Event) Handle(ctx context.Context, str string) {
	fmt.Println(str)
	time.Sleep(time.Second)
}

// Quit
func (e *Event) Quit() {

}

func Publish(str string) {
	EventBusHandler.PublishWithCtx(context.Background(), "test", str)
	fmt.Println("push", str)
}
