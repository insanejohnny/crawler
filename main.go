package main

import (
	"github.com/crawler/engine"
	"github.com/crawler/lianjia/parser"
)

func main() {
	// contents, err := ioutil.ReadFile("shanghailianjia.html")
	// if err != nil {
	// 	panic(err)
	// }
	engine.Run(engine.Request{
		Url:        "https://sh.lianjia.com/ershoufang/",
		ParserFunc: parser.ParseAreaList,
	})
}
