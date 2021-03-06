package parse

import (
	"crawl/LearnGo-crawl/engine"
	"regexp"
)

const BooklistRe = `<a.*?href="([^"]+)" title="([^"]+)"`

func ParseBookList(contents []byte) engine.ParseResult {
	//fmt.Printf("%s", contents)
	re := regexp.MustCompile(BooklistRe)

	matchers := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matchers {
		bookname := string(m[2])
		//result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParseFunc: func(c []byte) engine.ParseResult {
				return ParseBookDetail(c, bookname)
			},
		})
	}
	return result
}
