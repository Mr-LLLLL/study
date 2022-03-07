package chain_of_responsibility

import "chain-of-responsibility/base"

type Widget struct {
	_parent    *Widget
	_successor IHelpHandler
	_topic     base.Topic
}

func (b *Widget) HasHelp() bool {
	return b._topic != base.NO_HELP_TOPIC
}

func (b *Widget) HandleHelp() {
	if b._successor != nil {
		b._successor.HandleHelp()
	}
}

func NewWidget(w *Widget, t base.Topic) *Widget {
	return &Widget{
		_parent:    w,
		_successor: w,
		_topic:     t,
	}
}
