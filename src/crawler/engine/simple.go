package engine

import (
	"crawler/fetcher"
	"log"
)

type SimpleEngine struct {
}

func (e SimpleEngine) Run(seeds ...Request) {
	var q []Request
	for _, r := range seeds {
		q = append(q, r)
	}

	for len(q) > 0 {
		r := q[0]
		q = q[1:]

		parseResult, err := worker(r)
		if err != nil {
			continue
		}
		q = append(q, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}

	}
}

func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error "+"fecher url %s: %v\n", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body), nil
}
