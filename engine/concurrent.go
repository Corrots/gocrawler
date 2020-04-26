package engine

import "fmt"

type Scheduler interface {
	ReadyNotifier
	Register(Request)
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	// 执行创建Request chan和Worker chan等操作
	e.Scheduler.Run()
	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Register(r)
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

func (e *ConcurrentEngine) createWorker(out chan ParseResult, notifier ReadyNotifier) {
	in := make(chan Request)
	go func() {
		for {
			notifier.WorkerReady(in)
			r := <-in
			result, err := worker(r)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
