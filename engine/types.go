package engine

//Request  ...
type Request struct {
	URL        string
	ParserFunc func([]byte) ParserResult
}

//ParserResult 解析器返回内容的
type ParserResult struct {
	Requests []Request
	Items    []interface{}
}

//NilParser 处理nil的情况
func NilParser([]byte) ParserResult {
	return ParserResult{}
}
