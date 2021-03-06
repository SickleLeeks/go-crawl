package model

import "strconv"

type Bookdetail struct {
	Bookname  string
	Author    string
	Publicer  string
	Bookpages int
	Price     string
	Score     string
	Intro     string
}

func (b Bookdetail) String() string {
	return "书籍名称：" + b.Bookname + " 作者 :" + b.Author + " 出版社:" + b.Publicer + " 书籍页数：" + strconv.Itoa(b.Bookpages) + " 价格：" + b.Price + " 评分：" + b.Score + "\n简介:" + b.Intro
}
