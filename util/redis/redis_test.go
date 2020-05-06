package redis

import (
	"github.com/corrots/go-demo/gocrawler/engine"
	"testing"
)

var key = "mylist"

func TestLRange(t *testing.T) {
	count :=  LLen(key)
	t.Log(count)
}

func TestLPush(t *testing.T) {
	value := engine.Request{
		URL:        "http://www.baidu.com",
		ParserFunc: nil,
	}
	err := LPush(key, value)
	if err != nil {
		t.Error(err)
	}
	//t.Log("successful")
}