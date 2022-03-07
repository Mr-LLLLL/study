package observer

import "fmt"

type ISubject interface {
	Register(observer IObserver)
	Remove(observer IObserver)
	Notify(string)
}

type Subject struct {
	observers []IObserver
}

func (s *Subject) Register(observer IObserver) {
	s.observers = append(s.observers, observer)
}

func (s *Subject) Remove(observer IObserver) {
	for i := 0; i < len(s.observers); i++ {
		if s.observers[i] == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			i = i - 1
		}
	}
}

func (s *Subject) Notify(msg string) {
	for _, o := range s.observers {
		o.Update(msg)
	}
}

type IObserver interface {
	Update(msg string)
}

type Observer1 struct{}

func (Observer1) Update(msg string) {
	fmt.Printf("Observer1: %s", msg)
}

type Observer2 struct{}

func (Observer2) Update(msg string) {
	fmt.Printf("Observer2: %s", msg)
}
