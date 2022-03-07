package mediator

import "testing"

func TestDemo(t *testing.T) {
	usernameInput := Input("username")
	passwordInput := Input("passwordInput")
	repeatPasswordInput := Input("repeatPasswordInput")

	selection := Selection("login")
	d := &Dialog{
		Selection:           &selection,
		UsernameInput:       &usernameInput,
		PasswordInput:       &passwordInput,
		RepeatPasswordInput: &repeatPasswordInput,
	}
	d.HandleEvent(&selection)

	regSelection := Selection("register")
	d.Selection = &regSelection
	d.HandleEvent(&regSelection)
}
