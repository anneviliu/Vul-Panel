package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"html"
)

type Vul struct {
	gorm.Model
	Host     string
	Port     int
	Url      string `gorm:"size:1000"`
	Title    string
	Payload  string `gorm:"size:999999"`
	Request  string `gorm:"size:999999"`
	Response string `gorm:"size:999999"`
	Times    int64  `gorm:"size:100"`
	VulClass string `gorm:"size:100"`
}

type VulInfo struct {
	Timestamp int64  `json:"create_time"`
	Detail    Detail `json:"detail"`
	Plugin    string `json:"plugin"`
	VulClass  string `json:"vuln_class"`
}

type Detail struct {
	FileName  string `json:"filename"`
	Url       string `json:"url"`
	Host      string `json:"host"`
	Payload   string `json:"payload"`
	Port      int    `json:"port"`
	Request   string `json:"request"`
	Response  string `json:"response"`
	Request1  string `json:"request1"`
	Response1 string `json:"response1"`
}

// 从客户端获取漏洞信息
func (s *Service) getVulInfo(c *gin.Context) {
	var formData VulInfo
	err := c.BindJSON(&formData)
	if err != nil {
		c.JSON(400, gin.H{"errcode": 400, "description": "Post Data Err"})
		return
	}
	s.add(formData, c)
}

// 向数据库中插入漏洞信息
func (s *Service) add(data VulInfo, c *gin.Context) {
	if data.Detail.Request == "" {
		data.Detail.Request = data.Detail.Request1
	}

	if data.Detail.Response == "" {
		data.Detail.Response = data.Detail.Response1
	}

	if data.VulClass == "" {
		data.VulClass = data.Plugin + " [plugin]"
	}
	vulData := &Vul{
		Host:     data.Detail.Host,
		Port:     data.Detail.Port,
		Url:      data.Detail.Url,
		Title:    data.Plugin,
		Payload:  data.Detail.Payload,
		Request:  html.EscapeString(data.Detail.Request),
		Response: html.EscapeString(data.Detail.Response),
		Times:    data.Timestamp,
		VulClass: data.VulClass,
	}
	if !s.check(data) {
		fmt.Printf("重复插入记录")
	} else {
		s.Mysql.Create(vulData)

		s.writeHTML(data)

		s.StartWeChat(data)
	}
}

// 检查该漏洞是否已记录
func (s *Service) check(data VulInfo) bool {
	a := s.Mysql.Model(&Pushed{}).Where(&Pushed{
		Model: gorm.Model{},
		/*
		 *  修复 不得根据request和resp来判断
		 *  应使用url和plugin 和 payload来判断
		 */

		//Request:  data.Detail.Request,
		//Response: data.Detail.Response,

		Url:     data.Detail.Url,
		Title:   data.Plugin,
		Payload: data.Detail.Payload,
		Times:   data.Timestamp,
	}).First(&Pushed{}).RowsAffected
	if a == 1 {
		return false
	} else {
		return true
	}
}
