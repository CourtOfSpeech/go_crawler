package main

import (
	"crawler/engine"
	"crawler/persist"
	"crawler/scheduler"
	"crawler/zhenai/parser"
)

func main() {
	// engine.SimpleEngine{}.Run(engine.Request{
	// 	URL:        "https://www.zhenai.com/zhenghun",
	// 	ParserFunc: parser.ParseCityList,
	// })
	itemChan, err := persist.ItemServer("dating_profile")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueudeScheduler{},
		WorkerCount: 10,
		ItemChan:    itemChan,
	}
	// e.Run(engine.Request{
	// 	URL:        "https://www.zhenai.com/zhenghun",
	// 	ParserFunc: parser.ParseCityList,
	// })

	e.Run(engine.Request{
		URL:        "https://www.zhenai.com/zhenghun/shanghai",
		ParserFunc: parser.ParseCity,
	})
}
