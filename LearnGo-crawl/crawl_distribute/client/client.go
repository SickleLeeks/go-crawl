package client

import (
	"crawl/LearnGo-crawl/crawl_distribute/rpcsupport"
	"crawl/LearnGo-crawl/engine"
	"log"
)

func ItemSave(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)

	go func() {
		itemcount := 0
		for {
			item := <-out
			log.Printf("Item save: Got$%d,%v", itemcount, item)
			result := ""
			err = client.Call("ItemService.Save", item, &result)
			if err != nil {
				log.Printf("item save:error saving item %v:%w", item, err)
			}
			itemcount++
		}
	}()
	return out, nil
}
