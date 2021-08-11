package eventbus

type Sub interface {
	Receive(interface{})
}

type ReceiveFunc func(msg interface{})

func (f ReceiveFunc) Receive(msg interface{}) {
	f(msg)
}
