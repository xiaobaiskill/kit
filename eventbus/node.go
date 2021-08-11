package eventbus

import "sync"

type node struct {
	rw   sync.RWMutex
	subs []Sub
}

func newNode(subs ...Sub) *node {
	return &node{
		subs: subs,
	}
}

func (n *node) subsLen() int {
	return len(n.subs)
}

func (n *node) subscribe(sub Sub) {
	n.rw.Lock()
	n.subs = append(n.subs, sub)
	n.rw.Unlock()
}

func (n *node) clean() {
	n.rw.Lock()
	n.subs = []Sub{}
	n.rw.Unlock()
}

func (n *node) Publish(msg interface{}) {
	n.rw.RLock()
	go func(subs []Sub, msg interface{}) {
		for _, s := range subs {
			s.Receive(msg)
		}
	}(n.subs, msg)
	n.rw.RUnlock()
}
