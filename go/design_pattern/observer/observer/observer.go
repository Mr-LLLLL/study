package observer

import "reflect"

type ISubject interface {
	Attach(IObserver)
	Detach(IObserver)
	Notify()
}

type Subject struct {
	_observers []IObserver
}

func (s *Subject) Attach(o IObserver) {
	s._observers = append(s._observers, o)
}

func (s *Subject) Detach(o IObserver) {
	for i := 0; i < len(s._observers); i++ {
		if reflect.DeepEqual(s._observers[i], o) {
			s._observers = append(s._observers[:i], s._observers[i+1:]...)
			i = i - 1
		}
	}
}

func (s *Subject) Notify() {
	for _, v := range s._observers {
		v.Update(s)
	}
}

type ClockTimer struct {
	_observers []IObserver
}

func (c *ClockTimer) Attach(o IObserver) {
	c._observers = append(c._observers, o)
}

func (c *ClockTimer) Detach(o IObserver) {
	for i := 0; i < len(c._observers); i++ {
		if reflect.DeepEqual(c._observers[i], o) {
			c._observers = append(c._observers[:i], c._observers[i+1:]...)
			i = i - 1
		}
	}
}

func (c *ClockTimer) Notify() {
	for _, v := range c._observers {
		v.Update(c)
	}
}

func (c *ClockTimer) Tick() {
	c.Notify()
}

type IObserver interface {
	Update(theChangedSubject ISubject)
}

type DigitalClock struct {
	_subject *ClockTimer
}

func (d *DigitalClock) Update(theChangedSubject ISubject) {
	if reflect.DeepEqual(theChangedSubject, d._subject) {
		d.Draw()
	}
}

func (d *DigitalClock) Draw() {}

func NewDigitalClock(c *ClockTimer) *DigitalClock {
	d := &DigitalClock{
		_subject: c,
	}
	c.Attach(d)
	return d
}
