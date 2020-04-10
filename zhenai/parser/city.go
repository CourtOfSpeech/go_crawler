package parser

import (
	"crawler/engine"
	"regexp"
	"strings"
)

var (
	profileRe = regexp.MustCompile(
		`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityURLRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

//ParseCity 城市的解析器
func ParseCity(contents []byte, _ string) engine.ParserResult {
	matches := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}
	for _, m := range matches {
		url := strings.Replace(string(m[1]), "http", "https", 1)
		result.Requests = append(result.Requests,

			engine.Request{
				URL:        url,
				ParserFunc: ProfileParse(string(m[2])),
			})
	}
	//城市
	matches = cityURLRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		url := strings.Replace(string(m[1]), "http", "https", 1)
		result.Requests = append(result.Requests,
			engine.Request{
				URL:        url,
				ParserFunc: ParseCity,
			})
	}

	return result
}

//ProfileParse ProfileParse
func ProfileParse(name string) engine.ParserFunc {
	return func(c []byte, url string) engine.ParserResult {
		return ParseProfile(c, url, name)
	}
}
