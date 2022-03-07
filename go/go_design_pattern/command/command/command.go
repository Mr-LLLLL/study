package command

import "fmt"

type ICommand interface {
	Execute() error
}

type StartCommand struct{}

func (s *StartCommand) Execute() error {
	fmt.Println("game start")
	return nil
}

func NewStartCommand() *StartCommand {
	return new(StartCommand)
}

type ArchiveCommand struct{}

func NewArchiveCommand() *ArchiveCommand {
	return new(ArchiveCommand)
}

func (a *ArchiveCommand) Execute() error {
	fmt.Println("game archive")
	return nil
}
