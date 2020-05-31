package parser

import (
	"regexp"

	"github.com/crawler/engine"
)

const blockRe = `href="(https://sh.lianjia.com/ershoufang/[0-9]+.html)"[^>]*>([^<]+)</a>`

// <a href="/ershoufang/anshan/" >鞍山</a>

func ParseBlock(contents []byte) engine.ParseResult {
	compile := regexp.MustCompile(blockRe)
	matches := compile.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		title := string(m[2])
		//result.Items = append(result.Items, "title: "+title)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseHouseProfile(c, title)
			},
		})
	}
	return result
}
