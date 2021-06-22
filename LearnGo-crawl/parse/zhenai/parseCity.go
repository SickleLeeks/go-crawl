package zhenai

import (
	"crawl/LearnGo-crawl/engine"
	"regexp"
)

var cityRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[\d]+)" target="_blank">([^<]+)</a>`)
var cityUrlRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[^"]+)"`)

func ParseCity(contents []byte) engine.ParseResult {
	//matches := cityRe.FindAllStringSubmatch(contents, -1)
	matches := cityRe.FindAllSubmatch(contents,-1)
	result := engine.ParseResult{}
	for _, m := range matches {
		url := string(m[1])
		name := string(m[2])
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParseFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(
					c, url, name)
			},
		})
	}
	// 查找城市页面下的城市连接
	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(m[1]),
			ParseFunc: ParseCity,
		})
	}
	return result
}
