package zhenai

import (
	"crawl/LearnGo-crawl/engine"
	"regexp"
)

const cityListRe = `(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents string) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllStringSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		//result.Items = append(result.Items,string(m[2]))
		result.Requests = append(
			result.Requests, engine.Request{
				Url:       m[1],
				ParseFunc: ParseCity,
			})
	}
	return result
}
