package engine

import (
	"crawler/fetcher"
	"log"
)

func worker(r Request) (ParserResult, error) {
	log.Printf("Fething %s", r.URL)
	body, err := fetcher.Fetch(r.URL)
	if err != nil {
		//如果有错，就忽略这次的请求
		log.Printf("Fetcher: error fetcher url %s: %v", r.URL, err)
		//	continue
		return ParserResult{}, err
	}

	//如果没有错就解析body
	return r.ParserFunc(body, r.URL), nil
}
