package chain_of_responsibility

import "chain-of-responsibility/base"

type IHelpHandler interface {
	HasHelp() bool
	HandleHelp()
}

type HelpHandler struct {
	_successor IHelpHandler
	_topic     base.Topic
}

func (h *HelpHandler) HasHelp() bool {
	return h._topic != base.NO_HELP_TOPIC
}

func (h *HelpHandler) HandleHelp() {
	if h._successor != nil {
		h._successor.HandleHelp()
	}
}

func NewHelpHandler(h IHelpHandler, t base.Topic) *HelpHandler {
	return &HelpHandler{
		_successor: h,
		_topic:     t,
	}
}
