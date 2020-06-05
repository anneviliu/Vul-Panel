package main

import "github.com/jinzhu/gorm"

// 漏洞信息 数据库模型
type Vul struct {
	gorm.Model
	Host         string
	Port         int
	Url          string `gorm:"size:1000"`
	Title        string
	Details      string `sql:"TYPE:json"`
	Payload      string `gorm:"size:999999"`
	Request      string `gorm:"size:999999"`
	Response     string `gorm:"size:999999"`
	Times        int64  `gorm:"size:100"`
	VulClass     string `gorm:"size:100"`
	TempFilename string `gorm:"size:500"`
	Read         bool   `gorm:"default:false"`  // 是否已读
	Status       string `gorm:"default:'wait'"` // 漏洞是否有效
}
