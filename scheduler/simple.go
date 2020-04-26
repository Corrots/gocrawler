package scheduler

import "github.com/corrots/go-demo/gocrawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan = w
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) Register(r engine.Request) {
	go func() { s.workerChan <- r }()
}
