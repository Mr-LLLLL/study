package observer

import "testing"

func TestSubject_Notify(t *testing.T) {
	s := &Subject{}
	s.Register(&Observer1{})
	s.Register(&Observer2{})
	s.Notify("hi")
}
