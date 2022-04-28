package workpool

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type Student struct {
	wg *sync.WaitGroup
	id int
}

func (s *Student) Execute() {
	time.Sleep(time.Second * 1)
	fmt.Printf("%d doing...\n", s.id)
	s.wg.Done()
}

func TestStartDispathcher(t *testing.T) {
	start := time.Now()
	jobChan := StartDispathcher(2)
	var wg sync.WaitGroup
	for i := 1; i < 10; i++ {
		s := &Student{
			id: i,
			wg: &wg,
		}
		wg.Add(1)
		jobChan <- s
	}

	wg.Wait()
	StopDispathcher()

	sub := time.Since(start).Seconds()
	t.Log(sub)
}
