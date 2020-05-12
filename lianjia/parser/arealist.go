package parser

import (
	"regexp"

	"github.com/crawler/engine"
)

// const areaListRe = `<a href="(/ershoufang/[a-z]+/)"[^>]*>([^<]+)</a>`
const areaListRe = `<a href="(/ershoufang/yangpu/)"[^>]*>([^<]+)</a>`

func ParseAreaList(contents []byte) engine.ParseResult {
	compile := regexp.MustCompile(areaListRe)
	matches := compile.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "Area "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        "https://sh.lianjia.com" + string(m[1]),
			ParserFunc: ParseArea,
		})
	}
	return result
}
