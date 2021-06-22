package zhenai

import (
	"crawl/LearnGo-crawl/engine"
	"crawl/LearnGo-crawl/model"
	"regexp"
	"strconv"
)

var marryRe = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">(已婚|未婚|离异|丧偶)</div>`)
var ageRe = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">([\d]+)岁</div>`)
var constellationRe = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">([\S]+座\([\d]{2}.[\d]{2}-[\d]{2}.[\d]{2}\))</div>`)
var heightRe = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">([\d]+)cm</div>`)
var weightRe = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">([\d]+)kg</div>`)
var workareaRe = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">工作地:([^<]+)</div>`)
var salaryRe = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">月收入:([^<]+)</div>`)

var idRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

//解析器 解析用户
// name为上一级传递过来的
func ParseProfile(contents []byte, url string, name string) engine.ParseResult {
	//fmt.Println("url:"+url)
	//用户结构体
	profile := model.Profile{}
	profile.Name = name
	//年龄 string转换为int
	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}
	//婚姻状况
	profile.Marry = extractString(contents, marryRe)
	//身高
	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}
	//体重
	weight, err := strconv.Atoi(extractString(contents,weightRe))
	if err == nil {
		profile.Weight = weight
	}
	//工作地
	profile.WorkArea = extractString(contents, workareaRe)
	//薪水
	profile.Salary = extractString(contents, salaryRe)
	//星座
	profile.Constellation = extractString(contents, constellationRe)
	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}
func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindAllSubmatch(contents,-1)
	if len(match) >= 2 {
		return string(match[0][1])
	} else {
		return ""
	}
}
