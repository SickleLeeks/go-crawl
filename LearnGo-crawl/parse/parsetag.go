package parse

import (
	"crawl/LearnGo-crawl/engine"
	"regexp"
)

const regexpStr = `<a href="([^"]+)" class="tag">([^<]+)</a>`

func ParseTag(content []byte) engine.ParseResult {
	// <a href="/tag/励志" class="tag">励志</a>
	re := regexp.MustCompile(regexpStr)
	match := re.FindAllSubmatch(content, -1)

	result := engine.ParseResult{}
	for _, m := range match {

		result.Items = append(result.Items, m[2])
		result.Requests = append(result.Requests, engine.Request{
			Url:       "https://book.douban.com" + string(m[1]),
			ParseFunc: ParseBookList,
		})
	}

	return result
}
