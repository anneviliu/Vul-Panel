package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"io/ioutil"
	"log"
	"net/http"
)

type TimeRet struct {
	Code     int   `json:"code"`
	LastTime int64 `json:"msg"`
}

type Vul []struct {
	CreateTime int64 `json:"create_time"`
	Detail     struct {
		Host  string `json:"host"`
		Param struct {
			Key      string `json:"key"`
			Position string `json:"position"`
			Value    string `json:"value"`
		} `json:"param"`
		Payload  string `json:"payload"`
		Port     int    `json:"port"`
		Request  string `json:"request"`
		Response string `json:"response"`
		URL      string `json:"url"`
	} `json:"detail"`
	Plugin string `json:"plugin"`
	Target struct {
		URL    string `json:"url"`
		Params []struct {
			Position string   `json:"position"`
			Path     []string `json:"path"`
		} `json:"params"`
	} `json:"target"`
	VulnClass string `json:"vuln_class"`
}

func (s *Service) getTask() {
	s.watchFile("2.json")
}

func (s *Service) start(path string) {
	data := s.parseJson(s.readFile(path))
	time := s.getLastTime()
	for k, v := range data {
		// 如果某条json信息时间戳大于服务器记录时间戳，则向服务端发送该信息
		if v.CreateTime > time.LastTime {
			vulInfo, err_json := json.Marshal(v)
			if err_json != nil {
				fmt.Println("json字符串序列化失败")
			}
			fmt.Println("已发现新消息 " + string(k))
			req, err := http.NewRequest("POST", "http://"+s.Conf.ServerIP+"/getVulInfo", bytes.NewBuffer(vulInfo))
			req.Header.Set("Content-Type", "application/json")
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			fmt.Println(resp)
		}
	}
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

func (s *Service) readFile(path string) []byte {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("本地文件读取失败")
	}
	return contents
}

func (s *Service) parseJson(contents []byte) Vul {
	var data Vul
Repaired:
	jsonErr := json.Unmarshal(contents, &data)
	if jsonErr != nil {
		fmt.Println("json解析失败 正在尝试修复")
		//t := s.repairJson(contents)
		//fmt.Println(t)
		//s.parseJson([]byte(t))
		contents = []byte(s.repairJson(contents))
		goto Repaired
		return data
	}
	fmt.Println("json解析成功")
	return data
}

func (s *Service) repairJson(contents []byte) string {
	content := string(contents) + "]"
	//fmt.Println(content)
	return content
}

func (s *Service) watchFile(path string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	err = watcher.Add(path)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			select {
			case ev := <-watcher.Events:
				{
					//判断事件发生的类型，如下5种
					// Create 创建
					// Write 写入
					// Remove 删除
					// Rename 重命名
					// Chmod 修改权限
					if ev.Op&fsnotify.Write == fsnotify.Write {
						log.Println("写入文件 : ", ev.Name)
						s.start(path)
					}
				}
			case err := <-watcher.Errors:
				{
					log.Println("error : ", err)
					return
				}
			}
		}
	}()
	select {}
}
