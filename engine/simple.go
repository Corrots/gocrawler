package engine

import (
	"log"

	"github.com/corrots/go-demo/gocrawler/fetcher"
)

type SimpleEngine struct{}

func (e *SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		request := requests[0]
		requests = requests[1:]
		result, err := worker(request)
		if err != nil {
			continue
		}
		requests = append(requests, result.Requests...)
	}
}

func worker(r Request) (ParseResult, error) {
	c, err := fetcher.Fetch(r.URL)
	if err != nil {
		log.Printf("Fetching URL: %s, err: %v\n", r.URL, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(c), nil
}
