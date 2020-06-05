package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
)

/*
WebHook 用于接受外部漏洞消息，以替代使用客户端轮询文件的功能
*/

type info struct {
}

func (s *Service) getVulInfoFromScanner(c *gin.Context) {
	buf := make([]byte, 40960)
	n, _ := c.Request.Body.Read(buf)

	js, err := simplejson.NewJson(buf[0:n])
	if err != nil {
		fmt.Println("JSON 格式解析错误", err)
		return
	}
	detail, _ := js.Get("Details").Map()
	fmt.Println(detail)
	for k, v := range detail {
		fmt.Println(k, " ", v)
	}
	//request, err := js.Get("RawRequest").String()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(request)
	//
	//resp, err := js.Get("RawResponse").String()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(resp)

	//s.add()
	//fmt.Println(detail.Get("Param").String())
}
