package parser

import (
	"crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

//ParseCityList 城市的解析器
func ParseCityList(contents []byte) engine.ParserResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}
	for _, m := range matches {
		//将城市的名字作为Items
		result.Items = append(result.Items, string(m[2]))

		result.Requests = append(result.Requests,
			engine.Request{
				URL:        string(m[1]),
				ParserFunc: engine.NilParser,
			})
		//fmt.Printf("cityUrl: %s, cityName: %s\n", m[1], m[2])
	}
	//fmt.Printf("Matches found: %d\n", len(matches))
	return result
}
