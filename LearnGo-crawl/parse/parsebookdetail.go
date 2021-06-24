package parse

import (
	"crawl/LearnGo-crawl/engine"
	"crawl/LearnGo-crawl/model"
	"regexp"
	"strconv"
	"strings"
)

var authorRe = regexp.MustCompile(`<span class="pl">[\s]*?作者[:]?</span>[\d\D]*?<a.*?>([^<]+)</a>`)
var publicerRe = regexp.MustCompile(`<span class="pl">[\s]*?出版社[:]?</span>[\s]*?([^<]+)`)
var pageRe = regexp.MustCompile(`<span class="pl">页数:</span>[\s]*?([^<]+)`)
var priceRe = regexp.MustCompile(`<span class="pl">[\s]*?定价:</span>[\s]*?([^<]+)`)
var scoreRe = regexp.MustCompile(`<strong class="ll rating_num " property="v:average">([^<]+)</strong>`)
var inforRe = regexp.MustCompile(`<div class="intro">[\d\D]*?<p>([^<]+)</p></div>`)

func ParseBookDetail(contents []byte, bookname string) engine.ParseResult {
	//fmt.Printf("%s",contents)
	bookdetail := model.Bookdetail{}
	bookdetail.Author = strings.Replace(strings.Replace(ExtraString(contents, authorRe), "\n", "", -1), " ", "", -1)
	page, err := strconv.Atoi(strings.Replace(ExtraString(contents, pageRe), " ", "", -1))
	if err == nil {
		bookdetail.Bookpages = page
	}
	bookdetail.Bookname = bookname
	bookdetail.Publicer = ExtraString(contents, publicerRe)
	bookdetail.Price = ExtraString(contents, priceRe)
	bookdetail.Score = ExtraString(contents, scoreRe)
	bookdetail.Intro = ExtraString(contents, inforRe)
	result := engine.ParseResult{
		Items: []engine.Item{},
	}
	return result

}
func ExtraString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
