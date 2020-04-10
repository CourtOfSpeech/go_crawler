package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	//contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	//从拷贝下来的静态文件读取,拷贝的文件最好不要格式化
	contents, err := ioutil.ReadFile("city_list_data.html")
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%s\n", contents)
	result := ParseCityList(contents, "")

	const resultSize = 470
	expectdeURLs := []string{
		"http://www.zhenai.com/zhenghun/aba", "http://www.zhenai.com/zhenghun/akesu", "http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectdecities := []string{
		"City 阿坝", "City 阿克苏", "City 阿拉善盟",
	}
	//判断读取的内容长度是否相同
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests; but had %d", resultSize, len(result.Requests))
	}

	for i, url := range expectdeURLs {
		if result.Requests[i].URL != url {
			t.Errorf("expected url #%d: %s ;but was %s", i, url, result.Requests[i].URL)
		}
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d requests; but had %d", resultSize, len(result.Items))
	}

	for i, city := range expectdecities {
		if result.Items[i].Payload.(string) != city {
			t.Errorf("expected city #%d: %s ;but was %s", i, city, result.Items[i].Payload.(string))
		}
	}
}
