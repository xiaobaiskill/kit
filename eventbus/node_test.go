package eventbus

import (
	"fmt"
	"testing"
	"time"
)

func TestNode(t *testing.T) {
	s1 := newTestSub()
	s2 := newTestSub()
	s3 := newTestSub()

	n := newNode(s1, s2)
	n.subscribe(s3)

	var close = make(chan struct{})
	go func() {
		for i := 0; i < 10; i++ {
			n.Publish(fmt.Sprintf("data-%d", i))
		}
		close <- struct{}{}
	}()

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
	<-close
	time.Sleep(time.Millisecond * 100)
}
