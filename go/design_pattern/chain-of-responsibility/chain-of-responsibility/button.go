package chain_of_responsibility

import "chain-of-responsibility/base"

type Button struct {
	_parent    *Widget
	_successor IHelpHandler
	_topic     base.Topic
}

func (b *Button) HasHelp() bool {
	return b._topic != base.NO_HELP_TOPIC
}

func (b *Button) HandleHelp() {
	if b.HasHelp() {
	} else {
		NewHelpHandler(b._successor, b._topic).HandleHelp()
	}
}

func NewButton(w *Widget, t base.Topic) *Button {
	return &Button{
		_parent:    w,
		_successor: w,
		_topic:     t,
	}
}
