package chain_of_responsibility

import "chain-of-responsibility/base"

type Dialog struct {
	_parent    *Widget
	_successor IHelpHandler
	_topic     base.Topic
}

func (d *Dialog) HasHelp() bool {
	return d._topic != base.NO_HELP_TOPIC
}

func (d *Dialog) HandleHelp() {
	if d.HasHelp() {
	} else {
		NewHelpHandler(d._successor, d._topic).HandleHelp()
	}
}

func NewDialog(h IHelpHandler, t base.Topic) *Dialog {
	return &Dialog{
		_parent:    nil,
		_successor: h,
		_topic:     t,
	}
}
