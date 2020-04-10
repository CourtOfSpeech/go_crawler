package engine

//ParserFunc ParserFunc
type ParserFunc func(contents []byte, url string) ParserResult

//Request  ...
type Request struct {
	URL        string
	ParserFunc ParserFunc
}

//ParserResult 解析器返回内容的
type ParserResult struct {
	Requests []Request
	Items    []Items
}

//Items Items必须包含ID URL
type Items struct {
	URL     string      `json:"url,omitempty" xml:"url"`
	Type    string      `json:"type,omitempty" xml:"type"`
	ID      string      `json:"id,omitempty" xml:"id"`
	Payload interface{} `json:"payload,omitempty" xml:"payload"` //具体的内容
}

//NilParser 处理nil的情况
func NilParser([]byte) ParserResult {
	return ParserResult{}
}
