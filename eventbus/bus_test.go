package eventbus

import (
	"testing"
	"time"
)

var testbus = NewBus()

func TestBus(t *testing.T) {
	s1 := newTestSub()
	s2 := newTestSub()
	s3 := newTestSub()

	go func() {
		for {
			select {
			case d := <-s1.ch:
				t.Log("s1:", d)
			case d := <-s2.ch:
				t.Log("s2:", d)
			case d := <-s3.ch:
				t.Log("s3:", d)
			}
		}
	}()

	testbus.Subscribe("top1", s1)
	testbus.Subscribe("top2", s2)
	testbus.Subscribe("top1", s3)

	go testPublish("top1", "data1")
	go testPublish("top2", "data2")

	time.Sleep(time.Second)
}

func testPublish(topic string, msg string) {
	for {
		testbus.Publish(topic, msg)
		time.Sleep(time.Millisecond * 100)
	}
}
