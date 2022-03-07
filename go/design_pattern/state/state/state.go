package state

type TcpConnection struct {
	_state ITcpState
}

func (t *TcpConnection) ActiveOepn() {
	t._state.ActiveOepn(t)
}

func (t *TcpConnection) PassiveOpen() {
	t._state.PassiveOpen(t)
}

func (t *TcpConnection) Close() {
	t._state.Close(t)
}

func (t *TcpConnection) Send() {
	t._state.Send(t)
}

func (t *TcpConnection) Acknowledge() {
	t._state.Acknowledge(t)
}

func (t *TcpConnection) Synchronize() {
	t._state.Synchronize(t)
}

func (t *TcpConnection) ProcessOctet(*TcpOctetStream) {}

func (t *TcpConnection) ChangeState(state ITcpState) {
	t._state = state
}

func NewTcpConnection() *TcpConnection {
	return &TcpConnection{
		// _state: ,
	}
}

type TcpOctetStream struct{}

type ITcpState interface {
	Transmit(*TcpConnection, *TcpOctetStream)
	ActiveOepn(*TcpConnection)
	PassiveOpen(*TcpConnection)
	Close(*TcpConnection)
	Send(*TcpConnection)
	Acknowledge(*TcpConnection)
	Synchronize(*TcpConnection)
	ChangeState(*TcpConnection, ITcpState)
}

type TcpEstablished struct{}

func (t *TcpEstablished) Transmit(_ *TcpConnection, _ *TcpOctetStream) {
	panic("not implemented") // TODO: Implement
}

func (t *TcpEstablished) ActiveOepn(_ *TcpConnection) {
	panic("not implemented") // TODO: Implement
}

func (t *TcpEstablished) PassiveOpen(_ *TcpConnection) {
	panic("not implemented") // TODO: Implement
}

func (t *TcpEstablished) Close(_ *TcpConnection) {
	panic("not implemented") // TODO: Implement
}

func (t *TcpEstablished) Send(_ *TcpConnection) {
	panic("not implemented") // TODO: Implement
}

func (t *TcpEstablished) Acknowledge(_ *TcpConnection) {
	panic("not implemented") // TODO: Implement
}

func (t *TcpEstablished) Synchronize(_ *TcpConnection) {
	panic("not implemented") // TODO: Implement
}

func (t *TcpEstablished) ChangeState(_ *TcpConnection, _ ITcpState) {
	panic("not implemented") // TODO: Implement
}

var TcpEstablishedSingle *TcpEstablished

func NewTcpEstablished() *TcpEstablished {
	if TcpEstablishedSingle == nil {
		TcpEstablishedSingle = new(TcpEstablished)
	}
	return TcpEstablishedSingle
}

type TcpListen struct{}

func (t *TcpListen) Transmit(_ *TcpConnection, _ *TcpOctetStream) {
	panic("not implemented") // TODO: Implement
}

func (t *TcpListen) ActiveOepn(_ *TcpConnection) {
	panic("not implemented") // TODO: Implement
}

func (t *TcpListen) PassiveOpen(_ *TcpConnection) {
	panic("not implemented") // TODO: Implement
}

func (t *TcpListen) Close(_ *TcpConnection) {
	panic("not implemented") // TODO: Implement
}

func (t *TcpListen) Send(_ *TcpConnection) {
	panic("not implemented") // TODO: Implement
}

func (t *TcpListen) Acknowledge(_ *TcpConnection) {
	panic("not implemented") // TODO: Implement
}

func (t *TcpListen) Synchronize(_ *TcpConnection) {
	panic("not implemented") // TODO: Implement
}

func (t *TcpListen) ChangeState(_ *TcpConnection, _ ITcpState) {
	panic("not implemented") // TODO: Implement
}

var TcpListenSingle *TcpListen

func NewTcpListen() *TcpListen {
	if TcpListenSingle == nil {
		TcpListenSingle = new(TcpListen)
	}
	return TcpListenSingle
}

type TcpClosed struct{}

func (t *TcpClosed) Transmit(c *TcpConnection, o *TcpOctetStream) {
	c.ProcessOctet(o)
}

func (t *TcpClosed) ActiveOepn(c *TcpConnection) {
	t.ChangeState(c, NewTcpEstablished())
}

func (t *TcpClosed) PassiveOpen(c *TcpConnection) {
	t.ChangeState(c, NewTcpListen())
}

func (t *TcpClosed) Close(c *TcpConnection) {
	t.ChangeState(c, NewTcpListen())
}

func (t *TcpClosed) Send(c *TcpConnection) {
	t.ChangeState(c, NewTcpEstablished())
}

func (t *TcpClosed) Acknowledge(_ *TcpConnection) {
	panic("not implemented") // TODO: Implement
}

func (t *TcpClosed) Synchronize(_ *TcpConnection) {
	panic("not implemented") // TODO: Implement
}

func (t *TcpClosed) ChangeState(_ *TcpConnection, _ ITcpState) {
	panic("not implemented") // TODO: Implement
}

var TcpClosedSingle *TcpClosed

func NewTcpClosed() *TcpClosed {
	if TcpClosedSingle == nil {
		TcpClosedSingle = new(TcpClosed)
	}
	return TcpClosedSingle
}

