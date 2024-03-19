package main

import (
	"sign/conf"
	"sign/server"
)

// @title sign服务
// @version 1.0
func main() {
	if err := conf.Init(); err != nil {
		panic(err)
	}

	s := server.NewServer()
	if err := s.Init(conf.Conf); err != nil {
		panic(err)
	}

	go s.GrpcRun()

	err := s.GinRun()
	if err != nil {
		panic(err)
	}
}
