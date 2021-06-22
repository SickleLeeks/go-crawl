package main

import (
	"bufio"
	"crawl/LearnGo-crawl/engine"
	"crawl/LearnGo-crawl/parse"
	"crawl/LearnGo-crawl/persist"
	"crawl/LearnGo-crawl/scheduler"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"log"
	"regexp"
)

func main() {
	//engine.Run(engine.Request{
	//	//Url:       "https://book.douban.com/tag/%E7%BC%96%E7%A8%8B",
	//	//ParseFunc: parse.ParseBookList,
	//	Url:       "https://book.douban.com",
	//	ParseFunc: parse.ParseTag,
	//	//Url:       "https://book.douban.com/subject/4822685/",
	//	//ParseFunc: parse.ParseBookDetail,
	//})
	e := engine.ConcurrentEngine{
		&scheduler.QueueScheduler{},
		100,
		persist.ItemSave(),
	}
	e.Run(engine.Request{
		Url:       "https://book.douban.com",
		ParseFunc: parse.ParseTag,
		//Url: "http://www.zhenai.com/zhenghun",
		//ParseFunc: zhenai.ParseCity,
	})
}
func DeterminEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("fetch error : %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
func ParseContent(content []byte) {
	// <a href="/tag/励志" class="tag">励志</a>
	re := regexp.MustCompile(`<a href="([^"]+)" class="tag">([^"]+)</a>`)
	match := re.FindAllSubmatch(content, -1)
	for _, m := range match {
		//fmt.Printf("m[0]:%s m[1]:%s m[2]:%s\n", m[0], m[1], m[2])
		fmt.Printf("url:%s\n", "https://book.douban.com"+string(m[1]))
	}
}
