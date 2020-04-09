package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func (s *Service) initMysql() {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&loc=Local&charset=utf8mb4,utf8",
		s.Conf.Mysql.DBUsername,
		s.Conf.Mysql.DBPassword,
		s.Conf.Mysql.DBHost,
		s.Conf.Mysql.DBName,
	))

	if err != nil {
		log.Fatalln(err)
	}

	s.Mysql = db
	s.Mysql.AutoMigrate(&Vul{}, &Pushed{}, &RegUser{})
}
