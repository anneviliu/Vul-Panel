package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Service struct {
	Conf   *Config
	Router *gin.Engine
	Mysql  *gorm.DB
}

func (s *Service) init() {
	s.initConfig()
	s.initMysql()
	s.Router = s.initRouter()
	panic(s.Router.Run(s.Conf.Base.Port))
}
