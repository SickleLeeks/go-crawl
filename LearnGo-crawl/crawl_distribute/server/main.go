package main

import (
	"crawl/LearnGo-crawl/crawl_distribute/persist"
	"crawl/LearnGo-crawl/crawl_distribute/rpcsupport"
	"gopkg.in/olivere/elastic.v5"
)

func serveRpc(host string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}
	return rpcsupport.ServeRpc(host, &persist.ItemService{
		Client: client,
	})
}
func main() {
	serveRpc(":1234")
}
