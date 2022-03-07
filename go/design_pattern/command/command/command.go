package command

import "command/base"

type ICommand interface {
	Execute()
}

type OpenCommand struct {
	_response    string
	_application *base.Application
}

func (o *OpenCommand) AskUser() string {
	panic("implemented this")
}

func (o *OpenCommand) Execute() {
	name := o.AskUser()

	if name != "" {
		document := base.NewDocument()
		o._application.Add(document)
		document.Open()
	}
}

func NewOpenCommand(a *base.Application) *OpenCommand {
	return &OpenCommand{
		_application: a,
	}
}

type PasteCommand struct {
	_document *base.Document
}

func (p *PasteCommand) Execute() {
	p._document.Paste()
}

func NewPasteCommand(d *base.Document) *PasteCommand {
	return &PasteCommand{
		_document: d,
	}
}

type MacroCommand struct {
	_cmds []ICommand
}

func (m *MacroCommand) Execute() {
	for _, v := range m._cmds {
		v.Execute()
	}
}

func (m *MacroCommand) Add(c ICommand) {
	m._cmds = append(m._cmds, c)
}

func (m *MacroCommand) Remove(c ICommand) {
	for i, v := range m._cmds {
		if v == c {
			m._cmds = append(m._cmds[:i], m._cmds[i+1:]...)
		}
	}
}
