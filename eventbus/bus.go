package eventbus

import (
	"errors"
	"sync"
)

var ErrNotExist = errors.New("topic not exist")

type bus struct {
	rw    sync.RWMutex
	nodes map[string]*node
}

func NewBus() *bus {
	return &bus{
		nodes: make(map[string]*node),
	}
}

func (b *bus) Subscribe(topic string, sub Sub) {
	b.rw.Lock()
	if n, ok := b.nodes[topic]; ok {
		b.rw.Unlock()
		n.subscribe(sub)
	} else {
		defer b.rw.Unlock()
		b.nodes[topic] = newNode(sub)
	}
}

func (b *bus) Publish(topic string, msg interface{}) error {
	b.rw.RLock()
	defer b.rw.RUnlock()
	if n, ok := b.nodes[topic]; ok {
		n.Publish(msg)
	} else {
		return ErrNotExist
	}
	return nil
}
