package engine

import "fmt"

type Scheduler interface {
	Register(Request)
	ConfigureMasterWorkerChan(chan Request)
}

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	for _, r := range seeds {
		e.Scheduler.Register(r)
	}

	in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.ConfigureMasterWorkerChan(in)
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}
	// 从out_chan接收result
	for {
		result := <-out
		for _, item := range result.Items {
			fmt.Printf("Got item: %+v\n", item)
		}
		// result中的Request继续加入Request chan
		for _, r := range result.Requests {
			e.Scheduler.Register(r)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			r := <-in
			result, err := worker(r)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
