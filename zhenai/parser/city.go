package parser

import (
	"crawler/engine"
	"regexp"
)
const cityRe=`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

//ParseCity 城市的解析器
func ParseCity(contents []byte) engine.ParserResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}
	for _, m := range matches {
		//将用户的名字作为Items
		result.Items = append(result.Items, "User "+string(m[2]))

		result.Requests = append(result.Requests,
			engine.Request{ 
				URL:        string(m[1]),
				ParserFunc: engine.NilParser,
			})
	}
	return result
}