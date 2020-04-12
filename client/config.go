package main

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	Base `toml:"base"`
}

type Base struct {
	ServerIP string
	Debug    bool
}

func (s *Service) initConfig() {
	var conf *Config
	_, err := toml.DecodeFile("./config.toml", &conf)
	if err != nil {
		log.Fatalln(err)
	}
	s.Conf = conf
}
