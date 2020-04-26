package main

import (
	"github.com/corrots/go-demo/gocrawler/engine"
	"github.com/corrots/go-demo/gocrawler/scheduler"
	"github.com/corrots/go-demo/gocrawler/zhenai/parser"
)

const cityListURL = "http://www.zhenai.com/zhenghun"

func main() {
	//e := &engine.SimpleEngine{}
	//e := &engine.ConcurrentEngine{
	//	Scheduler:   &scheduler.SimpleScheduler{},
	//	WorkerCount: 10,
	//}
	e := &engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
	}
	e.Run(engine.Request{
		URL:        cityListURL,
		ParserFunc: parser.ParseCityList,
	})
}
