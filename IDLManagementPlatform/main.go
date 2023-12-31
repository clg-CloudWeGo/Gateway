// Code generated by hertz generator.

package main

import (
	"github.com/clg-CloudWeGo/idlmanagementservice/biz/handler/hertzSvr/service"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	service.DB.AutoMigrate(&service.IDLMessage{})
	h := server.New(server.WithHostPorts(":6666"))

	register(h)
	h.Spin()
}
