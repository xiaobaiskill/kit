package workpool

var (
	workerPool workerPoolType
	works      = make([]*worker, 0)
)

type workerPoolType chan chan Job

func StartDispathcher(work_num int) chan Job {
	workQueue := make(chan Job, work_num*2)
	workerPool = make(workerPoolType, work_num)

	for i := 0; i < work_num; i++ {
		newWorker(i, workerPool).start()
	}

	go func() {
		for work := range workQueue {
			workqueue := <-workerPool
			workqueue <- work
		}
	}()
	return workQueue
}

func StopDispathcher() {
	for _, work := range works {
		work.stop()
	}
}
