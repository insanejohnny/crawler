package parser

import (
	"regexp"

	"github.com/crawler/engine"
)

// <a href="/ershoufang/anshan/" >鞍山</a>
// const areaRe = `<a href="(/ershoufang/[a-z]+/)" >([^<]+)</a>`
const areaRe = `<a href="(/ershoufang/[a-z]+/)" >(鞍山)</a>`

func ParseArea(contents []byte) engine.ParseResult {
	compile := regexp.MustCompile(areaRe)
	matches := compile.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "Block "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        "https://sh.lianjia.com" + string(m[1]),
			ParserFunc: ParseBlock,
		})
	}
	return result
}
