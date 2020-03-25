package main

type Service struct {
	Conf *Config
}

func (s *Service) init() {
	s.initConfig()
	s.getTask()
}
