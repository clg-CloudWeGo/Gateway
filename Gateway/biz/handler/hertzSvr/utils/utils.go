package utils

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/adaptor"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/connpool"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/generic"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
)

// NewResolver 获取注册中心的resolver
func NewResolver() discovery.Resolver {
	r, err := etcd.NewEtcdResolver([]string{"localhost:2379"})
	if err != nil {
		log.Fatal("Error: fail to new etcd resolver---" + err.Error())
	}
	return r
}

// NewProvider 获取provider
// content 为热加载idl的内容
func NewProvider(content string) (*generic.ThriftContentProvider, error) {
	p, err := generic.NewThriftContentProvider(content, map[string]string{})
	return p, err
}

// NewClient 获取泛化调用的client
func NewClient(destServiceName string, provider *generic.ThriftContentProvider, resolver discovery.Resolver) (genericclient.Client, error) {
	g, err := generic.HTTPThriftGeneric(provider)
	if err != nil {
		return nil, err
	}
	// 对client的设置
	var opts []client.Option
	opts = append(opts, client.WithResolver(resolver))
	opts = append(opts, client.WithLongConnection(connpool.IdleConfig{
		MaxIdlePerAddress: 1000,
		MaxIdleGlobal:     1000 * 10,
	}))
	opts = append(opts, client.WithTag("Cluster", destServiceName+"Cluster"))
	cli, err := genericclient.NewClient(destServiceName, g, opts...)
	return cli, err
}

// GetHTTPGenericResponse 获取http泛化调用后的response
func GetHTTPGenericResponse(ctx context.Context, c *app.RequestContext, methodName string, cli genericclient.Client) (*generic.HTTPResponse, error) {
	httpReq, err := adaptor.GetCompatRequest(c.GetRequest())
	customReq, err := generic.FromHTTPRequest(httpReq)
	if err != nil {
		return nil, err
	}
	// customReq *generic.HttpRequest
	// 由于 hertz 泛化的 method 是通过 bam 规则从 hertz request 中获取的，所以填空就行
	fmt.Println(string(c.GetRequest().Body()))
	fmt.Println(cli)
	fmt.Println("this is genertic method name" + methodName)
	resp, err := cli.GenericCall(ctx, methodName, customReq)
	if err != nil {
		return nil, err
	}
	realResp := resp.(*generic.HTTPResponse)
	return realResp, nil
}
