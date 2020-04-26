package parser

import (
	"regexp"
	"strconv"

	"github.com/corrots/go-demo/gocrawler/engine"

	"github.com/corrots/go-demo/gocrawler/schema"
)

var (
	urlReg = regexp.MustCompile(`https://m.zhenai.com/u/(\d+).html`)

	genderReg    = regexp.MustCompile(`<span[^>]+>关注(\p{Han})</span>`)
	marriageReg  = regexp.MustCompile(`<div [^>]*tag[^>]+>([未婚离异丧偶]+)</div>`)
	ageReg       = regexp.MustCompile(`<div [^>]*tag[^>]+>(\d+)岁</div>`)
	xinzuoReg    = regexp.MustCompile(`<div [^>]*tag[^>]+>(\p{Han}+)\([0-9.-]+\)</div>`)
	heightReg    = regexp.MustCompile(`<div [^>]*tag[^>]+>(\d+)cm</div>`)
	weightReg    = regexp.MustCompile(`<div [^>]*tag[^>]+>(\d+)kg</div>`)
	incomeReg    = regexp.MustCompile(`<div [^>]*tag[^>]+>月收入:([^<]+)</div>`)
	educationReg = regexp.MustCompile(`<div [^>]*tag[^>]+>([高中及以下大学本科硕士]+)</div>`)
	hokouReg     = regexp.MustCompile(`<div [^>]*tag[^>]+>籍贯:(\p{Han}+)</div>`)
	carReg       = regexp.MustCompile(`<div [^>]*tag[^>]+>(\p{Han}+车)</div>`)
	houseReg     = regexp.MustCompile(`<div [^>]*tag[^>]+>(\p{Han}+房)</div>`)
)

func ParseProfile(c []byte, name string) engine.ParseResult {
	profile := schema.Profile{
		Name:      name,
		Gender:    convertGender(extractString(c, genderReg)),
		Age:       extractInt(c, ageReg),
		Height:    extractInt(c, heightReg),
		Weight:    extractInt(c, weightReg),
		Income:    extractString(c, incomeReg),
		Marriage:  extractString(c, marriageReg),
		Education: extractString(c, educationReg),
		Hokou:     extractString(c, hokouReg),
		Xinzuo:    extractString(c, xinzuoReg),
		House:     extractString(c, houseReg),
		Car:       extractString(c, carReg),
	}
	var result engine.ParseResult
	//fmt.Printf("Got profile: %+v\n", profile)
	result.Items = append(result.Items, profile)
	return result
}

func convertGender(gender string) string {
	if gender == "他" {
		return "男"
	}
	return "女"
}

func extractString(contents []byte, reg *regexp.Regexp) string {
	matches := reg.FindSubmatch(contents)
	if len(matches) >= 2 {
		return string(matches[1])
	}
	return ""
}

func extractInt(contents []byte, reg *regexp.Regexp) int {
	s, err := strconv.Atoi(extractString(contents, reg))
	if err == nil {
		return s
	}
	return 0
}
