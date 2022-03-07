package mediator

import (
	"fmt"
	"reflect"
)

type Input string

func (i Input) String() string {
	return string(i)
}

type Selection string

func (s Selection) Selected() string {
	return string(s)
}

type Button struct {
	onClick func()
}

func (b *Button) SetOnClick(f func()) {
	b.onClick = f
}

type IMediator interface {
	HandleEvent(component interface{})
}

type Dialog struct {
	LoginButton         *Button
	RegButton           *Button
	Selection           *Selection
	UsernameInput       *Input
	PasswordInput       *Input
	RepeatPasswordInput *Input
}

func (d *Dialog) HandleEvent(component interface{}) {
	switch {
	case reflect.DeepEqual(component, d.Selection):
		if d.Selection.Selected() == "login" {
			fmt.Println("select login")
			fmt.Printf("show: %s\n", d.UsernameInput)
			fmt.Printf("show: %s\n", d.PasswordInput)
		} else if d.Selection.Selected() == "register" {
			fmt.Println("select register")
			fmt.Printf("show: %s\n", d.UsernameInput)
			fmt.Printf("show: %s\n", d.PasswordInput)
			fmt.Printf("show: %s\n", d.RepeatPasswordInput)
		}
	}
}
