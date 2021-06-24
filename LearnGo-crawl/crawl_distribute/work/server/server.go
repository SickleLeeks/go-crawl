package server

import (
	"crawl/LearnGo-crawl/crawl_distribute/work"
	"crawl/LearnGo-crawl/engine"
)

type CrawlService struct {
}

func (CrawlService) Process(req engine.Request, result *work.ParseResult) error {
	engineReq, err := work.DeserializeRequest(req)
	if err != nil {
		return err
	}
	engineResult, err := engine.Worker(engineReq)
	*result = work.SerializeResult(engineResult)
	return nil
}
