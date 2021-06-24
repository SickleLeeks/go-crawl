package main

import (
	"crawl/LearnGo-crawl/crawl_distribute/client"
	client2 "crawl/LearnGo-crawl/crawl_distribute/work/client"
	"crawl/LearnGo-crawl/engine"
	"crawl/LearnGo-crawl/parse/zhenai"
	"crawl/LearnGo-crawl/scheduler"
)

func main() {
	itemsave, err := client.ItemSave(":1234")
	process, err := client2.CreateProcess()
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		&scheduler.QueueScheduler{},
		100,
		itemsave,
		process,
	}
	e.Run(engine.Request{
		Url:   "https://book.douban.com",
		Parse: engine.NewFuncparse(zhenai.ParseCityList, "ParseCityList"),
	})
}
