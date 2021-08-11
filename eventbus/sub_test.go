package eventbus

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testSub struct {
	ch chan string
}

func newTestSub() *testSub {
	return &testSub{ch: make(chan string)}
}

func (t *testSub) Receive(msg interface{}) {
	t.ch <- msg.(string)
}

var testData = "aaaa"

func TestSub(t *testing.T) {
	ts := newTestSub()
	go ts.Receive(testData)
	assert.Equal(t, <-ts.ch, testData)
}
