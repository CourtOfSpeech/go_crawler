package engine

//Request  ...
type Request struct {
	URL        string
	ParserFunc func([]byte) ParserResult
}

//ParserResult 解析器返回内容的
type ParserResult struct {
	Requests []Request
	Items    []Items
}

//Items Items必须包含ID URL
type Items struct {
	URL     string
	Type    string
	ID      string
	Payload interface{} //具体的内容
}

//NilParser 处理nil的情况
func NilParser([]byte) ParserResult {
	return ParserResult{}
}
