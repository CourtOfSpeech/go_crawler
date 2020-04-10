package persist

import (
	"context"
	"crawler/engine"
	"encoding/json"
	"errors"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

//ItemServer ItemServer
func ItemServer() chan engine.Items {
	out := make(chan engine.Items)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Server: got item: #%d: %v", itemCount, item)
			itemCount++

			err := save(item)
			if err != nil {
				log.Printf("Item Server Error: saving item %v:%v", item, err)
			}
		}

	}()
	return out
}

//save 向elasticsearch存储
func save(item engine.Items) error {
	if item.Type == "" {
		return errors.New("must supply Type")
	}

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return err
	}
	//log.Println(elasticsearch.Version)
	//log.Println(es.Info())
	var b strings.Builder
	profile, err := json.Marshal(item)
	if err != nil {
		//log.Printf("Profile json parsing Error: %s", err)
		return err
	}
	b.WriteString(string(profile))

	req := esapi.IndexRequest{
		Index:        "dating_profile",
		DocumentType: item.Type,
		Body:         strings.NewReader(b.String()),
		Refresh:      "true",
	}
	//log.Printf("IndexRequest: %v", req)
	if item.ID != "" {
		req.DocumentID = item.ID
	}
	//log.Printf("IndexRequest: %v", req)

	res, err := req.Do(context.Background(), es)
	if err != nil {
		//log.Fatalf("Error getting response: %s", err)
		return err
	}
	defer res.Body.Close()
	var r map[string]interface{}
	if res.IsError() {
		//log.Printf("[%s]", res.Status())
		return errors.New("Response Status Code " + res.Status())
	}
	// Deserialize the response into a map.
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		//log.Printf("Error parsing the response body: %s", err)
		return err
	}

	//log.Printf("[%s]; version=%d; id=%v", res.Status(), int(r["_version"].(float64)), r["_id"])
	return nil
}
