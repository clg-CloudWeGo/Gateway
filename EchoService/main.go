package main

import (
	api "github.com/clg-CloudWeGo/Echo/kitex_gen/api/echo"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

func main() {
	add, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8887")
	svr := api.NewServer(new(EchoImpl), server.WithServiceAddr(add))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
