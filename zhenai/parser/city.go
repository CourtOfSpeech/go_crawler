package parser

import (
	"crawler/engine"
	"regexp"
	"strings"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

//ParseCity 城市的解析器
func ParseCity(contents []byte) engine.ParserResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}
	for _, m := range matches {
		name := string(m[2])
		//将用户的名字作为Items
		result.Items = append(result.Items, "User "+name)

		url := strings.Replace(string(m[1]), "http", "https", 1)
		result.Requests = append(result.Requests,

			engine.Request{
				URL: url,
				ParserFunc: func(c []byte) engine.ParserResult {
					return ParseProfile(c, name)
				},
			})
	}
	return result
}
