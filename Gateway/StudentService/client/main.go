package main

import (
	"context"
	"github.com/Raccoon-njuse/rpcsvr/kitex_gen/demo"
	"github.com/Raccoon-njuse/rpcsvr/kitex_gen/demo/studentservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
)

func main() {
	cli, err := studentservice.NewClient("nju.student.rpc",
		client.WithHostPorts("127.0.0.1:8889"),
	)

	if err != nil {
		panic("unable to init client")
	}

	resp, err := cli.Query(context.Background(), &demo.QueryReq{
		Id: 1,
	})
	if err != nil {
		panic("send req failed: " + err.Error())
	}
	klog.Info("resp:", resp)
}
