package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
)

type Pushed struct {
	gorm.Model
	Host     string
	Port     int
	Url      string `gorm:"size:1000"`
	Title    string
	Payload  string `gorm:"size:999999"`
	Request  string `gorm:"size:999999"`
	Response string `gorm:"size:999999"`
	Times    int64  `gorm:"size:100"`
}

func (s *Service) addPushed(data VulInfo) {
	pushed := Pushed{
		Model:    gorm.Model{},
		Host:     data.Detail.Host,
		Port:     data.Detail.Port,
		Url:      data.Detail.Url,
		Title:    data.Plugin,
		Payload:  data.Detail.Payload,
		Request:  data.Detail.Request,
		Response: data.Detail.Response,
		Times:    data.Timestamp,
	}
	s.Mysql.Create(&pushed)
}

func (s *Service) getLastTime(c *gin.Context) {
	var time Pushed
	err := s.Mysql.Where(&Pushed{}).Last(&time).Error
	if err != nil {
		c.JSON(500, gin.H{"code": 500, "msg": "数据库错误"})
	}
	res := time.Times
	c.JSON(200, gin.H{"code": 200, "msg": strconv.FormatInt(res, 10)})
}
