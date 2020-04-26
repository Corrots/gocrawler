package parser

import (
	"regexp"

	"github.com/corrots/go-demo/gocrawler/engine"
)

const cityListReg = `<a href="(http://www.zhenai.com/zhenghun/[\w]+)" data-v-[\w]+>([\p{Han}]+)</a>`

func ParseCityList(c []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListReg)
	matches := re.FindAllSubmatch(c, -1)
	var result engine.ParseResult
	var limit = 2
	for _, val := range matches {
		//fmt.Printf("Got City: %s, URL: %s\n", fmt.Sprintf("%s", val[2]), val[1])
		result.Requests = append(result.Requests, engine.Request{
			URL:        string(val[1]),
			ParserFunc: ParseCity,
		})

		if limit == 0 {
			break
		}
		limit--
	}
	//fmt.Println(len(matches))
	return result
}
