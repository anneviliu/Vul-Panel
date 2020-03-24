package main

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	Mysql `toml:"mysql"`
	Base  `toml:"base"`
}

type Base struct {
	Port      string
	APPId     string
	APPSecret string
}

type Mysql struct {
	DBHost     string
	DBUsername string
	DBPassword string
	DBName     string
}

func (s *Service) initConfig() {
	var conf *Config
	_, err := toml.DecodeFile("./conf/config.toml", &conf)
	if err != nil {
		log.Fatalln(err)
	}

	s.Conf = conf
	log.Println("config load success")
}
