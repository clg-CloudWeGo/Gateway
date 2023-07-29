package main

import (
	"context"
	"github.com/clg-CloudWeGo/Echo/kitex_gen/api"
	"github.com/clg-CloudWeGo/Echo/kitex_gen/api/echo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"log"
	"time"
)

func main() {
	c, err := echo.NewClient("example-server", client.WithHostPorts("0.0.0.0:8887"))
	if err != nil {
		log.Fatal(err)
	}
	req := &api.Request{Message: "my request"}
	resp, err := c.Echo(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}
