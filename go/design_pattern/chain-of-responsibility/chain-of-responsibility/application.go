package chain_of_responsibility

import "chain-of-responsibility/base"

type Application struct {
	_successor IHelpHandler
	_topic     base.Topic
}

func (d *Application) HasHelp() bool {
	return d._topic != base.NO_HELP_TOPIC
}

func (d *Application) HandleHelp() {

}

func NewApplication(t base.Topic) *Application {
	return &Application{
		_successor: nil,
		_topic:     t,
	}
}
