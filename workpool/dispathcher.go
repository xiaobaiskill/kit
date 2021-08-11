package workpool

var (
	WorkQueue  chan Job
	workerPool workerPoolType
	works      = make([]*worker, 0)
)

type workerPoolType chan chan Job

func StartDispathcher(work_num int) {
	WorkQueue = make(chan Job, work_num*2)
	workerPool = make(workerPoolType, work_num)

	for i := 0; i < work_num; i++ {
		newWorker(i, workerPool).start()
	}

	go func() {
		for {
			select {
			case work := <-WorkQueue:
				workqueue := <-workerPool
				workqueue <- work
			}
		}
	}()
}

func StopDispathcher() {
	for _, work := range works {
		work.stop()
	}
}
