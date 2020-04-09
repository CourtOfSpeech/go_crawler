package persist

import (
	"bytes"
	"context"
	"crawler/model"
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/mitchellh/mapstructure"
)

func Test_save(t *testing.T) {
	type args struct {
		item interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"在水伊人", args{item: model.Profile{
			Name:   "在水伊人",
			Age:    44,
			Height: 155,
			Income: "8千-1.2万",
			Xinzuo: "魔羯座",
			Hokou:  "四川成都",
			House:  "已购房",
			Car:    "未买车"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := save(tt.args.item)
			if err != nil {
				fmt.Println(err)
			}
			//TODO: Try to start up elasticsearch
			//here using docker go client
			es, err := elasticsearch.NewDefaultClient()
			if err != nil {
				fmt.Println(err)
			}
			// Search for the indexed documents
			// Build the request body.
			var buf bytes.Buffer
			query := map[string]interface{}{
				"query": map[string]interface{}{
					"match": map[string]interface{}{
						"_id": id,
					},
				},
			}
			if err := json.NewEncoder(&buf).Encode(query); err != nil {
				log.Fatalf("Error encoding query: %s", err)
			}
			// Perform the search request.
			res, err := es.Search(
				es.Search.WithContext(context.Background()),
				es.Search.WithIndex("dating_profile"),
				es.Search.WithDocumentType("zhenai"),
				es.Search.WithBody(&buf),
				es.Search.WithTrackTotalHits(true),
				es.Search.WithPretty(),
			)
			if err != nil {
				log.Fatalf("Error getting response: %s", err)
			}
			defer res.Body.Close()
			var r map[string]interface{}
			if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
				log.Fatalf("Error parsing the response body: %s", err)
			}

			//Print the ID and document source for each hit.
			for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
				//log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
				var person model.Profile
				err := mapstructure.Decode(hit.(map[string]interface{})["_source"], &person)
				if err != nil {
					fmt.Println(err)
				}
				if tt.args.item != person {
					t.Errorf("got %v; expected %v", person, tt.args.item)
				}
			}
		})
	}
}
