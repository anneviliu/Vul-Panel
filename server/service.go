package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Service struct {
	Conf   *Config
	Router *gin.Engine
	Db     *gorm.DB
}

func (s *Service) init() {
	s.initConfig()
	s.initDb()
	s.Router = s.initRouter()
	panic(s.Router.Run(s.Conf.Base.Port))
}
