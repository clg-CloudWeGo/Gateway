package main

import (
	demo "github.com/Raccoon-njuse/rpcsvr/kitex_gen/demo/studentservice"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

func main() {
	add, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8889")
	s := new(StudentServiceImpl)
	s.InitDB()
	svr := demo.NewServer(s, server.WithServiceAddr(add))

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
