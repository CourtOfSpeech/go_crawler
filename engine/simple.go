package engine

import (
	"log"
)

//SimpleEngine 简单版的engine
type SimpleEngine struct {
}

//Run 根据传入的URL做事
func (e SimpleEngine) Run(seeds ...Request) {
	//将需要做的事情放入一个队列
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	//只有队列里有数据，就代表有事情要做
	for len(requests) > 0 {
		//拿到队列的第一个值，去做事情
		r := requests[0]
		requests = requests[1:]

		parserResult, err := worker(r)
		if err != nil {
			//如果有错，就忽略这次的请求
			log.Printf("Fetcher: error fetcher url %s: %v", r.URL, err)
			continue
		}

		//把这次拿到的URL，再次放入请求队列requests里
		requests = append(requests, parserResult.Requests...)

		for _, item := range parserResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}
