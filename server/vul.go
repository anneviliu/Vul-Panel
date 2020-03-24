package main

import (
	"github.com/jinzhu/gorm"
)
import "github.com/gin-gonic/gin"

type Vul struct {
	gorm.Model
	Host     string
	Port     int
	Url      string `gorm:"size:1000"`
	Title    string
	Payload  string `gorm:"size:9999"`
	Request  string `gorm:"size:999999"`
	Response string `gorm:"size:999999"`
}

type VulInfo struct {
	Detail Detail `json:"detail"`
	Plugin string `json:"plugin"`
}

type Detail struct {
	FileName string `json:"filename"`
	Url      string `json:"url"`
	Host     string `json:"host"`
	Payload  string `json:"payload"`
	Port     int    `json:"port"`
	Request  string `json:"request"`
	Response string `json:"response"`
}

func (s *Service) getVulInfo(c *gin.Context) {
	var formData VulInfo
	err := c.BindJSON(&formData)
	if err != nil {
		c.JSON(400, gin.H{"errcode": 400, "description": "Post Data Err"})
	}
	s.add(formData)
}

// 向数据库中插入漏洞信息
func (s *Service) add(data VulInfo) {
	vulData := &Vul{
		Host:     data.Detail.Host,
		Port:     data.Detail.Port,
		Url:      data.Detail.Url,
		Title:    data.Plugin,
		Payload:  data.Detail.Payload,
		Request:  data.Detail.Request,
		Response: data.Detail.Response,
	}
	s.Mysql.Create(vulData)
}
