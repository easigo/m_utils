package m_verify

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/handy-golang/go-tools/m_str"
)

type IPAddressType struct {
	Hostname  string
	ISP       string
	Operators string
}

// 获取IP地址归属
func GetIPS(ips []string) []IPAddressType {
	rList := []IPAddressType{}
	for _, val := range ips {
		res, _ := GetIPaddress(val)
		if len(res.Hostname) > 0 {
			rList = append(rList, res)
		}
	}

	return rList
}

/*
https://www.ipshudi.com/36.44.232.38.htm
*/

func GetIPaddress(ip string) (resData IPAddressType, resErr error) {
	if !IsIP(ip) {
		resErr = fmt.Errorf("ip地址不正确")
		return
	}

	HeaderMap := FileToHeader(WhatIsMyIpHeader)
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		for key, val := range HeaderMap {
			r.Headers.Set(key, val)
		}
	})
	c.OnError(func(r *colly.Response, errStr error) {
		resErr = errStr
	})
	// 获取IP
	c.OnHTML(".input-text", func(e *colly.HTMLElement) {
		value := e.Attr("value")
		resData.Hostname = value
	})
	// 获取运营商和归属地
	c.OnHTML("td.th+td", func(e *colly.HTMLElement) {
		isA := e.DOM.Find(".report")

		if len(isA.Text()) > 0 {
			span := e.DOM.Find("span")
			resData.ISP = span.Text()
		} else {
			span := e.DOM.Find("span")
			resData.Operators = span.Text()
		}
	})

	tmplStr := `https://www.ipshudi.com/${IP}.htm`
	tmplVal := map[string]string{
		"IP": ip,
	}
	FetchUrl := m_str.Temp(tmplStr, tmplVal)

	c.Visit(FetchUrl)
	if resErr != nil {
		return
	}
	// 请求完成
	if !IsIP(resData.Hostname) {
		resErr = fmt.Errorf("未获取到指定IP")
		return
	}

	return
}

// ========= 解析 Header 头 ===============

//go:embed WhatIsMyIpHeader.yaml
var WhatIsMyIpHeader string

func FileToHeader(cont string) map[string]string {
	strArr := strings.Split(cont, "\n")
	HeaderMap := make(map[string]string)
	for _, item := range strArr {
		kvArr := strings.Split(item, ": ")
		if len(kvArr) == 2 {
			k := kvArr[0]
			v := kvArr[1]
			HeaderMap[k] = v
		}
	}
	return HeaderMap
}
