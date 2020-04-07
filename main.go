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

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueudeScheduler{},
		WorkerCount: 10,
		ItemChan:    persist.ItemServer(),
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
