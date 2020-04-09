package persist

import (
	"context"
	"crawler/model"
	"encoding/json"
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

			_, err := save(item) 
			if err != nil {
				log.Printf("Item Server Error: saving item %v:%v", item, err)
			}
		}

	}()
	return out
}

//save 向elasticsearch存储
func save(item interface{}) (string, error) {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return "", err
	}
	//log.Println(elasticsearch.Version)
	//log.Println(es.Info())
	var b strings.Builder
	if profile, ok := item.(model.Profile); ok {
		proStr, err := json.Marshal(profile)
		if err != nil {
			log.Printf("Profile json parsing Error: %s", err)
			return "", err
		}
		b.WriteString(string(proStr))
		//fmt.Println(string(proStr))
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
		return "", err
	}
	defer res.Body.Close()
	var r map[string]interface{}
	if res.IsError() {
		log.Printf("[%s]", res.Status())
	} else {
		// Deserialize the response into a map.
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Printf("Error parsing the response body: %s", err)
			return "", err
		}
	}
	//log.Printf("[%s]; version=%d; id=%v", res.Status(), int(r["_version"].(float64)), r["_id"])
	return r["_id"].(string), nil
}
