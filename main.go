package main

import (
	"github.com/crawler/engine"
	"github.com/crawler/lianjia/parser"
	"github.com/crawler/persist"
	"github.com/crawler/scheduler"
)

func main() {
	// contents, err := ioutil.ReadFile("shanghailianjia.html")
	// if err != nil {
	// 	panic(err)
	// }
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    persist.ItemSaver(),
	}
	e.Run(engine.Request{
		Url:        "https://sh.lianjia.com/ershoufang/",
		ParserFunc: parser.ParseAreaList,
	})
}
