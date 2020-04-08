package persist

import (
	"context"
	"crawler/model"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

//ItemServer ItemServer
func ItemServer() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Server: got item: #%d: %v", itemCount, item)
			itemCount++

			save(item)
		}

	}()
	return out
}

//save 向elasticsearch存储
func save(item interface{}) {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		fmt.Println(err)
	}
	//log.Println(elasticsearch.Version)
	//log.Println(es.Info())
	var b strings.Builder
	if profile, ok := item.(model.Profile); ok {
		proStr, err := json.Marshal(profile)
		if err != nil {
			fmt.Printf("Profile json parsing Error: %s", err)
		}
		b.WriteString(string(proStr))
		fmt.Println(string(proStr))
	}
	req := esapi.IndexRequest{
		Index:        "dating_profile",
		DocumentType: "zhenai",
		Body:         strings.NewReader(b.String()),
		Refresh:      "true",
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("[%s]", res.Status())
	} else {
		// Deserialize the response into a map.
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and indexed document version.
			log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
		}
	}

}
