package main

import (
	api "github.com/clg-CloudWeGo/Echo/kitex_gen/api/echo"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func main() {
	add, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8887")
	s := new(EchoImpl)
	svr := InitEtcdRegistry(s, "EchoService", add)
	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
func InitEtcdRegistry(s *EchoImpl, serviceName string, addr *net.TCPAddr) server.Server {
	r, err := etcd.NewEtcdRegistry([]string{"localhost:2379"})
	if err != nil {
		log.Fatal("Error: fail to new etcd registry---" + err.Error())
	}

	ebi := &rpcinfo.EndpointBasicInfo{
		ServiceName: serviceName,
		Tags:        map[string]string{"Cluster": serviceName + "Cluster"},
	}

	svr := api.NewServer(s, server.WithRegistry(r), server.WithServiceAddr(addr), server.WithServerBasicInfo(ebi))
	return svr
}
