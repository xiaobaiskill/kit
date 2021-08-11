package eventbus

var globalBus *bus

func init() {
	globalBus = NewBus()
}

func Subscribe(topic string, sub Sub) {
	globalBus.Subscribe(topic, sub)
}

func SubscribeFunc(topic string, f func(msg interface{})) {
	globalBus.Subscribe(topic, ReceiveFunc(f))
}

func Publish(topic string, msg interface{}) error {
	return globalBus.Publish(topic, msg)
}
