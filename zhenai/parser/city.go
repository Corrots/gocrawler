package parser

import (
	"regexp"
	"strings"

	"github.com/corrots/go-demo/gocrawler/engine"
)

const cityReg = `<a href="(http://album.zhenai.com/u/\d+)"[^>]*>([^>]+)</a>`

func ParseCity(c []byte) engine.ParseResult {
	re := regexp.MustCompile(cityReg)
	matches := re.FindAllSubmatch(c, -1)
	var result engine.ParseResult
	for _, val := range matches {
		url := strings.Replace(string(val[1]), "album", "m", 1)
		name := string(val[2])
		//fmt.Printf("Got User: %s, url: %s\n", name, url)
		result.Requests = append(result.Requests, engine.Request{
			URL:        url,
			ParserFunc: ProfileParser(name),
		})
	}
	return result
}

func ProfileParser(name string) engine.ParserFunc {
	return func(c []byte) engine.ParseResult {
		return ParseProfile(c, name)
	}
}
