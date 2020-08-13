package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
	"net/url"
	"time"
)

/*
WebHook 用于接受外部漏洞消息，以替代使用客户端轮询文件的功能
*/

type info struct {
}

func (s *Service) getVulInfoFromScanner(c *gin.Context) {
	var request, resp, urlV, typeV, detail string

	buf := make([]byte, 40960)
	n, _ := c.Request.Body.Read(buf)

	js, err := simplejson.NewJson(buf[0:n])

	if err != nil {
		fmt.Println("JSON 格式解析错误", err)
		return
	}

	detail, _ = js.Get("details").String()
	// 判断各个json字段是否存在
	if _, flag := js.CheckGet("rawrequest"); flag {
		request, _ = js.Get("rawrequest").String()
	} else {
		request = ""
	}

	if _, flag := js.CheckGet("rawresponse"); flag {
		resp, _ = js.Get("rawresponse").String()
	} else {
		resp = ""
	}

	if _, flag := js.CheckGet("type"); flag {
		typeV, _ = js.Get("type").String()
	} else {
		typeV = ""
	}

	if _, flag := js.CheckGet("url"); flag {
		urlV, _ = js.Get("url").String()
	} else {
		urlV = ""
	}

	u, err := url.Parse(urlV)

	if err != nil {
		fmt.Println("URL解析失败")
		return
	}
	fmt.Println(u.Host)

	s.add(VulInfo{
		Timestamp: time.Now().UnixNano() / 1e6,
		Detail: Detail{
			Details:   detail,
			Url:       urlV,
			Host:      u.Host,
			Payload:   "",
			Port:      0,
			Request:   request,
			Response:  resp,
			Request1:  "",
			Response1: "",
		},
		Plugin:   "",
		VulClass: typeV,
	})
}
