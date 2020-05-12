package main

import (
	"github.com/crawler/engine"
	"github.com/crawler/lianjia/parser"
	"github.com/crawler/scheduler"
)

func main() {
	// contents, err := ioutil.ReadFile("shanghailianjia.html")
	// if err != nil {
	// 	panic(err)
	// }
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
	}
	e.Run(engine.Request{
		Url:        "https://sh.lianjia.com/ershoufang/",
		ParserFunc: parser.ParseAreaList,
	})
}
