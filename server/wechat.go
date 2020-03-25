package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type RetMsg struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
}

type SendMsg struct {
	Touser               string   `json:"touser"`
	MsgType              string   `json:"msgtype"`
	AgentID              int      `json:"agentid"`
	TextCard             TextCard `json:"textcard"`
	EnableIdTrans        int      `json:"enable_id_trans"`
	EnableDuplicateCheck int      `json:"enable_duplicate_check	"`
}

type TextCard struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	BtnTxt      string `json:"btntxt"`
}

func (s *Service) StartWeChat(data VulInfo) {
	accessToken := s.getAccessToken()
	fmt.Println(accessToken)
	template := s.buildHtml(data)
	sendUrl := fmt.Sprintf(s.Conf.SendMsgUrl, accessToken)
	req, err := http.NewRequest("POST", sendUrl, bytes.NewBuffer(template))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	var r RetMsg
	json.Unmarshal([]byte(string(body)), &r)
	if r.ErrMsg == "ok" {
		s.addPushed(data)
	}

}

func (s *Service) getAccessToken() string {
	url := fmt.Sprintf(s.Conf.AccessTokenUrl, s.Conf.CorpID, s.Conf.CorpSecret)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln("Wechat access_token get err", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("read resp body err", err)
	}
	var r RetMsg
	json.Unmarshal([]byte(string(body)), &r)
	//fmt.Println(r.AccessToken)
	return r.AccessToken
}

func (s *Service) buildHtml(data VulInfo) []byte {
	t := SendMsg{
		Touser:  "@all",
		MsgType: "textcard",
		AgentID: s.Conf.AgentID,
		TextCard: TextCard{
			Title: "Xray漏洞通知",
			Description: fmt.Sprintf("<div class=\"gray\">%s</div><div class=\"red\">Type: %s</div><br><div class=\"red\">Url: %s</div><br><div class=\"red\">Payload: %s</div>",
				time.Now().Format("2006-01-02 15:04:05"),
				data.Plugin,
				data.Detail.Url,
				data.Detail.Payload,
			),
			BtnTxt: "详情",
			Url:    "URL",
		},
		EnableIdTrans:        0,
		EnableDuplicateCheck: 0,
	}
	res, err := json.Marshal(t)
	if err != nil {
		log.Fatalln("read json err", err)
	}
	return res
}
