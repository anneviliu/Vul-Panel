package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (s *Service) pushVulInfo(data []byte) {
	req, err := http.NewRequest("POST", "http://"+s.Conf.ServerIP+"/getVulInfo", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

func (s *Service) beginScan(name string) {
	n := strings.Split(name, ".")
	tmp := n[0]

	tm := strconv.FormatInt(time.Now().Unix(), 10)
	jsonContent := "{\"host\":\"%s\",\"time\":%s}"
	jsonContent = fmt.Sprintf(jsonContent, tmp, tm)
	req, err := http.NewRequest("POST", "http://"+s.Conf.ServerIP+"/scannerStart", bytes.NewBuffer([]byte(jsonContent)))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

func (s *Service) getLastTime() TimeRet {
	resp, err := http.Get("http://" + s.Conf.Base.ServerIP + "/getLastTime")
	if err != nil {
		fmt.Println("请求时间戳接口失败")
		return TimeRet{
			Code:     200,
			LastTime: 0,
		}
	}
	//defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("获取时间戳响应体失败")
		return TimeRet{
			Code:     200,
			LastTime: 0,
		}
	}
	var timedata TimeRet
	json.Unmarshal(body, &timedata)
	return timedata
}
