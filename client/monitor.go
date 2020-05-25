package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

//TODO: 通过手机端实时查看扫描进程状态

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
		Payload   string `json:"payload"`
		Port      int    `json:"port"`
		Request   string `json:"request"`
		Response  string `json:"response"`
		Request1  string `json:"request1"`
		Response1 string `json:"response1"`
		URL       string `json:"url"`
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
	if s.Conf.Debug {
		s.watchFile("2.json")
	} else {
		var (
			name string
		)
		flag.StringVar(&name, "name", "", "Your json name here")
		flag.Parse()
		if name != "" {
			s.beginScan(name)
			s.watchFile(name)
		} else {
			flag.Usage()
		}
	}
}

func (s *Service) start(path string) {
	data := s.parseJson(s.readFile(path))
	timeL := s.getLastTime()

	for k, v := range data {
		// 如果某条json信息时间戳大于服务器记录时间戳，则向服务端发送该信息
		fmt.Println(v.CreateTime)
		if v.CreateTime > timeL.LastTime {
			vulInfo, errJson := json.Marshal(v)
			if errJson != nil {
				fmt.Println("json字符串序列化失败")
				return
			}
			fmt.Println("已发现新消息 " + string(k))
			s.pushVulInfo(vulInfo)
		} else {
			fmt.Println("监测到修改，无最新漏洞消息")
		}
	}
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
	var i int
Repaired:
	jsonErr := json.Unmarshal(contents, &data)
	if jsonErr != nil {
		if i > 5 {
			fmt.Println("json解析失败次数过多 终止运行")
			panic(jsonErr)
		}
		contents = []byte(s.repairJson(contents))
		i++
		goto Repaired
		return data
	}
	fmt.Println("json解析成功")
	return data
}

func (s *Service) repairJson(contents []byte) string {
	var content string
	if contents[len(contents)-1] == ',' {
		fmt.Println("json解析失败 正在尝试修复 ',' ")
		content = string(contents[0 : len(contents)-1])
		return content
	}
	if contents[len(contents)-1] == ']' {
		fmt.Println("json解析失败 正在尝试修复 ']' ")
		return string(contents)
	} else {
		content = string(contents) + "]"
	}
	return content
}

func (s *Service) watchFile(path string) {
	// 连接redis
	red, err := redis.Dial("tcp", s.Conf.Base.RedisIP)
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}

	defer red.Close()
	for {
		f, err := os.Open(path)
		if err != nil {
			fmt.Println("打开漏洞json文件错误", err)
		} else {
			fi, err2 := f.Stat()
			if err2 != nil {
				fmt.Println("获取文件状态失败")
				return
			}
			//文件每次修改时间存入redis
			oldTime, err := redis.String(red.Do("GET", "modTime"))

			if err != nil {
				fmt.Println("redis get failed:", err)
				return
			}

			var strInt int64
			strInt, err = strconv.ParseInt(oldTime, 10, 64)
			fmt.Println(fi.ModTime().Unix())
			if fi.ModTime().Unix() > strInt {
				s.start(path)
			} else {
				fmt.Println("正在轮询")
			}
			_, err = red.Do("SET", "modTime", fi.ModTime().Unix())

			if err != nil {
				fmt.Println("Redis set failed:", err)
			}
		}

		time.Sleep(time.Duration(1) * time.Second)
		f.Close()
	}
}
