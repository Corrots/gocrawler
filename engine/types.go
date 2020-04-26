package engine

type ParserFunc func(c []byte) ParseResult

type Request struct {
	URL string
	ParserFunc
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
