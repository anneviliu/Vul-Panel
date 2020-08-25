package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/url"
	"time"
)

/*
WebHook 用于接受外部漏洞消息，以替代使用客户端轮询文件的功能
*/

type WebHookInfo struct {
	Url string `json:"url"`
	Type string `json:"type"`
	Request   string `json:"request"`
	Response  string `json:"response"`
}

// type url request response
func (s *Service) getVulInfoFromScanner(c *gin.Context) {
	//var request, resp, urlV, typeV, detail string
	var formData WebHookInfo
	err := c.BindJSON(&formData)
	if err != nil {
		c.JSON(400, gin.H{"errcode": 400, "description": "Post Data Err"})
		return
	}
	//s.add(formData)


	//buf := make([]byte, 40960)
	//n, _ := c.Request.Body.Read(buf)
	//
	//js, err := simplejson.NewJson(buf[0:n])
	//
	//if err != nil {
	//	fmt.Println("JSON 格式解析错误", err)
	//	return
	//}

	//detail, _ = js.Get("details").String()
	// 判断各个json字段是否存在
	//if _, flag := js.CheckGet("request"); flag {
	//	request, _ = js.Get("request").String()
	//} else {
	//	request = ""
	//}
	//
	//if _, flag := js.CheckGet("response"); flag {
	//	resp, _ = js.Get("response").String()
	//} else {
	//	resp = ""
	//}
	//
	//if _, flag := js.CheckGet("type"); flag {
	//	typeV, _ = js.Get("type").String()
	//} else {
	//	typeV = ""
	//}
	//
	//if _, flag := js.CheckGet("url"); flag {
	//	urlV, _ = js.Get("url").String()
	//} else {
	//	urlV = ""
	//}

	u, err := url.Parse(formData.Url)

	if err != nil {
		fmt.Println("URL解析失败")
		return
	}
	fmt.Println(u.Host)

	s.add(VulInfo{
		Timestamp: time.Now().UnixNano() / 1e6,
		Detail: Detail{
			//Details:   detail,
			Url:       formData.Url,
			Host:      u.Host,
			Payload:   "",
			Port:      0,
			Request:   formData.Request,
			Response:  formData.Response,
			Request1:  "",
			Response1: "",
		},
		Plugin:   "",
		VulClass: formData.Type,
	})
}
