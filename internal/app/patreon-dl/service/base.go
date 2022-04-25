package service

import (
	"io/ioutil"
	"os"
)

const (
	perm   os.FileMode = 644
	format string      = "2006-01-02T15:04:05.999999999-0700"
)

var (
	header = map[string][]string{
		"accept":                    {"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"},
		"accept-encoding":           {"gzip, deflate, br"},
		"accept-language":           {"zh-CN,zh;q=0.9,en-GB;q=0.8,en;q=0.7,zh-TW;q=0.6"},
		"cache-control":             {"no-cache"},
		"dnt":                       {"1"},
		"pragma":                    {"no-cache"},
		"sec-ch-ua":                 {"\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"100\", \"Google Chrome\";v=\"100\""},
		"sec-ch-ua-mobile":          {"?0"},
		"sec-ch-ua-platform":        {"Windows"},
		"sec-fetch-dest":            {"document"},
		"sec-fetch-mode":            {"navigate"},
		"sec-fetch-site":            {"none"},
		"sec-fetch-user":            {"?1"},
		"upgrade-insecure-requests": {"1"},
		"user-agent":                {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36"},
	}
)

func init() {
	b, err := ioutil.ReadFile("cookies.txt")
	if err != nil {
		panic(err)
	}
	header["cookie"] = []string{string(b)}
}
