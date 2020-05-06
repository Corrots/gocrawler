package parser

import (
	"fmt"
	"github.com/corrots/go-demo/gocrawler/util/redis"
	"log"
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
		if isDuplicate(url) {
			continue
		}
		name := string(val[2])
		//fmt.Printf("Got User: %s, url: %s\n", name, url)
		result.Requests = append(result.Requests, engine.Request{
			URL:        url,
			ParserFunc: ProfileParser(name),
		})
	}
	return result
}

func isDuplicate(url string) bool {
	fmt.Println("duplicate validation: ",url)
	if redis.SIsMember("url", url) {
		return true
	}
	go func() {
		err := redis.SAdd("url", url)
		//log.Printf("redis set url: %s\n", r.URL)
		if err != nil {
			log.Printf("redis set url: %s err: %v\n", url, err)
		}
	}()
	return false
}

func ProfileParser(name string) engine.ParserFunc {
	return func(c []byte) engine.ParseResult {
		return ParseProfile(c, name)
	}
}
