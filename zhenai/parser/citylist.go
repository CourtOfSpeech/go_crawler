package parser

import (
	"crawler/engine"
	"regexp"
	"strings"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

//ParseCityList 城市列表的解析器
func ParseCityList(contents []byte) engine.ParserResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}

	limit := 100
	for _, m := range matches {
		//将城市的名字作为Items
		result.Items = append(result.Items, "City "+string(m[2]))

		url := strings.Replace(string(m[1]), "http", "https", 1)
		result.Requests = append(result.Requests,
			engine.Request{
				URL:        url,
				ParserFunc: ParseCity,
			})
		//fmt.Printf("cityUrl: %s, cityName: %s\n", m[1], m[2])
		limit--
		if limit == 0 {
			break
		}
	}
	//fmt.Printf("Matches found: %d\n", len(matches))
	return result
}
