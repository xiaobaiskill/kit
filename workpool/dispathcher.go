package workpool

type workerPoolType chan chan Job

type dispathcher struct {
	workerPool workerPoolType
	works      []*worker
	workQueue  chan Job
}

func NewDispathcher(num int) *dispathcher {
	workQueue := make(chan Job, num*2)
	workerPool := make(workerPoolType, num)
	works := make([]*worker, 0, num)
	for i := 0; i < num; i++ {
		work := newWorker(i, workerPool)
		work.start()
		works = append(works, work)
	}

	go func() {
		for work := range workQueue {
			workqueue := <-workerPool
			workqueue <- work
		}
	}()
	return &dispathcher{
		workerPool: workerPool,
		works:      works,
		workQueue:  workQueue,
	}
}

func (d *dispathcher) AddJob(job Job) {
	d.workQueue <- job
}

func (d *dispathcher) Stop() {
	for _, work := range d.works {
		work.stop()
	}
}
